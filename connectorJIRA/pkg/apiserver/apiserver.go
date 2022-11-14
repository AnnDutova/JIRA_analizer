package apiserver

import (
	"connectorJIRA/pkg/connector"
	"connectorJIRA/pkg/datapusher"
	"connectorJIRA/pkg/properties"
	"database/sql"
	"fmt"
	"net/http"
)

const (
	bindAddr = ":8050"
)

func Start(config *properties.Config) error {
	db, err := newDB(config)
	if err != nil {
		return err
	}
	defer db.Close()

	con, err := datapusher.New(db)
	if err != nil {
		return err
	}

	jcon, err := newJIRAConnection(config)
	if err != nil {
		return err
	}

	router := &Router{
		dbConnector:   con,
		JIRAConnector: jcon,
	}

	configureRouters(router)

	return http.ListenAndServe(bindAddr, nil)
}

func newDB(config *properties.Config) (*sql.DB, error) {
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

func newJIRAConnection(config *properties.Config) (*connector.Connection, error) {
	url := config.ProgramSettings.ApacheUrl
	return connector.GetConnection(url)
}
