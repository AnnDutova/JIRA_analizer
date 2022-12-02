package controllers

import (
	"Backend/pkg/services"
	u "Backend/pkg/utils"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

var GetCompareByGraphGroup = func(w http.ResponseWriter, r *http.Request) {
	projects := r.FormValue("project")
	re, _ := regexp.Compile("compare/(.*)")
	group, err := strconv.Atoi(strings.Split(re.FindString(r.URL.Path), "/")[1])
	project := strings.Split(projects, ",")

	if err != nil {
		u.RespondAny(w, nil, http.StatusInternalServerError)
	}

	switch group {
	case 1:
		getFirstGraph(w, project)
	case 2:
		getSecondGraph(w, project)
	case 3:
		getThirdGraph(w, project)
	case 4:
		getForthGraph(w, project)
	case 5:
		getFifthGraph(w, project)
	case 6:
		getSixthGraph(w, project)
	default:
		u.RespondAny(w, nil, http.StatusBadGateway)
	}
}

func getFirstGraph(w http.ResponseWriter, project []string) {
	data, status := services.GetFirstGraphOnCompare(project)
	if status != http.StatusOK {
		u.RespondAny(w, nil, http.StatusInternalServerError)
	}
	u.RespondAny(w, data, http.StatusOK)
}

func getSecondGraph(w http.ResponseWriter, project []string) {
	data, status := services.GetSecondGraphOnCompare(project)
	if status != http.StatusOK {
		u.RespondAny(w, nil, http.StatusInternalServerError)
	}
	u.RespondAny(w, data, http.StatusOK)
}

func getThirdGraph(w http.ResponseWriter, project []string) {
	data, status := services.GetThirdGraphOnCompare(project)
	if status != http.StatusOK {
		u.RespondAny(w, nil, http.StatusInternalServerError)
	}
	u.RespondAny(w, data, http.StatusOK)
}

func getForthGraph(w http.ResponseWriter, project []string) {
	data, status := services.GetForthGraphOnCompare(project)
	if status != http.StatusOK {
		u.RespondAny(w, nil, http.StatusInternalServerError)
	}
	u.RespondAny(w, data, http.StatusOK)
}

func getFifthGraph(w http.ResponseWriter, project []string) {
	data, status := services.GetFifthGraphOnCompare(project)
	if status != http.StatusOK {
		u.RespondAny(w, nil, http.StatusInternalServerError)
	}
	u.RespondAny(w, data, http.StatusOK)
}

func getSixthGraph(w http.ResponseWriter, project []string) {
	data, status := services.GetSixthGraphOnCompare(project)
	if status != http.StatusOK {
		u.RespondAny(w, nil, http.StatusInternalServerError)
	}
	u.RespondAny(w, data, http.StatusOK)
}
