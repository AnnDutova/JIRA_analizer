package datapusher

import (
	"connectorJIRA/pkg/datatransformer"
	"connectorJIRA/pkg/properties"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
)

const (
	YYYYMMDD = "2006-01-02"
)

func PushIssues(issues []datatransformer.Issue) {
	config := properties.GetConfig(os.Args[1])
	dbName := config.DbSettings.DbName
	dbUsername := config.DbSettings.DbUsername
	dbPassword := config.DbSettings.DbPassword
	dbPort := config.DbSettings.DbPort
	connStr := fmt.Sprintf("dbname=%s user=%s password=%s port=%s sslmode=disable", dbName, dbUsername, dbPassword,
		dbPort)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	for _, val := range issues {
		command := fmt.Sprintf("select insertIssue('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s' )", val.Key,
			val.CreatedTime.Format(YYYYMMDD), val.ClosedTime.Format(YYYYMMDD), val.Summary, val.Type, val.Priority, val.Status, val.ToJSON())
		fmt.Println(command)
		_, err = db.Exec(command)
		if err != nil {
			log.Fatal(err)
		}
	}
}
