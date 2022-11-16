package apiserver

import (
	"connectorJIRA/pkg/connector"
	"connectorJIRA/pkg/datapusher"
	"connectorJIRA/pkg/datatransformer"
	"context"
	"encoding/json"
	"errors"
	"golang.org/x/sync/errgroup"
	"net/http"
	"strconv"
	"sync"
)

type Router struct {
	dbConnector       *datapusher.PSQLConnector
	JIRAConnector     *connector.Connection
	issueInOneRequest uint
	threadCount       uint
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
	total, err := rout.JIRAConnector.GetTotalIssues(project)
	if err != nil {
		parseError(rw, r, http.StatusBadRequest, err)
		return
	}
	if total == 0 {
		parseError(rw, r, http.StatusBadRequest, errors.New("This project is empty"))
		return
	}

	issueInOneRequest := int(rout.issueInOneRequest)
	threadCount := int(rout.threadCount)

	if issueInOneRequest*(threadCount-1) >= total {
		if total%issueInOneRequest == 0 {
			threadCount = total / issueInOneRequest
		} else {
			threadCount = total/issueInOneRequest + 1
		}
	}

	g, _ := errgroup.WithContext(context.Background())
	var m1 sync.Mutex
	var m2 sync.Mutex

	var formattedIssues []datatransformer.Issue
	var histories []datatransformer.IssueStatusChanges

	for i := 0; i < threadCount; i++ {
		i := i
		g.Go(func() error {
			endAt := float64(total) / float64(threadCount) * float64(i+1)
			var startAt int
			if i == 0 {
				startAt = 0
			} else {
				startAtFloat := float64(total) / float64(threadCount) * float64(i)
				if int(startAtFloat)%issueInOneRequest == 0 {
					startAt = int(startAtFloat)
				} else {
					startAt = int(startAtFloat) + issueInOneRequest - int(startAtFloat)%issueInOneRequest
				}
			}

			for startAt < int(endAt) {
				issues, err := rout.JIRAConnector.GetExpandIssuesJSON(project, startAt, issueInOneRequest)
				if err != nil {
					return errors.New("Error in thread " + strconv.Itoa(i+1) + " when GetExpandIssuesJSON:" + err.Error())
				}
				formattedIssuesInThread, err := datatransformer.FormatIssues(issues)
				if err != nil {
					return errors.New("Error in thread " + strconv.Itoa(i+1) + " when FormatIssues:" + err.Error())
				}

				m1.Lock()
				formattedIssues = append(formattedIssues, formattedIssuesInThread...)
				m1.Unlock()

				historiesInThread := make([]datatransformer.IssueStatusChanges, len(formattedIssuesInThread))

				for inx, issue := range formattedIssuesInThread {
					changelog, err := rout.JIRAConnector.GetIssueChangelogJSON(issue.Key)
					if err != nil {
						return errors.New("Error in thread " + strconv.Itoa(i+1) + " when GetIssueChangelogJSON:" + err.Error())
					}
					statusChanges := datatransformer.FormatChangelog(changelog)
					historiesInThread[inx] = statusChanges
				}

				m2.Lock()
				histories = append(histories, historiesInThread...)
				m2.Unlock()

				startAt += issueInOneRequest
			}
			return nil
		})

	}
	if err := g.Wait(); err != nil {
		parseError(rw, r, http.StatusBadRequest, err)
		return
	}
	if err := rout.dbConnector.UpdateData(formattedIssues, histories); err != nil {
		parseError(rw, r, http.StatusBadRequest, err)
		return
	}
	respond(rw, r, http.StatusOK, "Project updated in DB")
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
