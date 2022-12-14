package repository

import (
	"Backend/pkg/models"
	"database/sql"
	"gorm.io/gorm"
	"math"
	"strconv"
	"strings"
)

type ProjectRepository struct {
	db *gorm.DB
}

func NewProjectRepository(db_ *gorm.DB) *ProjectRepository {
	return &ProjectRepository{
		db: db_,
	}
}

func (r *ProjectRepository) ReturnAllProjects(limit int64, page uint64, search string) ([]*models.Project, uint64, error) {
	projects := make([]*models.Project, 0)
	var err error
	var pageCount uint64

	absLimit := uint64(math.Abs(float64(limit)))
	pageCount, err = r.getPageCount(absLimit, search)

	if page != 0 && page <= pageCount {
		if limit != -1 {
			searchName := "%" + strings.ToLower(search) + "%"
			err = r.db.Raw("select * "+
				"from project where lower(title) like ? LIMIT ? OFFSET ?",
				searchName, limit, uint64(limit)*(page-1)).
				Scan(&projects).Error

			if err != nil {
				return nil, 0, err
			}
		} else {
			searchName := "%" + strings.ToLower(search) + "%"
			err = r.db.Raw("select * "+
				"from project where lower(title) like ?",
				searchName).
				Scan(&projects).Error

			pageCount = 1
		}
	}

	if err != nil {
		return nil, 0, err
	}

	return projects, pageCount, err
}

func (r *ProjectRepository) getPageCount(limit uint64, search string) (uint64, error) {
	var pageCount int64
	searchName := "%" + strings.ToLower(search) + "%"
	err := r.db.Raw("select count(*) from project where lower(title) like ?",
		searchName).Scan(&pageCount).Error

	if err != nil {
		return 0, err
	}

	if limit == 0 {
		pageCount = int64(math.Ceil(float64(pageCount) / float64(1)))
	} else {
		pageCount = int64(math.Ceil(float64(pageCount) / float64(limit)))
	}

	return uint64(pageCount), nil
}

func (r *ProjectRepository) ReturnProjectAnalytic(id string) (*models.ProjectAnalytic, error) {
	ProjectAnalytic := &models.ProjectAnalytic{}
	ProjectAnalytic.Id, _ = strconv.Atoi(id)

	err := r.db.Raw("select title "+
		"from project where Id = ?", id).Scan(&ProjectAnalytic.Title).Error

	if err != nil {
		return nil, err
	}

	err = r.db.Raw("select count(*) "+
		"from issues where projectID = ?", id).Scan(&ProjectAnalytic.AllIssuesCount).Error

	if err != nil {
		return nil, err
	}

	err = r.db.Raw("select count(*) "+
		"from issues where projectID = ? AND status = 'Open'", id).Scan(&ProjectAnalytic.OpenIssuesCount).Error

	if err != nil {
		return nil, err
	}

	err = r.db.Raw("select count(*) "+
		"from issues where projectID = ? AND status = 'Closed'", id).Scan(&ProjectAnalytic.CloseIssuesCount).Error

	if err != nil {
		return nil, err
	}

	var averageTime sql.NullFloat64
	err = r.db.Raw("select AVG(timeSpent) "+
		"from issues where projectID = ?", id).Scan(&averageTime).Error

	if !averageTime.Valid {
		ProjectAnalytic.AverageTime = 0
	} else {
		ProjectAnalytic.AverageTime = averageTime.Float64
	}

	if err != nil {
		return nil, err
	}

	var averageIssuesCount sql.NullFloat64
	err = r.db.Raw("select AVG(a.count) from (select createdTime, count(*) as count from issues "+
		"where projectID = ? and EXTRACT(DAY FROM CURRENT_DATE - createdTime) <= 7 GROUP BY createdTime)"+
		" as a", id).Scan(&averageIssuesCount).Error

	if !averageIssuesCount.Valid {
		ProjectAnalytic.AverageIssuesCount = 0
	} else {
		ProjectAnalytic.AverageIssuesCount = averageIssuesCount.Float64
	}

	if err != nil {
		return nil, err
	}

	return ProjectAnalytic, err
}

func (r *ProjectRepository) DeleteProjectById(id string) (*models.Project, error) {
	project := &models.Project{}

	err := r.db.Raw("DELETE FROM issues WHERE projectId = ?", id).Scan(project).Error
	if err != nil {
		return nil, err
	}

	err = r.db.Raw("DELETE FROM project WHERE id = ? RETURNING *", id).Scan(project).Error
	if err != nil {
		return nil, err
	}

	return project, nil
}
