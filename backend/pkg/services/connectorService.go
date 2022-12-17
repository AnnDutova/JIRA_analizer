package services

import (
	"Backend/pkg/repository"
	u "Backend/pkg/utils"
	"net/http"
)

func AddProjectToDB(url string, project string) (map[string]interface{}, int) {
	logger := u.GetLogger()
	logger.Info("Send requests to connector: Get AddProjectToDB request")

	projects, err := repository.DbCon.GetRepository().AddProjectToDB(project)
	if err != nil {
		logger.Error("Something went wrong on AddProjectToDB ", err.Error())
		return u.Message(false, err.Error(),
			"Jira Analyzer Backend Get Projects", url), http.StatusBadRequest
	}
	logger.Info("Send requests to connector: Get result of AddProjectToDB request")
	resp := u.Message(true, "success",
		"Jira Analyzer Backend Get Projects", url)

	resp["data"] = projects

	return resp, http.StatusOK
}

func ReturnAllProjectsFromConnector(url string, limit, page, search string) (map[string]interface{}, int) {
	logger := u.GetLogger()
	logger.Info("Send requests to connector: Get ReturnAllProjectsFromConnector request")

	project, err := repository.DbCon.GetRepository().ReturnAllProjectsFromConnector(limit, page, search)
	if err != nil {
		logger.Error("Something went wrong on ReturnAllProjectsFromConnector")
		return u.Message(false, err.Error(),
			"Jira Analyzer Backend Get Projects", url), http.StatusBadRequest
	}

	logger.Info("Get result from connector: Get ReturnAllProjectsFromConnector request")
	resp := u.Message(true, "success",
		"Jira Analyzer Backend Get Projects", url)

	resp["data"] = project.Projects
	resp["pageCount"] = project.PageInfo.PageCount
	resp["pageInfo"] = project.PageInfo

	return resp, http.StatusOK
}
