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
	re, _ := regexp.Compile("graph/get/(.*)")
	group, err := strconv.Atoi(strings.Split(re.FindString(r.URL.Path), "/")[2])

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
	default:
		u.RespondAny(w, nil, http.StatusBadGateway)
	}
}

func getFirstGraphGroup(w http.ResponseWriter, project string) {
	resp, status := services.GetReturnTimeCountOfIssuesInCloseState(project)
	if status > 300 && status < 200 {
		u.RespondAny(w, nil, http.StatusInternalServerError)
	}
	u.RespondAny(w, resp, http.StatusOK)
}

func getSecondGraphGroup(w http.ResponseWriter, project string) {
	resp, status := services.GetReturnTaskStateTime(project)
	if status > 300 && status < 200 {
		u.RespondAny(w, nil, http.StatusInternalServerError)
	}
	u.RespondAny(w, resp, http.StatusOK)
}

func getThirdGraphGroup(w http.ResponseWriter, project string) {
	data, status := services.GetReturnActivityByTask(project)
	if status > 300 && status < 200 {
		u.RespondAny(w, nil, http.StatusInternalServerError)
	}
	u.RespondAny(w, data, http.StatusOK)
}

func getForthGraphGroup(w http.ResponseWriter, project string) {
	resp, status := services.GetReturnTimeSpentOnAllTasks(project)
	if status > 300 && status < 200 {
		u.RespondAny(w, nil, http.StatusInternalServerError)
	}
	u.RespondAny(w, resp, http.StatusOK)
}

func getFifthGraphGroup(w http.ResponseWriter, project string) {
	resp, status := services.GetReturnPriorityCountOfProjectOpen(project)
	if status > 300 && status < 200 {
		u.RespondAny(w, nil, http.StatusInternalServerError)
	}
	u.RespondAny(w, resp, http.StatusOK)
}

func getSixGraphGroup(w http.ResponseWriter, project string) {
	resp, status := services.GetReturnPriorityCountOfProjectClose(project)
	if status > 300 && status < 200 {
		u.RespondAny(w, nil, http.StatusInternalServerError)
	}
	u.RespondAny(w, resp, http.StatusOK)
}

var MakeGraphByGroup = func(w http.ResponseWriter, r *http.Request) {
	project := r.FormValue("project")
	re, _ := regexp.Compile("graph/make/(.*)")
	group, err := strconv.Atoi(strings.Split(re.FindString(r.URL.Path), "/")[2])

	if err != nil {
		u.RespondAny(w, nil, http.StatusInternalServerError)
	}

	switch group {
	case 1:
		makeFirstGraphGroup(w, project)
	case 2:
		makeSecondGraphGroup(w, project)
	case 3:
		makeThirdGraphGroup(w, project)
	case 4:
		makeForthGraphGroup(w, project)
	case 5:
		makeFifthGraphGroup(w, project)
	case 6:
		makeSixGraphGroup(w, project)
	default:
		u.RespondAny(w, nil, http.StatusBadGateway)
	}
}

func makeFirstGraphGroup(w http.ResponseWriter, project string) {
	resp, status := services.MakeTimeCountOfIssuesInCloseState(project)
	if status != http.StatusOK {
		u.RespondAny(w, nil, http.StatusInternalServerError)
	}
	u.RespondAny(w, resp, http.StatusOK)
}

func makeSecondGraphGroup(w http.ResponseWriter, project string) {
	resp, status := services.MakeTaskStateTime(project)
	if status != http.StatusOK {
		u.RespondAny(w, nil, http.StatusInternalServerError)
	}
	u.RespondAny(w, resp, http.StatusOK)
}

func makeThirdGraphGroup(w http.ResponseWriter, project string) {
	data, status := services.MakeActivityByTask(project)
	if status != http.StatusOK {
		u.RespondAny(w, nil, http.StatusInternalServerError)
	}
	u.RespondAny(w, data, http.StatusOK)
}

func makeForthGraphGroup(w http.ResponseWriter, project string) {
	resp, status := services.MakeTimeSpentOnAllTasks(project)
	if status != http.StatusOK {
		u.RespondAny(w, nil, http.StatusInternalServerError)
	}
	u.RespondAny(w, resp, http.StatusOK)
}

func makeFifthGraphGroup(w http.ResponseWriter, project string) {
	resp, status := services.MakePriorityCountOfProjectOpen(project)
	if status != http.StatusOK {
		u.RespondAny(w, nil, http.StatusInternalServerError)
	}
	u.RespondAny(w, resp, http.StatusOK)
}

func makeSixGraphGroup(w http.ResponseWriter, project string) {
	resp, status := services.MakePriorityCountOfProjectClose(project)
	if status != http.StatusOK {
		u.RespondAny(w, nil, http.StatusInternalServerError)
	}
	u.RespondAny(w, resp, http.StatusOK)
}

var IsAnalyzed = func(w http.ResponseWriter, r *http.Request) {
	project := r.FormValue("project")
	resp, status := services.IsAnalyzedGraph(project)
	if status != http.StatusOK {
		u.RespondAny(w, nil, http.StatusInternalServerError)
	}
	u.RespondAny(w, resp, http.StatusOK)
}

var DeleteGraphByProject = func(w http.ResponseWriter, r *http.Request) {
	project := r.FormValue("project")

	resp, status := services.DeleteGraphsByProject(project)
	if status != http.StatusOK {
		u.RespondAny(w, nil, http.StatusInternalServerError)
	}
	u.RespondAny(w, resp, http.StatusOK)
}

var OptionsReq = func(w http.ResponseWriter, r *http.Request) {
	u.RespondAny(w, nil, http.StatusOK)
}
