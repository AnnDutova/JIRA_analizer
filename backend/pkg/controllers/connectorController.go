package controllers

import (
	"Backend/pkg/services"
	u "Backend/pkg/utils"
	"net/http"
)

var AddProjectToDB = func(w http.ResponseWriter, r *http.Request) {
	project := r.FormValue("project")

	resp, status := services.AddProjectToDB(r.URL.Path, project)

	u.RespondAny(w, resp, status)
}

var GetAllProjectsFromConnector = func(w http.ResponseWriter, r *http.Request) {
	limit := r.FormValue("limit")
	page := r.FormValue("page")
	search := r.FormValue("search")

	resp, status := services.ReturnAllProjectsFromConnector(r.URL.Path, limit, page, search)

	u.RespondAny(w, resp, status)
}
