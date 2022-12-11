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

type IAnalytic interface {
	ReturnTheMostActiveCreators(projectName string) ([]models.GraphOutput, error)
	ReturnTimeCountOfIssuesInCloseState(projectName string) ([]models.GraphOutput, error)
	ReturnTimeSpentOnAllTasks(projectName string) ([]models.GraphOutput, error)

	ReturnPriorityCountOfProjectOpen(projectName string) ([]models.GraphOutput, error)
	ReturnPriorityCountOfProjectClose(projectName string) ([]models.GraphOutput, error)

	ReturnCountOpenTaskInDay(projectName string) ([]models.GraphOutput, error)
	ReturnCountCloseTaskInDay(projectName string) ([]models.GraphOutput, error)

	ReturnCountTimeOfOpenStateInCloseTask(projectName string) ([]models.GraphOutput, error)
	ReturnCountTimeOfResolvedStateInCloseTask(projectName string) ([]models.GraphOutput, error)
	ReturnCountTimeOfReopenedStateInCloseTask(projectName string) ([]models.GraphOutput, error)
	ReturnCountTimeOfInProgressStateInCloseTask(projectName string) ([]models.GraphOutput, error)
}

type ICompareGraphs interface {
	CheckExistenceOnOpenTaskTimeTable(projectName string) ([]models.GraphOutput, error)

	CheckExistenceOnTaskStateTimeTableOpen(projectName string) ([]models.GraphOutput, error)
	CheckExistenceOnTaskStateTimeTableResolved(projectName string) ([]models.GraphOutput, error)
	CheckExistenceOnTaskStateTimeTableReopened(projectName string) ([]models.GraphOutput, error)
	CheckExistenceOnTaskStateTimeTableInProgress(projectName string) ([]models.GraphOutput, error)

	CheckExistenceOnActivityByTaskTableClose(projectName string) ([]models.GraphOutput, error)
	CheckExistenceOnActivityByTaskTableOpen(projectName string) ([]models.GraphOutput, error)

	CheckExistenceOnComplexityTaskTimeTable(projectName string) ([]models.GraphOutput, error)

	CheckExistenceOnTaskPriorityCountTableOpen(projectName string) ([]models.GraphOutput, error)
	CheckExistenceOnTaskPriorityCountTableClose(projectName string) ([]models.GraphOutput, error)
}

type Repository struct {
	IIssue
	IProject
	IHistories
	IConnector
	IAnalytic
	ICompareGraphs
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		IIssue:         NewIssueRepository(db),
		IProject:       NewProjectRepository(db),
		IHistories:     NewHistoriesRepository(db),
		IConnector:     NewConnectorRepository(db),
		IAnalytic:      NewAnalyticRepository(db),
		ICompareGraphs: NewCompareGraphRepository(db),
	}
}
