package repository

import (
	"Backend/pkg/models"
	"gorm.io/gorm"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type IIssue interface {
}

type IConnector interface {
	AddProjectToDB(project string) (*models.Project, error)
	ReturnAllProjectsFromConnector(limit, page, search string) (*models.Projects, error)
}

type IProject interface {
	ReturnAllProjects(limit int64, page uint64, search string) ([]*models.Project, uint64, error)
	ReturnProjectAnalytic(string) (*models.ProjectAnalytic, error)
}

type IHistories interface {
}

type Repository struct {
	IIssue
	IProject
	IHistories
	IConnector
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		IIssue:     NewIssueRepository(db),
		IProject:   NewProjectRepository(db),
		IHistories: NewHistoriesRepository(db),
		IConnector: NewConnectorRepository(db),
	}
}
