package apiserver

import (
	"connectorJIRA/pkg/connector"
	"connectorJIRA/pkg/datapusher"
	"connectorJIRA/pkg/logging"
	"connectorJIRA/pkg/properties"
	"database/sql"
	"fmt"
	"net/http"
	"os"
)

func Start() error {

	logger := logging.GetLogger()
	logger.Info("Starting apiserver")
	config := properties.GetConfig(os.Args[1])

	logger.Info("Try connecting to DB")
	db, err := newDB()
	if err != nil {
		logger.Fatal(err)
	}
	logger.Info("Connection to DB was successful")
	defer db.Close()

	con := datapusher.New(db)

	logger.Info("Try connecting to JIRA")
	jcon, err := newJIRAConnection()
	if err != nil {
		logger.Fatal(err)
	}
	logger.Info("Connection to JIRA was successful")

	router := &Router{
		logger:            logger,
		dbConnector:       con,
		JIRAConnector:     jcon,
		issueInOneRequest: config.ProgramSettings.IssueInOneRequest,
		threadCount:       config.ProgramSettings.ThreadCount,
	}
	logger.Info("Configurate routers")
	configureRouters(router)
	logger.Info("Apiserver is started!")

	return http.ListenAndServe(config.ProgramSettings.BindAddress, nil)
}

func newDB() (*sql.DB, error) {
	config := properties.GetConfig(os.Args[1])
	dbName := config.DbSettings.DbName
	dbUsername := config.DbSettings.DbUsername
	dbPassword := config.DbSettings.DbPassword
	dbPort := config.DbSettings.DbPort

	connStr := fmt.Sprintf("dbname=%s user=%s password=%s port=%s sslmode=disable", dbName, dbUsername, dbPassword,
		dbPort)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func newJIRAConnection() (*connector.Connection, error) {
	return connector.GetConnection(os.Args[1])
}
