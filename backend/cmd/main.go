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
