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

func GetReturnPriorityCountOfProjectOpen(project string) (map[string]interface{}, int) {
	issues, err := repository.DbCon.GetRepository().ReturnPriorityCountOfProjectOpen(project)
	if err != nil {
		return u.Message(false, err.Error(),
			"Jira Analyzer Backend GetReturnPriorityCountOfProject", project), http.StatusBadRequest
	}

	resp := u.Message(true, "success",
		"Jira Analyzer Backend GetReturnPriorityCountOfProject", project)
	resp["data"] = issues

	return resp, http.StatusOK
}

func GetReturnPriorityCountOfProjectClose(project string) (map[string]interface{}, int) {
	issues, err := repository.DbCon.GetRepository().ReturnPriorityCountOfProjectClose(project)
	if err != nil {
		return u.Message(false, err.Error(),
			"Jira Analyzer Backend GetReturnPriorityCountOfProject", project), http.StatusBadRequest
	}

	resp := u.Message(true, "success",
		"Jira Analyzer Backend GetReturnPriorityCountOfProject", project)
	resp["data"] = issues

	return resp, http.StatusOK
}

func GetReturnTaskStateTime(project string) (map[string]interface{}, int) {
	openTasks, err := repository.DbCon.GetRepository().ReturnCountTimeOfOpenStateInCloseTask(project)
	if err != nil {
		return u.Message(false, err.Error(),
			"Jira Analyzer Backend GetReturnCountTimeOfOpenStateInCloseTask", project), http.StatusBadRequest
	}
	resolveTask, err := repository.DbCon.GetRepository().ReturnCountTimeOfResolvedStateInCloseTask(project)
	if err != nil {
		return u.Message(false, err.Error(),
			"Jira Analyzer Backend GetReturnCountTimeOfResolvedStateInCloseTask", project), http.StatusBadRequest
	}

	reopenedTask, err := repository.DbCon.GetRepository().ReturnCountTimeOfReopenedStateInCloseTask(project)
	if err != nil {
		return u.Message(false, err.Error(),
			"Jira Analyzer Backend GetReturnCountTimeOfReopenedStateInCloseTask", project), http.StatusBadRequest
	}

	inProgressTask, err := repository.DbCon.GetRepository().ReturnCountTimeOfInProgressStateInCloseTask(project)
	if err != nil {
		return u.Message(false, err.Error(),
			"Jira Analyzer Backend GetReturnCountTimeOfInProgressStateInCloseTask", project), http.StatusBadRequest
	}

	resp := u.Message(true, "success",
		"Jira Analyzer REST API GetReturnTaskStateTime", project)
	resp["Open"] = openTasks
	resp["Resolve"] = resolveTask
	resp["Reopen"] = reopenedTask
	resp["In progress"] = inProgressTask

	return resp, http.StatusOK
}

func GetReturnActivityByTask(project string) (map[string]interface{}, int) {
	closeTasks, err := repository.DbCon.GetRepository().ReturnCountCloseTaskInDay(project)
	if err != nil {
		return u.Message(false, err.Error(),
			"Jira Analyzer Backend GetReturnCountCloseTaskInDay", project), http.StatusBadRequest
	}
	openTasks, err := repository.DbCon.GetRepository().ReturnCountOpenTaskInDay(project)
	if err != nil {
		return u.Message(false, err.Error(),
			"Jira Analyzer Backend GetReturnCountOpenTaskInDay", project), http.StatusBadRequest
	}

	resp := u.Message(true, "success",
		"Jira Analyzer Backend GetReturnActivityByTask", project)
	resp["Open"] = openTasks
	resp["Close"] = closeTasks

	return resp, http.StatusOK
}
