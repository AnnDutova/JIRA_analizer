package repository

import (
	"database/sql"
)

type IssueRepository struct {
	db *sql.DB
}

func NewIssueRepository(db_ *sql.DB) *IssueRepository {
	return &IssueRepository{
		db: db_,
	}
}
