package services

import (
	"Backend/pkg/repository"
	u "Backend/pkg/utils"
	"net/http"
)

func AddProjectToDB(url string, project string) (map[string]interface{}, int) {

	projects, err := repository.DbCon.GetRepository().AddProjectToDB(project)
	if err != nil {
		return u.Message(false, err.Error(),
			"Jira Analyzer Backend Get Projects", url), http.StatusBadRequest
	}

	resp := u.Message(true, "success",
		"Jira Analyzer Backend Get Projects", url)

	resp["data"] = projects

	return resp, http.StatusOK
}

func ReturnAllProjectsFromConnector(url string, limit, page, search string) (map[string]interface{}, int) {

	project, err := repository.DbCon.GetRepository().ReturnAllProjectsFromConnector(limit, page, search)
	if err != nil {
		return u.Message(false, err.Error(),
			"Jira Analyzer Backend Get Projects", url), http.StatusBadRequest
	}

	resp := u.Message(true, "success",
		"Jira Analyzer Backend Get Projects", url)

	resp["data"] = project.Projects
	resp["pageCount"] = project.PageInfo.PageCount
	resp["pageInfo"] = project.PageInfo

	return resp, http.StatusOK
}
