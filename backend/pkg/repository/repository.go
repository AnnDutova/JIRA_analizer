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
}

type IHistories interface {
}

type IAnalytic interface {
	ReturnTheMostActiveCreators(projectName string) ([]models.GraphOutput, error)         //done
	ReturnTimeCountOfIssuesInCloseState(projectName string) ([]models.GraphOutput, error) //done
	ReturnTimeSpentOnAllTasks(projectName string) ([]models.GraphOutput, error)           //done

	ReturnPriorityCountOfProjectOpen(projectName string) ([]models.GraphOutput, error)
	ReturnPriorityCountOfProjectClose(projectName string) ([]models.GraphOutput, error)

	ReturnCountOpenTaskInDay(projectName string) ([]models.GraphOutput, error)
	ReturnCountCloseTaskInDay(projectName string) ([]models.GraphOutput, error)

	ReturnCountTimeOfOpenStateInCloseTask(projectName string) ([]models.GraphOutput, error)
	ReturnCountTimeOfResolvedStateInCloseTask(projectName string) ([]models.GraphOutput, error)
	ReturnCountTimeOfReopenedStateInCloseTask(projectName string) ([]models.GraphOutput, error)
	ReturnCountTimeOfInProgressStateInCloseTask(projectName string) ([]models.GraphOutput, error)
}

type Repository struct {
	IIssue
	IProject
	IHistories
	IConnector
	IAnalytic
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		IIssue:     NewIssueRepository(db),
		IProject:   NewProjectRepository(db),
		IHistories: NewHistoriesRepository(db),
		IConnector: NewConnectorRepository(db),
		IAnalytic:  NewAnalyticRepository(db),
	}
}
