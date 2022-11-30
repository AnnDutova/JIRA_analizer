package repository

import "gorm.io/gorm"

type ProjectRepository struct {
	db *gorm.DB
}

func NewProjectRepository(db_ *gorm.DB) *ProjectRepository {
	return &ProjectRepository{
		db: db_,
	}
}
