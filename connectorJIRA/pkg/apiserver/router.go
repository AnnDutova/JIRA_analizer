package apiserver

import (
	"connectorJIRA/pkg/connector"
	"connectorJIRA/pkg/datapusher"
	"connectorJIRA/pkg/datatransformer"
	"connectorJIRA/pkg/properties"
	"encoding/json"
	"net/http"
	"os"
)

func GetIssues(connectionApache *connector.Connection, projectName string) []datatransformer.Issue {
	startAt := 0
	total, _ := connectionApache.GetTotalIssues(projectName)
	var issues []datatransformer.Issue
	for ; startAt < total; startAt += 50 {
		issuesRaw, _ := connectionApache.GetExpandIssuesJSON(projectName, startAt)
		formattedIssues, _ := datatransformer.FormatIssues(issuesRaw)
		issues = append(issues, formattedIssues...)
	}
	return issues
}

type Router struct {
	dbConnector   *datapusher.PSQLConnector
	JIRAConnector *connector.Connection
}

func configureRouters(r *Router) {
	http.HandleFunc("/test", r.handleTestAnswer)
	http.HandleFunc("/issues", r.handleIssues)
	http.HandleFunc("/updateProject", r.handleUpdateProject)
}

func (rout *Router) handleTestAnswer(rw http.ResponseWriter, r *http.Request) {
	respond(rw, r, http.StatusOK, "test")
}

func (rout *Router) handleIssues(rw http.ResponseWriter, r *http.Request) {
	config := properties.GetConfig(os.Args[1])
	issues := GetIssues(rout.JIRAConnector, config.ProgramSettings.ProjectNames)
	respond(rw, r, http.StatusOK, issues)
}

func (rout *Router) handleUpdateProject(rw http.ResponseWriter, r *http.Request) {
	config := properties.GetConfig(os.Args[1])
	issues := GetIssues(rout.JIRAConnector, config.ProgramSettings.ProjectNames)
	histories := make([]datatransformer.IssueStatusChanges, len(issues))

	for i, issue := range issues {
		changelog := rout.JIRAConnector.GetIssueChangelogJSON(issue.Key)
		statusChanges := datatransformer.FormatChangelog(changelog)
		histories[i] = statusChanges
	}

	if err := rout.dbConnector.PushFirstData(issues, histories); err != nil {
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
