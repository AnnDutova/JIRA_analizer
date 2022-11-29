package repository

import (
	"gorm.io/gorm"
)

type IssueRepository struct {
	db *gorm.DB
}

func NewIssueRepository(db_ *gorm.DB) *IssueRepository {
	return &IssueRepository{
		db: db_,
	}
}
