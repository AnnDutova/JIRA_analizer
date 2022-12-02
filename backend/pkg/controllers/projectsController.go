package controllers

import (
	"Backend/pkg/services"
	u "Backend/pkg/utils"
	"net/http"
	"regexp"
	"strings"
)

var GetProjectsFor = func(w http.ResponseWriter, r *http.Request) {
	limit := r.FormValue("limit")
	page := r.FormValue("page")
	search := r.FormValue("search")

	resp, status := services.GetProjects(r.URL.Path, limit, page, search)

	u.RespondAny(w, resp, status)
}

var GetProjectAnalytic = func(w http.ResponseWriter, r *http.Request) {
	re, _ := regexp.Compile("projects/(.*)")
	id := re.FindString(r.URL.Path)

	resp, status := services.GetProjectAnalytic(strings.Split(id, "/")[1], r.URL.Path)

	u.RespondAny(w, resp, status)
}
