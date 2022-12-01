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

	if limit != "" {
		iLimit, err = strconv.ParseInt(limit, 10, 64)
		if err != nil {
			return u.Message(false, "Limit parameter must be integer",
				"Jira Analyzer Backend Get Projects", url), http.StatusBadRequest
		}
	} else {
		iLimit = -1
	}

	if page != "" {
		uPage, err = strconv.ParseUint(page, 10, 64)
		if err != nil {
			return u.Message(false, "Page parameter must be integer",
				"Jira Analyzer Backend Get Projects", url), http.StatusBadRequest
		}
	} else {
		uPage = 1
	}

	projects, pageCount, err := repository.DbCon.GetRepository().ReturnAllProjects(iLimit, uPage, search)
	if err != nil {
		return u.Message(false, err.Error(),
			"Jira Analyzer Backend Get Projects", url), http.StatusBadRequest
	}

	resp := u.Message(true, "success",
		"Jira Analyzer Backend Get Projects", url)

	resp["pageCount"] = pageCount
	resp["data"] = projects

	return resp, http.StatusOK
}

func GetProjectAnalytic(id string, url string) (map[string]interface{}, int) {
	projectStat, err := repository.DbCon.GetRepository().ReturnProjectAnalytic(id)
	if err != nil {
		return u.Message(false, err.Error(),
			"Jira Analyzer Backend Get Project", url), http.StatusBadRequest
	}

	resp := u.Message(true, "success",
		"Jira Analyzer Backend Get Project", url)

	resp["data"] = projectStat

	return resp, http.StatusOK
}
