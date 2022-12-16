package repository

import (
	"Backend/pkg/app"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type DBController struct {
	db   *gorm.DB
	repo *Repository
}

func NewDBController(path string) *DBController {
	config := app.NewConfig(path)
	db, err := setDBConnection(config)
	if err != nil {
		log.Fatal(err)
	}

	return &DBController{
		db:   db,
		repo: NewRepository(db),
	}
}

var DbCon *DBController

func setDBConnection(config *app.Config) (*gorm.DB, error) {
	username := config.DbSettings.DbUsername
	password := config.DbSettings.DbPassword
	dbName := config.DbSettings.DbName
	dbHost := config.DbSettings.DbHost
	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)

	db, err := gorm.Open(postgres.Open(dbUri), &gorm.Config{})
	if err != nil {
		fmt.Print(err)
	}
	return db, nil
}

func (db_ *DBController) GetDB() *gorm.DB {
	return db_.db
}

func (db_ *DBController) GetRepository() *Repository {
	return db_.repo
}
