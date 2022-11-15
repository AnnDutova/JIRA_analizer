package apiserver

import (
	"connectorJIRA/pkg/connector"
	"connectorJIRA/pkg/datapusher"
	"connectorJIRA/pkg/datatransformer"
	"encoding/json"
	"net/http"
)

type Router struct {
	dbConnector   *datapusher.PSQLConnector
	JIRAConnector *connector.Connection
}

func configureRouters(r *Router) {
	http.HandleFunc("/test", r.handleTestAnswer)
	http.HandleFunc("/issues", r.handleIssues)
	http.HandleFunc("/updateProject", r.handleUpdateProject)
	http.HandleFunc("/allProjects", r.handleAllProjects)
}

func (rout *Router) handleTestAnswer(rw http.ResponseWriter, r *http.Request) {
	respond(rw, r, http.StatusOK, "test")
}

func (rout *Router) handleAllProjects(rw http.ResponseWriter, r *http.Request) {
	projects, _ := rout.JIRAConnector.GetAllFormattedProjects()
	respond(rw, r, http.StatusOK, projects)
}

func (rout *Router) handleIssues(rw http.ResponseWriter, r *http.Request) {
	project := r.FormValue("project")
	issues, _ := rout.JIRAConnector.GetFormattedIssues(project)
	respond(rw, r, http.StatusOK, issues)
}

func (rout *Router) handleUpdateProject(rw http.ResponseWriter, r *http.Request) {
	project := r.FormValue("project")
	issues, _ := rout.JIRAConnector.GetFormattedIssues(project)
	histories := make([]datatransformer.IssueStatusChanges, len(issues))

	for i, issue := range issues {
		changelog := rout.JIRAConnector.GetIssueChangelogJSON(issue.Key)
		statusChanges := datatransformer.FormatChangelog(changelog)
		histories[i] = statusChanges
	}

	if err := rout.dbConnector.UpdateData(issues, histories); err != nil {
		parseError(rw, r, http.StatusBadRequest, err)
	}
	respond(rw, r, http.StatusOK, "push")
}

func parseError(w http.ResponseWriter, r *http.Request, code int, err error) {
	respond(w, r, code, map[string]string{"error": err.Error()})
}

func respond(w http.ResponseWriter, r *http.Request, code int, date interface{}) {
	w.WriteHeader(code)
	if date != nil {
		json.NewEncoder(w).Encode(date)
	}
}
