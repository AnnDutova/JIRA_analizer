package controllers

import (
	"Backend/pkg/services"
	u "Backend/pkg/utils"
	"log"
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
		log.Print("Graph 2")
	case 3:
		log.Print("Graph 3")
	case 4:
		getForthGraphGroup(w, project)
	case 5:
		log.Print("Graph 5 Priority")
	case 6:
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

func getForthGraphGroup(w http.ResponseWriter, project string) {
	var data []any
	resp2, status := services.GetReturnTimeSpentOnAllTasks(project)
	if status != http.StatusOK {
		u.RespondAny(w, nil, http.StatusInternalServerError)
	}
	data = append(data, resp2)
	u.RespondVariable(w, http.StatusOK, data...)
}
