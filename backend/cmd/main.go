package main

import (
	"Backend/pkg/app"
	"Backend/pkg/controllers"
	"Backend/pkg/repository"
	"Backend/pkg/utils"
	"flag"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "../config/config.yaml", "path to config file")
}

func main() {
	config := app.NewConfig(configPath)
	repository.DbCon = repository.NewDBController(config)
	utils.Init()
	logger := utils.GetLogger()
	logger.Info("Start backend work")
	router := mux.NewRouter()

	router.HandleFunc("/api/v1/projects",
		controllers.GetProjectsFor).Methods("GET")

	router.HandleFunc("/api/v1/projects/{id:[0-9]+}",
		controllers.GetProjectAnalytic).Methods("GET")

	router.HandleFunc("/api/v1/projects/{id:[0-9]+}",
		controllers.DeleteProjectById).Methods("DELETE")

	router.HandleFunc("/api/v1/projects/{id:[0-9]+}",
		controllers.OptionsReq).Methods("OPTIONS")

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

	router.HandleFunc("/api/v1/graph/delete",
		controllers.OptionsReq).Methods("OPTIONS")

	router.HandleFunc("/api/v1/isAnalyzed",
		controllers.IsAnalyzed).Methods("GET")

	router.HandleFunc("/api/v1/compare/{group:[0-9]}",
		controllers.GetCompareByGraphGroup).Methods("GET")

	port := fmt.Sprintf("%d", config.Backend.Port)
	fmt.Println(port)
	logger.Info("Backend work on port ", port)
	err := http.ListenAndServe(":"+port, router)

	if err != nil {
		fmt.Print(err)
	}
}
