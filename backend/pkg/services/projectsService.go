package services

import (
	"Backend/pkg/repository"
	u "Backend/pkg/utils"
	"net/http"
	"strconv"
)

func GetProjects(url, limit, page, search string) (map[string]interface{}, int) {
	var uPage uint64
	var iLimit int64
	var err error
	logger := u.GetLogger()
	logger.Info("Get GetProjects request")

	if limit != "" {
		iLimit, err = strconv.ParseInt(limit, 10, 64)
		if err != nil {
			logger.Error("Limit parameter must be integer ", err.Error())
			return u.Message(false, "Limit parameter must be integer",
				"Jira Analyzer Backend Get Projects", url), http.StatusBadRequest
		}
	} else {
		iLimit = -1
	}

	if page != "" {
		uPage, err = strconv.ParseUint(page, 10, 64)
		if err != nil {
			logger.Error("Page parameter must be integer ", err.Error())
			return u.Message(false, "Page parameter must be integer",
				"Jira Analyzer Backend Get Projects", url), http.StatusBadRequest
		}
	} else {
		uPage = 1
	}

	logger.Info("Send request on ReturnAllProjects")
	projects, pageCount, err := repository.DbCon.GetRepository().ReturnAllProjects(iLimit, uPage, search)
	if err != nil {
		logger.Error("Something went wrong on ReturnAllProjects(iLimit, uPage, search) ", err.Error())
		return u.Message(false, err.Error(),
			"Jira Analyzer Backend Get Projects", url), http.StatusBadRequest
	}
	logger.Info("Get result of request on ReturnAllProjects")

	resp := u.Message(true, "success",
		"Jira Analyzer Backend Get Projects", url)

	resp["pageCount"] = pageCount
	resp["data"] = projects

	return resp, http.StatusOK
}

func GetProjectAnalytic(id string, url string) (map[string]interface{}, int) {
	logger := u.GetLogger()
	logger.Info("Get GetProjectAnalytic request")
	projectStat, err := repository.DbCon.GetRepository().ReturnProjectAnalytic(id)
	if err != nil {
		logger.Info("Something went wrong on GetProjectAnalytic ", err.Error())
		return u.Message(false, err.Error(),
			"Jira Analyzer Backend Get Project", url), http.StatusBadRequest
	}
	logger.Info("Get result of GetProjectAnalytic request")

	resp := u.Message(true, "success",
		"Jira Analyzer Backend Get Project", url)

	resp["data"] = projectStat

	return resp, http.StatusOK
}

func DeleteProjectById(id string, url string) (map[string]interface{}, int) {
	logger := u.GetLogger()
	logger.Info("Get DeleteProjectById request")
	project, err := repository.DbCon.GetRepository().DeleteProjectById(id)
	if err != nil {
		logger.Error("Something went wrong on DeleteProjectById ", err.Error())
		return u.Message(false, err.Error(),
			"Jira Analyzer Backend DeleteProject", url), http.StatusBadRequest
	}
	logger.Info("Get result of DeleteProjectById request")

	resp := u.Message(true, "success",
		"Jira Analyzer Backend DeleteProject", url)

	resp["data"] = project

	return resp, http.StatusOK
}
