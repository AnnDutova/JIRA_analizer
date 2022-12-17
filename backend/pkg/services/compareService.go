package services

import (
	"Backend/pkg/repository"
	u "Backend/pkg/utils"
	"fmt"
	"net/http"
)

func GetFirstGraphOnCompare(projects []string) (map[string]interface{}, int) {
	logger := u.GetLogger()
	logger.Info("Sent GetFirstGraphOnCompare request")

	data := make(map[string]interface{}, 0)
	count := make(map[string]interface{}, 0)
	var emptyProject string
	for ipr, project := range projects {
		logger.Info("Sent CheckExistenceOnOpenTaskTimeTable request for project ", project)
		issues, err := repository.DbCon.GetRepository().CheckExistenceOnOpenTaskTimeTable(project)
		if err != nil {
			logger.Error("Something went wrong on CheckExistenceOnOpenTaskTimeTable ", err.Error())
			return u.Message(false, err.Error(),
				"Jira Analyzer Backend CheckExistenceOnOpenTaskTimeTable", project), http.StatusBadRequest
		}
		if issues == nil {
			emptyProject = project
			break
		}
		for _, el := range issues {
			if val, ok := count[el.Title]; ok {
				val.([]int)[ipr] = el.Count
				count[el.Title] = val
			} else {
				arr := make([]int, len(projects))
				arr[ipr] = el.Count
				count[el.Title] = arr
			}
		}
	}
	var resp map[string]interface{}
	if len(emptyProject) == 0 {
		data["count"] = count
		data["categories"] = u.SortCategories(count)
		resp = u.Message(true, "success",
			"Jira Analyzer Backend GetFirstGraphOnCompare", fmt.Sprint(projects))
		resp["data"] = data
	} else {
		data["count"] = nil
		data["categories"] = nil
		data["projects"] = projects
		resp = u.Message(true, "accepted",
			fmt.Sprintf("Jira Analyzer Backend GetFirstGraphOnCompare. Has empty data for %s", emptyProject), fmt.Sprint(projects))
		resp["data"] = data

	}
	logger.Info("Get result GetFirstGraphOnCompare request")
	return resp, http.StatusOK

}

