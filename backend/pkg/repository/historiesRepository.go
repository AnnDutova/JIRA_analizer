package repository

import (
	"database/sql"
)

type HistoriesRepository struct {
	db *sql.DB
}

func NewHistoriesRepository(db_ *sql.DB) *HistoriesRepository {
	return &HistoriesRepository{
		db: db_,
	}
}
