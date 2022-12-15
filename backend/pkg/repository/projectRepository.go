package repository

import (
	"Backend/pkg/models"
	"database/sql"
	"fmt"
	"math"
	"strconv"
)

type ProjectRepository struct {
	db *sql.DB
}

func NewProjectRepository(db_ *sql.DB) *ProjectRepository {
	return &ProjectRepository{
		db: db_,
	}
}

func (r *ProjectRepository) ReturnAllProjects(limit int64, page uint64, search string) ([]models.Project, uint64, error) {
	projects := make([]models.Project, 0)
	var err error
	var pageCount uint64

	absLimit := uint64(math.Abs(float64(limit)))
	pageCount, err = r.getPageCount(absLimit, search)

	if page != 0 && page <= pageCount {
		var project models.Project
		if limit != -1 {
			searchName := "%" + search + "%"
			query := fmt.Sprintf("select * "+
				"from project where title like '%s' LIMIT %d OFFSET %d",
				searchName, limit, uint64(limit)*(page-1))
			rows, err := r.db.Query(query)
			if err != nil {
				return nil, 0, err
			}
			for rows.Next() {
				err := rows.Scan(&project.Id, &project.Title)
				if err != nil {
					return nil, 0, err
				}
				projects = append(projects, project)
			}

		} else {
			searchName := "%" + search + "%"
			query := fmt.Sprintf("select * from project where title like '%s'", searchName)
			rows, err := r.db.Query(query)
			if err != nil {
				return nil, 0, err
			}
			for rows.Next() {
				err := rows.Scan(&project.Id, &project.Title)
				if err != nil {
					return nil, 0, err
				}
				projects = append(projects, project)
			}
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
	err := r.db.QueryRow("select count(*) from project").Scan(&pageCount)

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
	ProjectAnalytic := models.ProjectAnalytic{}
	ProjectAnalytic.Id, _ = strconv.Atoi(id)

	query := fmt.Sprintf("select title from project where id = %s", id)
	err := r.db.QueryRow(query).Scan(&ProjectAnalytic.Title)

	if err != nil {
		return nil, err
	}

	query = fmt.Sprintf("select count(*) from issues where projectID = %s", id)
	err = r.db.QueryRow(query).Scan(&ProjectAnalytic.AllIssuesCount)

	if err != nil {
		return nil, err
	}

	query = fmt.Sprintf("select count(*) from issues where projectID = %s AND status = 'Open'", id)
	err = r.db.QueryRow(query).Scan(&ProjectAnalytic.OpenIssuesCount)

	if err != nil {
		return nil, err
	}

	query = fmt.Sprintf("select count(*) from issues where projectID = %s AND status = 'Closed'", id)
	err = r.db.QueryRow(query).Scan(&ProjectAnalytic.CloseIssuesCount)

	if err != nil {
		return nil, err
	}

	var averageTime sql.NullFloat64
	query = fmt.Sprintf("select AVG(timeSpent) from issues where projectID = %s", id)
	err = r.db.QueryRow(query).Scan(&averageTime)

	if !averageTime.Valid {
		ProjectAnalytic.AverageTime = 0
	} else {
		ProjectAnalytic.AverageTime = int(averageTime.Float64)
	}

	if err != nil {
		return nil, err
	}

	var averageIssuesCount sql.NullFloat64
	query = fmt.Sprintf("select AVG(a.count) from (select createdTime, count(*) as count from issues "+
		"where projectID = %s and EXTRACT(DAY FROM CURRENT_DATE - createdTime) <= 7 GROUP BY createdTime)"+
		" as a", id)
	err = r.db.QueryRow(query).Scan(&averageIssuesCount)

	if !averageIssuesCount.Valid {
		ProjectAnalytic.AverageIssuesCount = 0
	} else {
		ProjectAnalytic.AverageIssuesCount = int(averageIssuesCount.Float64)
	}

	if err != nil {
		return nil, err
	}

	return &ProjectAnalytic, err
}

func (r *ProjectRepository) DeleteProjectById(id string) (*models.Project, error) {
	project := models.Project{}

	_, err := r.db.Exec("DELETE FROM \"taskPriorityCount\" WHERE projectId = ?", id)
	if err != nil {
		return nil, err
	}

	_, err = r.db.Exec("DELETE FROM \"activityByTask\" WHERE projectId = ?", id)
	if err != nil {
		return nil, err
	}

	_, err = r.db.Exec("DELETE FROM \"taskStateTime\" WHERE projectId = ?", id)
	if err != nil {
		return nil, err
	}

	_, err = r.db.Exec("DELETE FROM \"complexityTaskTime\" WHERE projectId = ?", id)
	if err != nil {
		return nil, err
	}

	_, err = r.db.Exec("DELETE FROM \"openTaskTime\" WHERE projectId = ?", id)
	if err != nil {
		return nil, err
	}

	err = r.db.QueryRow("DELETE FROM issues WHERE projectId = ?", id).Scan(project)
	if err != nil {
		return nil, err
	}

	err = r.db.QueryRow("DELETE FROM project WHERE id = ? RETURNING *", id).Scan(project)
	if err != nil {
		return nil, err
	}

	return &project, nil
}