func GetSecondGraphOnCompare(projects []string) (map[string]interface{}, int) {
	logger := u.GetLogger()
	logger.Info("Sent GetFirstGraphOnCompare request")

	data := make(map[string]interface{}, 0)
	open := make(map[string]interface{}, 0)
	resolve := make(map[string]interface{}, 0)
	reopen := make(map[string]interface{}, 0)
	progress := make(map[string]interface{}, 0)
	var emptyProject string

	for ipr, project := range projects {
		logger.Info("Sent CheckExistenceOnTaskStateTimeTableOpen request for project ", project)
		openTask, err := repository.DbCon.GetRepository().CheckExistenceOnTaskStateTimeTableOpen(project)
		if err != nil {
			logger.Error("Something went wrong on CheckExistenceOnTaskStateTimeTableOpen ", err.Error())
			return u.Message(false, err.Error(),
				"Jira Analyzer Backend CheckExistenceOnTaskStateTimeTableOpen", project), http.StatusBadRequest
		}
		logger.Info("Sent CheckExistenceOnTaskStateTimeTableResolved request for project ", project)
		resolvedTask, err := repository.DbCon.GetRepository().CheckExistenceOnTaskStateTimeTableResolved(project)
		if err != nil {
			logger.Error("Something went wrong on CheckExistenceOnTaskStateTimeTableResolved", err.Error())
			return u.Message(false, err.Error(),
				"Jira Analyzer Backend CheckExistenceOnTaskStateTimeTableResolved", project), http.StatusBadRequest
		}
		logger.Info("Sent CheckExistenceOnTaskStateTimeTableReopened request for project ", project)
		reopenedTask, err := repository.DbCon.GetRepository().CheckExistenceOnTaskStateTimeTableReopened(project)
		if err != nil {
			logger.Error("Something went wrong on CheckExistenceOnTaskStateTimeTableReopened ", err.Error())
			return u.Message(false, err.Error(),
				"Jira Analyzer Backend CheckExistenceOnTaskStateTimeTableReopened", project), http.StatusBadRequest
		}
		logger.Info("Sent CheckExistenceOnTaskStateTimeTableInProgress request for project ", project)
		progressTask, err := repository.DbCon.GetRepository().CheckExistenceOnTaskStateTimeTableInProgress(project)
		if err != nil {
			logger.Error("Something went wrong on CheckExistenceOnTaskStateTimeTableInProgress ", err.Error())
			return u.Message(false, err.Error(),
				"Jira Analyzer Backend CheckExistenceOnTaskStateTimeTableInProgress", project), http.StatusBadRequest
		}

		if openTask == nil && resolvedTask == nil && reopenedTask == nil && progressTask == nil {
			emptyProject = project
			break
		}

		open = u.CreateResultMap(len(projects), ipr, openTask, open)
		resolve = u.CreateResultMap(len(projects), ipr, resolvedTask, resolve)
		reopen = u.CreateResultMap(len(projects), ipr, reopenedTask, reopen)
		progress = u.CreateResultMap(len(projects), ipr, progressTask, progress)
	}

	var resp map[string]interface{}
	if len(emptyProject) == 0 {
		category := make(map[string]interface{}, 0)
		if len(open) > 0 {
			data["open"] = open
			category["open"] = u.SortCategories(open)
		} else {
			data["open"] = nil
			category["open"] = nil
		}
		if len(reopen) > 0 {
			data["reopen"] = reopen
			category["reopen"] = u.SortCategories(reopen)
		} else {
			data["reopen"] = nil
			category["reopen"] = nil
		}
		if len(progress) > 0 {
			data["progress"] = progress
			category["progress"] = u.SortCategories(progress)
		} else {
			data["progress"] = nil
			category["progress"] = nil
		}
		if len(resolve) > 0 {
			data["resolve"] = resolve
			category["resolve"] = u.SortCategories(resolve)
		} else {
			data["resolve"] = nil
			category["resolve"] = nil
		}
		data["projects"] = projects
		data["categories"] = category
		resp = u.Message(true, "success",
			"Jira Analyzer Backend GetSecondGraphOnCompare", fmt.Sprint(projects))
		resp["data"] = data
	} else {
		data["open"] = nil
		data["reopen"] = nil
		data["progress"] = nil
		data["resolve"] = nil
		data["categories"] = nil
		data["projects"] = projects
		resp = u.Message(true, "accepted",
			fmt.Sprintf("Jira Analyzer Backend GetSecondGraphOnCompare. Has empty data for %s", emptyProject), fmt.Sprint(projects))
		resp["data"] = data
	}
	logger.Info("Get result of GetFirstGraphOnCompare request")
	return resp, http.StatusOK
}

