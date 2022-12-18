package services

import (
	"Backend/pkg/repository"
	u "Backend/pkg/utils"
	"net/http"
)

func GetReturnTimeCountOfIssuesInCloseState(project string) (map[string]interface{}, int) {
	logger := u.GetLogger()
	logger.Info("Send GetReturnTimeCountOfIssuesInCloseState request")

	data := make(map[string]interface{}, 0)
	resp := make(map[string]interface{}, 0)

	logger.Info("Send ReturnTimeCountOfIssuesInCloseState request")

	logger.Info("Check on emptiness")
	res, err := repository.DbCon.GetRepository().IsEmpty(project)
	if err != nil {
		logger.Error("Something went wrong when check project on emptiness")
		return u.Message(false, err.Error(),
			"Jira Analyzer Backend IsEmpty", project), http.StatusBadRequest
	}
	if !res {
		issues, err := repository.DbCon.GetRepository().ReturnTimeCountOfIssuesInCloseState(project)
		if err != nil {
			logger.Error("Something went wrong ReturnTimeCountOfIssuesInCloseState")
			return u.Message(false, err.Error(),
				"Jira Analyzer Backend GetReturnTimeCountOfIssuesInCloseState", project), http.StatusBadRequest
		}

		resp = u.Message(true, "success",
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
		logger.Info("Get result of ReturnTimeCountOfIssuesInCloseState request")
		return resp, http.StatusOK
	} else {
		resp = u.Message(true, "success",
			"Jira Analyzer Backend GetReturnTimeCountOfIssuesInCloseState", project)
		data["empty"] = res
		resp["data"] = data
		logger.Info("Get result of ReturnTimeCountOfIssuesInCloseState request")
		return resp, http.StatusNoContent
	}
}

func GetReturnTimeSpentOnAllTasks(project string) (map[string]interface{}, int) {
	logger := u.GetLogger()
	logger.Info("Send GetReturnTimeSpentOnAllTasks request")

	data := make(map[string]interface{}, 0)
	resp := make(map[string]interface{}, 0)

	res, err := repository.DbCon.GetRepository().IsEmpty(project)
	if err != nil {
		logger.Error("Something went wrong when check project on emptiness")
		return u.Message(false, err.Error(),
			"Jira Analyzer Backend IsEmpty", project), http.StatusBadRequest
	}
	if !res {
		logger.Info("Send ReturnTimeSpentOnAllTasks request")
		issues, err := repository.DbCon.GetRepository().ReturnTimeSpentOnAllTasks(project)
		if err != nil {
			logger.Error("Something went wrong on ReturnTimeSpentOnAllTasks request")
			return u.Message(false, err.Error(),
				"Jira Analyzer Backend GetReturnTimeSpentOnAllTasks", project), http.StatusBadRequest
		}

		resp = u.Message(true, "success",
			"Jira Analyzer Backend GetReturnTimeSpentOnAllTasks", project)
		if len(issues) > 0 {
			category := make(map[string]interface{}, len(issues))
			for _, el := range issues {
				category[el.Title] = el.Count
			}
			data["count"] = category
			data["categories"] = u.SortMinutesCategories(category)
			resp["data"] = data
		} else {
			resp["data"] = nil
		}
		logger.Info("Get result from GetReturnTimeSpentOnAllTasks request")
		return resp, http.StatusOK
	} else {
		resp = u.Message(true, "success",
			"Jira Analyzer Backend GetReturnTimeSpentOnAllTasks", project)
		data["empty"] = res
		resp["data"] = data
		logger.Info("Get result from GetReturnTimeSpentOnAllTasks request")
		return resp, http.StatusNoContent
	}
}

func GetReturnPriorityCountOfProjectOpen(project string) (map[string]interface{}, int) {
	logger := u.GetLogger()
	logger.Info("Send GetReturnTimeSpentOnAllTasks request")

	data := make(map[string]interface{}, 0)
	resp := make(map[string]interface{}, 0)

	res, err := repository.DbCon.GetRepository().IsEmpty(project)
	if err != nil {
		logger.Error("Something went wrong when check project on emptiness")
		return u.Message(false, err.Error(),
			"Jira Analyzer Backend IsEmpty", project), http.StatusBadRequest
	}
	if !res {
		logger.Info("Send ReturnPriorityCountOfProjectOpen request")
		issues, err := repository.DbCon.GetRepository().ReturnPriorityCountOfProjectOpen(project)
		if err != nil {
			logger.Error("Something went wrong ReturnPriorityCountOfProjectOpen", err.Error())
			return u.Message(false, err.Error(),
				"Jira Analyzer Backend GetReturnPriorityCountOfProject", project), http.StatusBadRequest
		}

		resp = u.Message(true, "success",
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
			resp["data"] = data
		} else {
			resp["data"] = nil
		}
		logger.Info("Get result from GetReturnPriorityCountOfProjectOpen request")
		return resp, http.StatusOK
	} else {
		resp = u.Message(true, "success",
			"Jira Analyzer Backend GetReturnPriorityCountOfProject", project)
		data["empty"] = res
		resp["data"] = data
		logger.Info("Get result from GetReturnPriorityCountOfProjectOpen request")
		return resp, http.StatusNoContent
	}
}

func GetReturnPriorityCountOfProjectClose(project string) (map[string]interface{}, int) {
	logger := u.GetLogger()
	logger.Info("Send GetReturnPriorityCountOfProjectClose request")

	data := make(map[string]interface{}, 0)
	resp := make(map[string]interface{}, 0)

	res, err := repository.DbCon.GetRepository().IsEmpty(project)
	if err != nil {
		logger.Error("Something went wrong when check project on emptiness")
		return u.Message(false, err.Error(),
			"Jira Analyzer Backend IsEmpty", project), http.StatusBadRequest
	}
	if !res {
		logger.Info("Send ReturnPriorityCountOfProjectClose request")
		issues, err := repository.DbCon.GetRepository().ReturnPriorityCountOfProjectClose(project)
		if err != nil {
			logger.Error("Something went wrong on ReturnPriorityCountOfProjectClose request")
			return u.Message(false, err.Error(),
				"Jira Analyzer Backend GetReturnPriorityCountOfProject", project), http.StatusBadRequest
		}

		resp = u.Message(true, "success",
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
			resp["data"] = data
		} else {
			resp["data"] = nil
		}
		logger.Info("Get result of GetReturnPriorityCountOfProjectClose request")
		return resp, http.StatusOK
	} else {
		resp = u.Message(true, "success",
			"Jira Analyzer Backend GetReturnPriorityCountOfProject", project)
		data["empty"] = res
		resp["data"] = data
		return resp, http.StatusNoContent
	}
}

func GetReturnTaskStateTime(project string) (map[string]interface{}, int) {
	logger := u.GetLogger()
	logger.Info("Send GetReturnTaskStateTime request")

	open := make(map[string]interface{}, 0)
	resolve := make(map[string]interface{}, 0)
	reopen := make(map[string]interface{}, 0)
	progress := make(map[string]interface{}, 0)
	category := make(map[string]interface{}, 0)
	data := make(map[string]interface{}, 0)
	resp := make(map[string]interface{}, 0)

	res, err := repository.DbCon.GetRepository().IsEmpty(project)
	if err != nil {
		logger.Error("Something went wrong when check project on emptiness")
		return u.Message(false, err.Error(),
			"Jira Analyzer Backend IsEmpty", project), http.StatusBadRequest
	}
	if !res {
		logger.Info("Send ReturnCountTimeOfOpenStateInCloseTask request")
		openTasks, err := repository.DbCon.GetRepository().ReturnCountTimeOfOpenStateInCloseTask(project)
		if err != nil {
			logger.Error("Something went wrong on ReturnCountTimeOfOpenStateInCloseTask request")
			return u.Message(false, err.Error(),
				"Jira Analyzer Backend GetReturnCountTimeOfOpenStateInCloseTask", project), http.StatusBadRequest
		}

		logger.Info("Send ReturnCountTimeOfResolvedStateInCloseTask request")
		resolveTask, err := repository.DbCon.GetRepository().ReturnCountTimeOfResolvedStateInCloseTask(project)
		if err != nil {
			logger.Error("Something went wrong on ReturnCountTimeOfResolvedStateInCloseTask request")
			return u.Message(false, err.Error(),
				"Jira Analyzer Backend ReturnCountTimeOfResolvedStateInCloseTask", project), http.StatusBadRequest
		}

		logger.Info("Send ReturnCountTimeOfReopenedStateInCloseTask request")
		reopenedTask, err := repository.DbCon.GetRepository().ReturnCountTimeOfReopenedStateInCloseTask(project)
		if err != nil {
			logger.Error("Something went wrong on ReturnCountTimeOfReopenedStateInCloseTask request")
			return u.Message(false, err.Error(),
				"Jira Analyzer Backend ReturnCountTimeOfReopenedStateInCloseTask", project), http.StatusBadRequest
		}

		logger.Info("Send ReturnCountTimeOfInProgressStateInCloseTask request")
		inProgressTask, err := repository.DbCon.GetRepository().ReturnCountTimeOfInProgressStateInCloseTask(project)
		if err != nil {
			logger.Error("Something went wrong on ReturnCountTimeOfInProgressStateInCloseTask request")
			return u.Message(false, err.Error(),
				"Jira Analyzer Backend ReturnCountTimeOfInProgressStateInCloseTask", project), http.StatusBadRequest
		}

		resp = u.Message(true, "success",
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

		if len(openTasks) > 0 || len(reopenedTask) > 0 || len(resolveTask) > 0 || len(inProgressTask) > 0 {
			data["categories"] = category
			resp["data"] = data
		} else if len(openTasks) == 0 && len(reopenedTask) == 0 && len(resolveTask) > 0 && len(inProgressTask) > 0 {
			resp["data"] = nil
		}
		logger.Info("Get result of GetReturnTaskStateTime request")
		return resp, http.StatusOK
	} else {
		resp = u.Message(true, "success",
			"Jira Analyzer REST API GetReturnTaskStateTime", project)
		data["empty"] = res
		resp["data"] = data
		logger.Info("Get empty result of GetReturnTaskStateTime request")
		return resp, http.StatusNoContent
	}
}

func GetReturnActivityByTask(project string) (map[string]interface{}, int) {
	logger := u.GetLogger()
	logger.Info("Send GetReturnTimeSpentOnAllTasks request")

	open := make(map[string]interface{}, 0)
	closed := make(map[string]interface{}, 0)
	category := make(map[string]interface{}, 0)
	data := make(map[string]interface{}, 0)
	all := make(map[string]interface{}, 0)
	resp := make(map[string]interface{}, 0)

	res, err := repository.DbCon.GetRepository().IsEmpty(project)
	if err != nil {
		logger.Error("Something went wrong when check project on emptiness")
		return u.Message(false, err.Error(),
			"Jira Analyzer Backend IsEmpty", project), http.StatusBadRequest
	}
	if !res {
		logger.Info("Send ReturnCountCloseTaskInDay request")
		closeTasks, err := repository.DbCon.GetRepository().ReturnCountCloseTaskInDay(project)
		if err != nil {
			logger.Error("Something went wrong on ReturnCountCloseTaskInDay request")
			return u.Message(false, err.Error(),
				"Jira Analyzer Backend ReturnCountCloseTaskInDay", project), http.StatusBadRequest
		}
		logger.Info("Send ReturnCountOpenTaskInDay request")
		openTasks, err := repository.DbCon.GetRepository().ReturnCountOpenTaskInDay(project)
		if err != nil {
			logger.Error("Something went wrong on ReturnCountOpenTaskInDay request")
			return u.Message(false, err.Error(),
				"Jira Analyzer Backend ReturnCountOpenTaskInDay", project), http.StatusBadRequest
		}

		resp = u.Message(true, "success",
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
		all = u.JoinToMap(openTasks, closeTasks, all)
		if len(all) > 0 {
			allCategories, err := u.SortDatesForActivityGraph(all)
			if err != nil {
				return u.Message(false, err.Error(),
					"Jira Analyzer Backend GetReturnActivityByTask. Fail on SortDatesForActivityGraph",
					"Dates for all categories"), http.StatusBadRequest
			}
			category["all"] = allCategories
		} else {
			category["all"] = nil
		}
		if len(openTasks) > 0 || len(closeTasks) > 0 {
			data["categories"] = category
			resp["data"] = data
		} else if len(openTasks) == 0 && len(closeTasks) == 0 {
			resp["data"] = nil
		}
		logger.Info("Get result of GetReturnActivityByTask request")
		return resp, http.StatusOK
	} else {
		resp = u.Message(true, "success",
			"Jira Analyzer Backend GetReturnActivityByTask", project)
		data["empty"] = res
		resp["data"] = data
		return resp, http.StatusNoContent
	}
}

func MakeTimeCountOfIssuesInCloseState(project string) (map[string]interface{}, int) {
	logger := u.GetLogger()
	logger.Info("Send GetReturnTimeSpentOnAllTasks request")

	logger.Info("Send MakeTimeCountOfIssuesInCloseState request")
	err := repository.DbCon.GetRepository().MakeTimeCountOfIssuesInCloseState(project)
	if err != nil {
		logger.Error("Something went wrong on MakeTimeCountOfIssuesInCloseState request")
		return u.Message(false, err.Error(),
			"Jira Analyzer Backend MakeTimeCountOfIssuesInCloseState", project), http.StatusBadRequest
	}
	resp := u.Message(true, "success",
		"Jira Analyzer Backend MakeTimeCountOfIssuesInCloseState", project)
	logger.Info("Get result of MakeTimeCountOfIssuesInCloseState request")
	return resp, http.StatusOK
}

func MakeTaskStateTime(project string) (map[string]interface{}, int) {
	logger := u.GetLogger()
	logger.Info("Send GetReturnTimeSpentOnAllTasks request")

	logger.Info("Send MakeCountTimeOfOpenStateInCloseTask request")
	err := repository.DbCon.GetRepository().MakeCountTimeOfOpenStateInCloseTask(project)
	if err != nil {
		logger.Error("Something went wrong on MakeCountTimeOfOpenStateInCloseTask request")
		return u.Message(false, err.Error(),
			"Jira Analyzer Backend MakeCountTimeOfOpenStateInCloseTask", project), http.StatusBadRequest
	}

	logger.Info("Send MakeCountTimeOfResolvedStateInCloseTask request")
	err = repository.DbCon.GetRepository().MakeCountTimeOfResolvedStateInCloseTask(project)
	if err != nil {
		logger.Error("Something went wrong on MakeCountTimeOfResolvedStateInCloseTask request")
		return u.Message(false, err.Error(),
			"Jira Analyzer Backend MakeCountTimeOfResolvedStateInCloseTask", project), http.StatusBadRequest
	}

	logger.Info("Send MakeCountTimeOfReopenedStateInCloseTask request")
	err = repository.DbCon.GetRepository().MakeCountTimeOfReopenedStateInCloseTask(project)
	if err != nil {
		logger.Error("Something went wrong on MakeCountTimeOfReopenedStateInCloseTask request")
		return u.Message(false, err.Error(),
			"Jira Analyzer Backend MakeCountTimeOfReopenedStateInCloseTask", project), http.StatusBadRequest
	}

	logger.Info("Send MakeCountTimeOfInProgressStateInCloseTask request")
	err = repository.DbCon.GetRepository().MakeCountTimeOfInProgressStateInCloseTask(project)
	if err != nil {
		logger.Error("Something went wrong on MakeCountTimeOfInProgressStateInCloseTask request")
		return u.Message(false, err.Error(),
			"Jira Analyzer Backend MakeCountTimeOfInProgressStateInCloseTask", project), http.StatusBadRequest
	}

	resp := u.Message(true, "success",
		"Jira Analyzer Backend MakeTaskStateTime", project)
	logger.Info("Get result of MakeTaskStateTime request")
	return resp, http.StatusOK
}

func MakeActivityByTask(project string) (map[string]interface{}, int) {
	logger := u.GetLogger()
	logger.Info("Send GetReturnTimeSpentOnAllTasks request")

	logger.Info("Send MakeCountCloseTaskInDay request")
	err := repository.DbCon.GetRepository().MakeCountCloseTaskInDay(project)
	if err != nil {
		logger.Error("Something went wrong on MakeCountCloseTaskInDay request")
		return u.Message(false, err.Error(),
			"Jira Analyzer Backend MakeCountCloseTaskInDay", project), http.StatusBadRequest
	}

	logger.Info("Send MakeCountOpenTaskInDay request")
	err = repository.DbCon.GetRepository().MakeCountOpenTaskInDay(project)
	if err != nil {
		logger.Error("Something went wrong on MakeCountOpenTaskInDay request")
		return u.Message(false, err.Error(),
			"Jira Analyzer Backend MakeCountOpenTaskInDay", project), http.StatusBadRequest
	}
	resp := u.Message(true, "success",
		"Jira Analyzer Backend MakeActivityByTask", project)
	logger.Info("Get result of MakeActivityByTask request")
	return resp, http.StatusOK
}

func MakeTimeSpentOnAllTasks(project string) (map[string]interface{}, int) {
	logger := u.GetLogger()
	logger.Info("Send GetReturnTimeSpentOnAllTasks request")

	logger.Info("Send MakeTimeSpentOnAllTasks request")
	err := repository.DbCon.GetRepository().MakeTimeSpentOnAllTasks(project)
	if err != nil {
		logger.Error("Something went wrong on MakeTimeSpentOnAllTasks request")
		return u.Message(false, err.Error(),
			"Jira Analyzer Backend MakeTimeSpentOnAllTasks", project), http.StatusBadRequest
	}

	resp := u.Message(true, "success",
		"Jira Analyzer Backend MakeTimeSpentOnAllTasks", project)
	logger.Info("Get result of MakeTimeSpentOnAllTasks request")
	return resp, http.StatusOK
}

func MakePriorityCountOfProjectOpen(project string) (map[string]interface{}, int) {
	logger := u.GetLogger()
	logger.Info("Send GetReturnTimeSpentOnAllTasks request")

	logger.Info("Send MakePriorityCountOfProjectOpen request")
	err := repository.DbCon.GetRepository().MakePriorityCountOfProjectOpen(project)
	if err != nil {
		logger.Error("Something went wrong on MakePriorityCountOfProjectOpen request")
		return u.Message(false, err.Error(),
			"Jira Analyzer Backend MakePriorityCountOfProjectOpen", project), http.StatusBadRequest
	}
	resp := u.Message(true, "success",
		"Jira Analyzer Backend MakePriorityCountOfProjectOpen", project)
	logger.Info("Get result of MakePriorityCountOfProjectOpen request")
	return resp, http.StatusOK
}

func MakePriorityCountOfProjectClose(project string) (map[string]interface{}, int) {
	logger := u.GetLogger()
	logger.Info("Send GetReturnTimeSpentOnAllTasks request")

	logger.Info("Send MakePriorityCountOfProjectClose request")
	err := repository.DbCon.GetRepository().MakePriorityCountOfProjectClose(project)
	if err != nil {
		logger.Error("Something went wrong on MakePriorityCountOfProjectClose request")
		return u.Message(false, err.Error(),
			"Jira Analyzer Backend MakePriorityCountOfProjectClose", project), http.StatusBadRequest
	}
	resp := u.Message(true, "success",
		"Jira Analyzer Backend MakePriorityCountOfProjectClose", project)
	logger.Info("Get result of MakePriorityCountOfProjectClose request")
	return resp, http.StatusOK
}

func IsAnalyzedGraph(project string) (map[string]interface{}, int) {
	logger := u.GetLogger()
	logger.Info("Send IsAnalyzedGraph request")

	logger.Info("Send IsAnalyzed request")
	ok, err := repository.DbCon.GetRepository().IsAnalyzed(project)
	if err != nil {
		logger.Error("Something went wrong on IsAnalyzed request")
		return u.Message(false, err.Error(),
			"Jira Analyzer Backend IsAnalyzedGraph", project), http.StatusBadRequest
	}

	resp := u.Message(true, "success",
		"Jira Analyzer Backend IsAnalyzedGraph", project)

	data := make(map[string]interface{}, 1)
	data["isAnalyzed"] = ok
	resp["data"] = data
	logger.Info("Get result of IsAnalyzedGraph request")
	return resp, http.StatusOK
}

func IsEmptyProject(project string) (map[string]interface{}, int) {
	logger := u.GetLogger()
	logger.Info("Send IsEmptyProject request")

	logger.Info("Send IsEmpty request")
	ok, err := repository.DbCon.GetRepository().IsEmpty(project)
	if err != nil {
		logger.Error("Something went wrong on IsEmpty request")
		return u.Message(false, err.Error(),
			"Jira Analyzer Backend IsEmptyProject", project), http.StatusBadRequest
	}

	resp := u.Message(true, "success",
		"Jira Analyzer Backend IsEmptyProject", project)

	data := make(map[string]interface{}, 1)
	data["isEmpty"] = ok
	resp["data"] = data
	logger.Info("Get result of IsEmptyProject request")
	return resp, http.StatusOK
}

func DeleteGraphsByProject(project string) (map[string]interface{}, int) {
	logger := u.GetLogger()
	logger.Info("Send GetReturnTimeSpentOnAllTasks request")

	logger.Info("Send DeleteOpenTaskTime request")
	err := repository.DbCon.GetRepository().DeleteOpenTaskTime(project)
	if err != nil {
		logger.Error("Something went wrong on DeleteOpenTaskTime request")
		return u.Message(false, err.Error(),
			"Jira Analyzer Backend DeleteOpenTaskTime", project), http.StatusBadRequest
	}

	logger.Info("Send DeleteTaskStateTimeOpen request")
	err = repository.DbCon.GetRepository().DeleteTaskStateTimeOpen(project)
	if err != nil {
		logger.Error("Something went wrong on DeleteTaskStateTimeOpen request")
		return u.Message(false, err.Error(),
			"Jira Analyzer Backend DeleteTaskStateTimeOpen", project), http.StatusBadRequest
	}

	logger.Info("Send DeleteTaskStateTimeResolved request")
	err = repository.DbCon.GetRepository().DeleteTaskStateTimeResolved(project)
	if err != nil {
		logger.Error("Something went wrong on DeleteTaskStateTimeResolved request")
		return u.Message(false, err.Error(),
			"Jira Analyzer Backend DeleteTaskStateTimeResolved", project), http.StatusBadRequest
	}

	logger.Info("Send DeleteTaskStateTimeReopened request")
	err = repository.DbCon.GetRepository().DeleteTaskStateTimeReopened(project)
	if err != nil {
		logger.Error("Something went wrong on DeleteTaskStateTimeReopened request")
		return u.Message(false, err.Error(),
			"Jira Analyzer Backend DeleteTaskStateTimeReopened", project), http.StatusBadRequest
	}

	logger.Info("Send DeleteTaskStateTimeInProgress request")
	err = repository.DbCon.GetRepository().DeleteTaskStateTimeInProgress(project)
	if err != nil {
		logger.Error("Something went wrong on DeleteTaskStateTimeInProgress request")
		return u.Message(false, err.Error(),
			"Jira Analyzer Backend DeleteTaskStateTimeInProgress", project), http.StatusBadRequest
	}

	logger.Info("Send DeleteActivityByTaskOpen request")
	err = repository.DbCon.GetRepository().DeleteActivityByTaskOpen(project)
	if err != nil {
		logger.Error("Something went wrong on DeleteActivityByTaskOpen request")
		return u.Message(false, err.Error(),
			"Jira Analyzer Backend DeleteActivityByTaskOpen", project), http.StatusBadRequest
	}

	logger.Info("Send DeleteActivityByTaskClose request")
	err = repository.DbCon.GetRepository().DeleteActivityByTaskClose(project)
	if err != nil {
		logger.Error("Something went wrong on DeleteActivityByTaskClose request")
		return u.Message(false, err.Error(),
			"Jira Analyzer Backend DeleteActivityByTaskClose", project), http.StatusBadRequest
	}

	logger.Info("Send DeleteComplexityTaskTime request")
	err = repository.DbCon.GetRepository().DeleteComplexityTaskTime(project)
	if err != nil {
		logger.Error("Something went wrong on DeleteComplexityTaskTime request")
		return u.Message(false, err.Error(),
			"Jira Analyzer Backend DeleteComplexityTaskTime", project), http.StatusBadRequest
	}

	logger.Info("Send DeleteTaskPriorityCountOpen request")
	err = repository.DbCon.GetRepository().DeleteTaskPriorityCountOpen(project)
	if err != nil {
		logger.Error("Something went wrong on DeleteTaskPriorityCountOpen request")
		return u.Message(false, err.Error(),
			"Jira Analyzer Backend DeleteTaskPriorityCountOpen", project), http.StatusBadRequest
	}

	logger.Info("Send DeleteTaskPriorityCountClose request")
	err = repository.DbCon.GetRepository().DeleteTaskPriorityCountClose(project)
	if err != nil {
		logger.Error("Something went wrong on DeleteTaskPriorityCountClose request")
		return u.Message(false, err.Error(),
			"Jira Analyzer Backend DeleteTaskPriorityCountClose", project), http.StatusBadRequest
	}

	resp := u.Message(true, "success",
		"Jira Analyzer Backend DeleteGraphsByProject", project)
	logger.Info("Get result of DeleteGraphsByProject request")
	return resp, http.StatusOK
}
