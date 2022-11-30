package repository

import (
	"gorm.io/gorm"
)

type HistoriesRepository struct {
	db *gorm.DB
}

func NewHistoriesRepository(db_ *gorm.DB) *HistoriesRepository {
	return &HistoriesRepository{
		db: db_,
	}
}
