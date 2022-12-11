package services

import (
	"Backend/pkg/repository"
	u "Backend/pkg/utils"
	"net/http"
)

func GetReturnTimeCountOfIssuesInCloseState(project string) (map[string]interface{}, int) {
	data := make(map[string]interface{}, 0)
	issues, err := repository.DbCon.GetRepository().ReturnTimeCountOfIssuesInCloseState(project)
	if err != nil {
		return u.Message(false, err.Error(),
			"Jira Analyzer Backend GetReturnTimeCountOfIssuesInCloseState", project), http.StatusBadRequest
	}

	resp := u.Message(true, "success",
		"Jira Analyzer Backend GetReturnTimeCountOfIssuesInCloseState", project)
	if len(issues) > 0 {
		category := make(map[string]interface{}, len(issues))
		for _, el := range issues {
			category[el.Title] = el.Count
		}
		data["categories"] = u.SortCategories(category)
		data["count"] = category
		resp["data"] = data
	} else {
		resp["data"] = nil
	}

	return resp, http.StatusOK
}

func GetReturnTimeSpentOnAllTasks(project string) (map[string]interface{}, int) {
	data := make(map[string]interface{}, 0)
	issues, err := repository.DbCon.GetRepository().ReturnTimeSpentOnAllTasks(project)
	if err != nil {
		return u.Message(false, err.Error(),
			"Jira Analyzer Backend GetReturnTimeSpentOnAllTasks", project), http.StatusBadRequest
	}

	resp := u.Message(true, "success",
		"Jira Analyzer Backend GetReturnTimeSpentOnAllTasks", project)
	if len(issues) > 0 {
		category := make(map[string]interface{}, len(issues))
		for _, el := range issues {
			category[el.Title] = el.Count
		}
		data["count"] = category
		data["categories"] = u.SortMinutesCategories(category)
	} else {
		data["count"] = nil
		data["categories"] = nil
	}
	resp["data"] = data
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
	if len(issues) > 0 {
		resp["data"] = issues
		category := make(map[string]interface{}, len(issues))
		for _, el := range issues {
			category[el.Title] = el.Count
		}
		resp["categories"] = u.SortCategories(category)
	} else {
		resp["data"] = nil
		resp["categories"] = nil
	}

	return resp, http.StatusOK
}

func GetReturnPriorityCountOfProjectOpen(project string) (map[string]interface{}, int) {
	data := make(map[string]interface{}, 0)
	issues, err := repository.DbCon.GetRepository().ReturnPriorityCountOfProjectOpen(project)
	if err != nil {
		return u.Message(false, err.Error(),
			"Jira Analyzer Backend GetReturnPriorityCountOfProject", project), http.StatusBadRequest
	}

	resp := u.Message(true, "success",
		"Jira Analyzer Backend GetReturnPriorityCountOfProject", project)
	if len(issues) > 0 {
		category := make([]string, 0, len(issues))
		names := make(map[string]interface{}, 0)
		for _, el := range issues {
			category = append(category, el.Title)
			names[el.Title] = el.Count
		}
		data["count"] = names
		data["categories"] = category
	} else {
		data["count"] = nil
		data["categories"] = nil
	}
	resp["data"] = data
	return resp, http.StatusOK
}

func GetReturnPriorityCountOfProjectClose(project string) (map[string]interface{}, int) {
	data := make(map[string]interface{}, 0)
	issues, err := repository.DbCon.GetRepository().ReturnPriorityCountOfProjectClose(project)
	if err != nil {
		return u.Message(false, err.Error(),
			"Jira Analyzer Backend GetReturnPriorityCountOfProject", project), http.StatusBadRequest
	}

	resp := u.Message(true, "success",
		"Jira Analyzer Backend GetReturnPriorityCountOfProject", project)
	if len(issues) > 0 {
		category := make([]string, 0, len(issues))
		names := make(map[string]interface{}, 0)
		for _, el := range issues {
			category = append(category, el.Title)
			names[el.Title] = el.Count
		}
		data["count"] = names
		data["categories"] = category
	} else {
		data["count"] = nil
		data["categories"] = nil
	}
	resp["data"] = data
	return resp, http.StatusOK
}

func GetReturnTaskStateTime(project string) (map[string]interface{}, int) {
	open := make(map[string]interface{}, 0)
	resolve := make(map[string]interface{}, 0)
	reopen := make(map[string]interface{}, 0)
	progress := make(map[string]interface{}, 0)
	category := make(map[string]interface{}, 0)
	data := make(map[string]interface{}, 0)

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

	if len(openTasks) > 0 {
		for _, el := range openTasks {
			open[el.Title] = el.Count
		}
		data["open"] = open
		category["open"] = u.SortCategories(open)
	} else {
		data["open"] = nil
		category["open"] = nil
	}
	if len(resolveTask) > 0 {
		for _, el := range resolveTask {
			resolve[el.Title] = el.Count
		}
		data["resolve"] = resolve
		category["resolve"] = u.SortCategories(resolve)
	} else {
		data["resolve"] = nil
		category["resolve"] = nil
	}
	if len(reopenedTask) > 0 {
		for _, el := range reopenedTask {
			reopen[el.Title] = el.Count
		}
		data["reopen"] = reopen
		category["reopen"] = u.SortCategories(reopen)
	} else {
		data["reopen"] = nil
		category["reopen"] = nil
	}
	if len(inProgressTask) > 0 {
		for _, el := range inProgressTask {
			progress[el.Title] = el.Count
		}
		data["progress"] = progress
		category["progress"] = u.SortCategories(progress)
	} else {
		data["progress"] = nil
		category["progress"] = nil
	}
	data["categories"] = category
	resp["data"] = data
	return resp, http.StatusOK
}

func GetReturnActivityByTask(project string) (map[string]interface{}, int) {
	open := make(map[string]interface{}, 0)
	closed := make(map[string]interface{}, 0)
	category := make(map[string]interface{}, 0)
	data := make(map[string]interface{}, 0)

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
	if len(openTasks) > 0 {
		for _, el := range openTasks {
			open[el.Title] = el.Count
		}
		openDates, err := u.SortDatesForActivityGraph(open)
		if err != nil {
			return u.Message(false, err.Error(),
				"Jira Analyzer Backend GetReturnActivityByTask. Fail on SortDatesForActivityGraph",
				"Dates for open tasks"), http.StatusBadRequest
		}
		data["open"] = open
		category["open"] = openDates
	} else {
		data["open"] = nil
		category["open"] = nil
	}
	if len(closeTasks) > 0 {
		for _, el := range closeTasks {
			closed[el.Title] = el.Count
		}
		closeDates, err := u.SortDatesForActivityGraph(closed)
		if err != nil {
			return u.Message(false, err.Error(),
				"Jira Analyzer Backend GetReturnActivityByTask. Fail on SortDatesForActivityGraph",
				"Dates for close tasks"), http.StatusBadRequest
		}
		data["close"] = closed
		category["close"] = closeDates
	} else {
		data["close"] = nil
		category["close"] = nil
	}
	data["categories"] = category
	resp["data"] = data
	return resp, http.StatusOK
}