func GetThirdGraphOnCompare(projects []string) (map[string]interface{}, int) {
	logger := u.GetLogger()
	logger.Info("Sent GetThirdGraphOnCompare request")

	data := make(map[string]interface{}, 0)
	open := make(map[string]interface{}, 0)
	cls := make(map[string]interface{}, 0)
	all := make(map[string]interface{}, 0)

	var emptyProject string

	for ipr, project := range projects {
		logger.Info("Sent CheckExistenceOnActivityByTaskTableClose request for project ", project)
		closeTask, err := repository.DbCon.GetRepository().CheckExistenceOnActivityByTaskTableClose(project)
		if err != nil {
			logger.Error("Send result from CheckExistenceOnActivityByTaskTableClose ", err.Error())
			return u.Message(false, err.Error(),
				"Jira Analyzer Backend CheckExistenceOnActivityByTaskTableClose", project), http.StatusBadRequest
		}
		logger.Info("Sent CheckExistenceOnActivityByTaskTableOpen request for project ", project)
		openTask, err := repository.DbCon.GetRepository().CheckExistenceOnActivityByTaskTableOpen(project)
		if err != nil {
			logger.Error("Send result from CheckExistenceOnActivityByTaskTableOpen ", err.Error())
			return u.Message(false, err.Error(),
				"Jira Analyzer Backend CheckExistenceOnActivityByTaskTableOpen", project), http.StatusBadRequest
		}

		if openTask == nil && closeTask == nil {
			emptyProject = project
			break
		}

		open = u.CreateResultMap(len(projects), ipr, openTask, open)
		cls = u.CreateResultMap(len(projects), ipr, closeTask, cls)
		all = u.JoinToMap(openTask, closeTask, all)
	}
	var resp map[string]interface{}
	if len(emptyProject) == 0 {
		category := make(map[string]interface{}, 0)
		if len(open) > 0 {
			data["open"] = open
			openDates, err := u.SortDatesForActivityGraph(open)
			if err != nil {
				return u.Message(false, err.Error(),
					"Jira Analyzer Backend CheckExistenceOnActivityByTaskTableOpen. Fail on SortDatesForActivityGraph",
					"Dates for open tasks"), http.StatusBadRequest
			}
			category["open"] = openDates
		} else {
			data["open"] = nil
			category["open"] = nil
		}
		if len(cls) > 0 {
			data["close"] = cls
			closeDates, err := u.SortDatesForActivityGraph(cls)
			if err != nil {
				return u.Message(false, err.Error(),
					"Jira Analyzer Backend CheckExistenceOnActivityByTaskTableOpen. Fail on SortDatesForActivityGraph",
					"Dates for close tasks"), http.StatusBadRequest
			}
			category["close"] = closeDates
		} else {
			data["close"] = nil
			category["close"] = nil
		}
		if len(all) > 0 {
			allCategories, err := u.SortDatesForActivityGraph(all)
			if err != nil {
				return u.Message(false, err.Error(),
					"Jira Analyzer Backend CheckExistenceOnActivityByTaskTableOpen. Fail on SortDatesForActivityGraph",
					"Dates for all categories"), http.StatusBadRequest
			}
			category["all"] = allCategories
		} else {
			category["all"] = nil
		}
		resp = u.Message(true, "success",
			"Jira Analyzer Backend GetThirdGraphOnCompare", fmt.Sprint(projects))
		data["projects"] = projects
		data["categories"] = category
		resp["data"] = data
	} else {
		data["open"] = nil
		data["close"] = nil
		resp = u.Message(true, "accepted",
			fmt.Sprintf("Jira Analyzer Backend GetThirdGraphOnCompare. Has empty data for %s", emptyProject), fmt.Sprint(projects))
		data["projects"] = projects
		data["categories"] = nil
		resp["data"] = data
	}
	logger.Info("Get result of GetThirdGraphOnCompare request")
	return resp, http.StatusOK
}

func GetForthGraphOnCompare(projects []string) (map[string]interface{}, int) {
	logger := u.GetLogger()
	logger.Info("Sent GetForthGraphOnCompare request")

	data := make(map[string]interface{}, 0)
	category := make(map[string]interface{}, 0)
	var emptyProject string
	for ipr, project := range projects {
		logger.Info("Sent CheckExistenceOnComplexityTaskTimeTable request", project)
		issues, err := repository.DbCon.GetRepository().CheckExistenceOnComplexityTaskTimeTable(project)
		if err != nil {
			logger.Info("Something went wrong CheckExistenceOnComplexityTaskTimeTable request", err.Error())
			return u.Message(false, err.Error(),
				"Jira Analyzer Backend CheckExistenceOnComplexityTaskTimeTable", project), http.StatusBadRequest
		}
		if issues == nil {
			emptyProject = project
			break
		}
		category = u.CreateResultMap(len(projects), ipr, issues, category)
	}

	var resp map[string]interface{}
	if len(emptyProject) == 0 {
		resp = u.Message(true, "success",
			"Jira Analyzer Backend GetForthGraphOnCompare", fmt.Sprint(projects))
		data["projects"] = projects
		data["complexity"] = category
		data["categories"] = u.SortMinutesCategories(category)
		resp["data"] = data
	} else {
		resp = u.Message(true, "accepted",
			fmt.Sprintf("Jira Analyzer Backend GetForthGraphOnCompare. Has empty data for %s", emptyProject), fmt.Sprint(projects))
		data["projects"] = projects
		data["categories"] = nil
		resp["data"] = data
	}
	logger.Info("Get result of GetForthGraphOnCompare request")
	return resp, http.StatusOK
}

