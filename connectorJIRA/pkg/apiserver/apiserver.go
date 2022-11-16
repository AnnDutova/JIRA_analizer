package apiserver

import (
	"connectorJIRA/pkg/connector"
	"connectorJIRA/pkg/datapusher"
	"connectorJIRA/pkg/properties"
	"database/sql"
	"fmt"
	"net/http"
	"os"
)

const (
	bindAddr = ":8050"
)

func Start() error {
	config, err := properties.GetConfig(os.Args[1])
	if err != nil {
		return err
	}

	db, err := newDB()
	if err != nil {
		return err
	}
	defer db.Close()

	con, err := datapusher.New(db)
	if err != nil {
		return err
	}

	jcon, err := newJIRAConnection()
	if err != nil {
		return err
	}

	router := &Router{
		dbConnector:       con,
		JIRAConnector:     jcon,
		issueInOneRequest: config.ProgramSettings.IssueInOneRequest,
		threadCount:       config.ProgramSettings.ThreadCount,
	}

	configureRouters(router)

	return http.ListenAndServe(bindAddr, nil)
}

func newDB() (*sql.DB, error) {
	config, err := properties.GetConfig(os.Args[1])
	if err != nil {
		return nil, err
	}
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
	return connector.GetConnection()
}
