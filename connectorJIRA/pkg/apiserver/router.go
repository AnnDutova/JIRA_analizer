package apiserver

import (
	"connectorJIRA/pkg/connector"
	"connectorJIRA/pkg/datapusher"
	"connectorJIRA/pkg/datatransformer"
	"context"
	"encoding/json"
	"errors"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
	"net/http"
	"strconv"
	"sync"
)

type Router struct {
	logger            *logrus.Logger
	dbConnector       *datapusher.PSQLConnector
	JIRAConnector     *connector.Connection
	issueInOneRequest uint
	threadCount       uint
}

func configureRouters(r *Router) {
	http.HandleFunc("/test", r.handleTestAnswer)
	http.HandleFunc("/issues", r.handleIssues)
	http.HandleFunc("/updateProject", r.handleUpdateProject)
	http.HandleFunc("/projects", r.handleProjects)
}

func getProjectsParams(r *http.Request) (int, int, string, error) {
	limitStr := r.FormValue("limit")
	pageStr := r.FormValue("page")
	search := r.FormValue("search")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		if limitStr == "" {
			limit = 20
		} else {
			return 0, 0, "", errors.New("Parameter limit is not a number")
		}
	}
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		if pageStr == "" {
			page = 1
		} else {
			return 0, 0, "", errors.New("Parameter page is not a number")
		}
	}
	if page < 1 {
		return 0, 0, "", errors.New("Parameter page cannot be less than 1")
	}
	if limit < 1 {
		return 0, 0, "", errors.New("Parameter limit cannot be less than 1")
	}
	return limit, page, search, nil
}

func correctThreadCount(issueInOneRequest int, threadCount int, total int) int {
	if issueInOneRequest*(threadCount-1) >= total {
		if total%issueInOneRequest == 0 {
			threadCount = total / issueInOneRequest
		} else {
			threadCount = total/issueInOneRequest + 1
		}
	}
	return threadCount
}

func getStartAndEnd(total int, threadCount int, issueInOneRequest int, threadNum int) (int, int) {
	endAt := float64(total) / float64(threadCount) * float64(threadNum+1)
	var startAt int
	if threadNum == 0 {
		startAt = 0
	} else {
		startAtFloat := float64(total) / float64(threadCount) * float64(threadNum)
		if int(startAtFloat)%issueInOneRequest == 0 {
			startAt = int(startAtFloat)
		} else {
			startAt = int(startAtFloat) + issueInOneRequest - int(startAtFloat)%issueInOneRequest
		}
	}
	return startAt, int(endAt)
}

func checkContext(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		return nil
	}
}

func getIssuesWithHistories(con *connector.Connection, project string, startAt int,
	issueInOneRequest int) ([]datatransformer.Issue, []datatransformer.IssueStatusChanges, error) {
	issues, err := con.GetExpandIssuesJSON(project, startAt, issueInOneRequest)
	if err != nil {
		return nil, nil, errors.New("Error when GetExpandIssuesJSON:" + err.Error())
	}
	formattedIssues, err := datatransformer.FormatIssues(issues)
	if err != nil {
		return nil, nil, errors.New("Error when FormatIssues:" + err.Error())
	}

	histories := make([]datatransformer.IssueStatusChanges, len(formattedIssues))

	for inx, issue := range formattedIssues {
		changelog, err := con.GetIssueChangelogJSON(issue.Key)
		if err != nil {
			return nil, nil, errors.New("Error when GetIssueChangelogJSON:" + err.Error())
		}
		statusChanges, err := datatransformer.FormatChangelog(changelog)
		if err != nil {
			return nil, nil, errors.New("Error when FormatChangelog:" + err.Error())
		}
		histories[inx] = statusChanges
	}

	return formattedIssues, histories, nil
}

func (rout *Router) handleTestAnswer(rw http.ResponseWriter, r *http.Request) {
	respond(rw, r, http.StatusOK, "test")
}