func GetFifthGraphOnCompare(projects []string) (map[string]interface{}, int) {
	logger := u.GetLogger()
	logger.Info("Sent GetFifthGraphOnCompare request")

	data := make(map[string]interface{}, 0)
	category := make(map[string]interface{}, 0)
	var emptyProject string

	for ipr, project := range projects {
		logger.Info("Sent CheckExistenceOnTaskPriorityCountTableOpen request", project)
		issues, err := repository.DbCon.GetRepository().CheckExistenceOnTaskPriorityCountTableOpen(project)
		if err != nil {
			logger.Info("Something went wrong CheckExistenceOnTaskPriorityCountTableOpen request", err.Error())
			return u.Message(false, err.Error(),
				"Jira Analyzer Backend CheckExistenceOnTaskPriorityCountTable", project), http.StatusBadRequest
		}
		if issues == nil {
			emptyProject = project
			break
		}
		category = u.CreateResultMap(len(projects), ipr, issues, category)
	}

	var resp map[string]interface{}
	if len(emptyProject) == 0 {
		resp = u.Message(true, "success",
			"Jira Analyzer REST API GetFifthGraphOnCompare", fmt.Sprint(projects))
		data["projects"] = projects
		keys := make([]string, 0, len(category))
		for k := range category {
			keys = append(keys, k)
		}
		data["priority"] = category
		data["categories"] = keys
		resp["data"] = data
	} else {
		data["priority"] = nil
		resp = u.Message(true, "accepted",
			fmt.Sprintf("Jira Analyzer Backend GetFifthGraphOnCompare. Has empty data for %s", emptyProject), fmt.Sprint(projects))
		data["projects"] = projects
		data["categories"] = nil
		resp["data"] = data
	}
	logger.Info("Get result of GetFifthGraphOnCompare request")
	return resp, http.StatusOK
}

func GetSixthGraphOnCompare(projects []string) (map[string]interface{}, int) {
	logger := u.GetLogger()
	logger.Info("Sent GetSixthGraphOnCompare request")

	data := make(map[string]interface{}, 0)
	category := make(map[string]interface{}, 0)
	var emptyProject string
	for ipr, project := range projects {
		logger.Info("Sent CheckExistenceOnTaskPriorityCountTableClose request", project)
		issues, err := repository.DbCon.GetRepository().CheckExistenceOnTaskPriorityCountTableClose(project)
		if err != nil {
			logger.Info("Something wrong CheckExistenceOnTaskPriorityCountTableClose request", err.Error())
			return u.Message(false, err.Error(),
				"Jira Analyzer Backend CheckExistenceOnTaskPriorityCountTable", project), http.StatusBadRequest
		}
		if issues == nil {
			emptyProject = project
			break
		}
		category = u.CreateResultMap(len(projects), ipr, issues, category)
	}

	var resp map[string]interface{}
	if len(emptyProject) == 0 {
		resp = u.Message(true, "success",
			"Jira Analyzer Backend GetSixthGraphOnCompare", fmt.Sprint(projects))
		data["projects"] = projects
		keys := make([]string, 0, len(category))
		for k := range category {
			keys = append(keys, k)
		}
		data["priority"] = category
		data["categories"] = keys
		resp["data"] = data
	} else {
		data["priority"] = nil
		resp = u.Message(true, "accepted",
			fmt.Sprintf("Jira Analyzer Backend GetSixthGraphOnCompare. Has empty data for %s", emptyProject), fmt.Sprint(projects))
		data["projects"] = projects
		data["categories"] = nil
		resp["data"] = data
	}
	logger.Info("Get result of GetSixthGraphOnCompare request")
	return resp, http.StatusOK
}
