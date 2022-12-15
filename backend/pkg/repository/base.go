package repository

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
)

type DBController struct {
	db   *sql.DB
	repo *Repository
}

func NewDBController() *DBController {
	db, err := setDBConnection()
	if err != nil {
		log.Fatal(err)
	}

	return &DBController{
		db:   db,
		repo: NewRepository(db),
	}
}

var DbCon *DBController

func init() {
	DbCon = NewDBController()
	fmt.Println("Говно какое-то")
}

func setDBConnection() (*sql.DB, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")

	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password) //Создать строку подключения

	db, err := sql.Open("postgres", dbUri)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	/*db, err := gorm.Open(postgres.Open(dbUri), &gorm.Config{})
	if err != nil {
		fmt.Print(err)
	}*/
	return db, nil
}

func (db_ *DBController) GetDB() *sql.DB {
	return db_.db
}

func (db_ *DBController) GetRepository() *Repository {
	return db_.repo
}