func (rout *Router) handleProjects(rw http.ResponseWriter, r *http.Request) {
	rout.logger.Info("/projects is called")
	rout.logger.Info("Check parameters validity")
	limit, page, search, err := getProjectsParams(r)
	if err != nil {
		rout.logger.Warning(err)
		parseError(rw, r, http.StatusBadRequest, err)
		return
	}
	rout.logger.Infof("Parameters page = %d and limit = %d is valid", page, limit)
	rout.logger.Info("Begin load projects")
	projects, err := rout.JIRAConnector.GetAllFormattedProjects(limit, page, search)
	if err != nil {
		rout.logger.Error(err)
		parseError(rw, r, http.StatusBadRequest, err)
		return
	}
	rout.logger.Info("Projects are loaded, return projects")
	respond(rw, r, http.StatusOK, projects)
}

func (rout *Router) handleIssues(rw http.ResponseWriter, r *http.Request) {
	rout.logger.Info("/issues is called")
	project := r.FormValue("project")
	rout.logger.Info("Load issues")
	issues, err := rout.JIRAConnector.GetFormattedIssues(project)
	if err != nil {
		rout.logger.Error(err)
		parseError(rw, r, http.StatusBadRequest, err)
		return
	}
	respond(rw, r, http.StatusOK, issues)
}

func (rout *Router) handleUpdateProject(rw http.ResponseWriter, r *http.Request) {
	project := r.FormValue("project")
	rout.logger.Infof("/updateProject is called with parameter project = %s", project)
	total, err := rout.JIRAConnector.GetTotalIssues(project)
	if err != nil {
		rout.logger.Error(err)
		parseError(rw, r, http.StatusBadRequest, err)
		return
	}
	if total == 0 {
		projectId, projectName, err := rout.JIRAConnector.GetProjectInfo(project)
		if err != nil {
			rout.logger.Error("Error when get project info from JIRA: " + err.Error())
			parseError(rw, r, http.StatusBadRequest, err)
			return
		}
		if err := rout.dbConnector.AddProject(projectName, projectId); err != nil {
			rout.logger.Error("Error when push data to DB: " + err.Error())
			parseError(rw, r, http.StatusBadRequest, err)
			return
		}
		return
	}

	issueInOneRequest := int(rout.issueInOneRequest)
	threadCount := int(rout.threadCount)
	threadCount = correctThreadCount(issueInOneRequest, threadCount, total)
	rout.logger.Infof("threadCount = %d", threadCount)

	g, ctx := errgroup.WithContext(context.Background())
	var m sync.Mutex

	var formattedIssues []datatransformer.Issue
	var histories []datatransformer.IssueStatusChanges

	for i := 0; i < threadCount; i++ {
		i := i
		g.Go(func() error {
			startAt, endAt := getStartAndEnd(total, threadCount, issueInOneRequest, i)
			rout.logger.Infof("Thread №%d StartAt:%d EndAt:%d", i+1, startAt, int(endAt))

			for startAt < endAt {
				err = checkContext(ctx)
				if err != nil {
					return err
				}
				rout.logger.Infof("Thread №%d begin load data with StartAt:%d", i+1, startAt)
				formattedIssuesInThread, historiesInThread, err := getIssuesWithHistories(rout.JIRAConnector, project, startAt, issueInOneRequest)
				if err != nil {
					rout.logger.Error("Error in thread " + strconv.Itoa(i+1) + " when getIssuesWithHistories:" + err.Error())
					return err
				}

				m.Lock()
				formattedIssues = append(formattedIssues, formattedIssuesInThread...)
				histories = append(histories, historiesInThread...)
				m.Unlock()

				startAt += issueInOneRequest
			}
			return nil
		})

	}
	if err := g.Wait(); err != nil {
		parseError(rw, r, http.StatusBadRequest, err)
		return
	}
	rout.logger.Info("Begin update data in DB")
	if err := rout.dbConnector.UpdateData(formattedIssues, histories); err != nil {
		rout.logger.Error("Error when push data to DB: " + err.Error())
		parseError(rw, r, http.StatusBadRequest, err)
		return
	}
	rout.logger.Info("Project is updated, return positive response")
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
