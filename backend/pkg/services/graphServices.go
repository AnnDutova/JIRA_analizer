package services

import (
	"Backend/pkg/repository"
	u "Backend/pkg/utils"
	"net/http"
)

func GetReturnTimeCountOfIssuesInCloseState(project string) (map[string]interface{}, int) {
	issues, err := repository.DbCon.GetRepository().ReturnTimeCountOfIssuesInCloseState(project)
	if err != nil {
		return u.Message(false, err.Error(),
			"Jira Analyzer Backend GetReturnTimeCountOfIssuesInCloseState", project), http.StatusBadRequest
	}

	resp := u.Message(true, "success",
		"Jira Analyzer Backend GetReturnTimeCountOfIssuesInCloseState", project)
	resp["data"] = issues

	return resp, http.StatusOK
}

func GetReturnTimeSpentOnAllTasks(project string) (map[string]interface{}, int) {
	issues, err := repository.DbCon.GetRepository().ReturnTimeSpentOnAllTasks(project)
	if err != nil {
		return u.Message(false, err.Error(),
			"Jira Analyzer Backend GetReturnTimeSpentOnAllTasks", project), http.StatusBadRequest
	}

	resp := u.Message(true, "success",
		"Jira Analyzer Backend GetReturnTimeSpentOnAllTasks", project)
	resp["data"] = issues

	return resp, http.StatusOK
}

func GetReturnTheMostActiveCreators(project string) (map[string]interface{}, int) {
	issues, err := repository.DbCon.GetRepository().ReturnTheMostActiveCreators(project)
	if err != nil {
		return u.Message(false, err.Error(),
			"Jira Analyzer Backend GetReturnTheMostActiveCreators", project), http.StatusBadRequest
	}

	resp := u.Message(true, "success",
		"Jira Analyzer Backend GetReturnTheMostActiveCreators", project)
	resp["data"] = issues

	return resp, http.StatusOK
}
