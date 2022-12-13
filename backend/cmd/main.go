package main

import (
	"Backend/pkg/controllers"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/v1/projects",
		controllers.GetProjectsFor).Methods("GET")

	router.HandleFunc("/api/v1/projects/{id:[0-9]+}",
		controllers.GetProjectAnalytic).Methods("GET")

	router.HandleFunc("/api/v1/connector/projects",
		controllers.GetAllProjectsFromConnector).Methods("GET")

	router.HandleFunc("/api/v1/connector/updateProject",
		controllers.AddProjectToDB).Methods("POST")

	router.HandleFunc("/api/v1/graph/get/{group:[0-9]}",
		controllers.GetGraphByGroup).Methods("GET")

	router.HandleFunc("/api/v1/graph/make/{group:[0-9]}",
		controllers.MakeGraphByGroup).Methods("POST")

	router.HandleFunc("/api/v1/graph/delete",
		controllers.DeleteGraphByProject).Methods("DELETE")

	router.HandleFunc("/api/v1/isAnalyzed",
		controllers.IsAnalyzed).Methods("GET")

	router.HandleFunc("/api/v1/compare/{group:[0-9]}",
		controllers.GetCompareByGraphGroup).Methods("GET")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	fmt.Println(port)

	err := http.ListenAndServe(":"+port, router)

	if err != nil {
		fmt.Print(err)
	}
}
