package controllers

import (
	"Backend/pkg/services"
	u "Backend/pkg/utils"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

var GetGraphByGroup = func(w http.ResponseWriter, r *http.Request) {
	project := r.FormValue("project")
	re, _ := regexp.Compile("graph/(.*)")
	group, err := strconv.Atoi(strings.Split(re.FindString(r.URL.Path), "/")[1])

	if err != nil {
		u.RespondAny(w, nil, http.StatusInternalServerError)
	}

	switch group {
	case 1:
		getFirstGraphGroup(w, project)
	case 2:
		getSecondGraphGroup(w, project)
	case 3:
		getThirdGraphGroup(w, project)
	case 4:
		getForthGraphGroup(w, project)
	case 5:
		getFifthGraphGroup(w, project)
	case 6:
		getSixGraphGroup(w, project)
	case 7:
		resp, status := services.GetReturnTheMostActiveCreators(project)
		u.RespondAny(w, resp, status)
	}

}

func getFirstGraphGroup(w http.ResponseWriter, project string) {
	var data []any
	resp2, status := services.GetReturnTimeCountOfIssuesInCloseState(project)
	if status != http.StatusOK {
		u.RespondAny(w, nil, http.StatusInternalServerError)
	}
	data = append(data, resp2)
	u.RespondVariable(w, http.StatusOK, data...)
}

func getSecondGraphGroup(w http.ResponseWriter, project string) {
	resp, status := services.GetReturnTaskStateTime(project)
	if status != http.StatusOK {
		u.RespondAny(w, nil, http.StatusInternalServerError)
	}
	u.RespondAny(w, resp, http.StatusOK)
}

func getThirdGraphGroup(w http.ResponseWriter, project string) {
	data, status := services.GetReturnActivityByTask(project)
	if status != http.StatusOK {
		u.RespondAny(w, nil, http.StatusInternalServerError)
	}
	u.RespondAny(w, data, http.StatusOK)
}

func getForthGraphGroup(w http.ResponseWriter, project string) {
	var data []any
	resp2, status := services.GetReturnTimeSpentOnAllTasks(project)
	if status != http.StatusOK {
		u.RespondAny(w, nil, http.StatusInternalServerError)
	}
	data = append(data, resp2)
	u.RespondVariable(w, http.StatusOK, data...)
}

func getFifthGraphGroup(w http.ResponseWriter, project string) {
	var data []any
	resp, status := services.GetReturnPriorityCountOfProjectOpen(project)
	if status != http.StatusOK {
		u.RespondAny(w, nil, http.StatusInternalServerError)
	}
	data = append(data, resp)
	u.RespondVariable(w, http.StatusOK, data...)
}

func getSixGraphGroup(w http.ResponseWriter, project string) {
	var data []any
	resp, status := services.GetReturnPriorityCountOfProjectClose(project)
	if status != http.StatusOK {
		u.RespondAny(w, nil, http.StatusInternalServerError)
	}
	data = append(data, resp)
	u.RespondVariable(w, http.StatusOK, data...)
}
