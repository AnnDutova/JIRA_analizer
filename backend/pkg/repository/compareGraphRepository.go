package repository

import (
	"Backend/pkg/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
)

type CompareGraphRepository struct {
	db *gorm.DB
}

func NewCompareGraphRepository(db_ *gorm.DB) *CompareGraphRepository {
	return &CompareGraphRepository{
		db: db_,
	}
}

func (r *CompareGraphRepository) CheckExistenceOnOpenTaskTimeTable(projectName string) ([]models.GraphOutput, error) {
	var data []models.GraphOutput
	var request []byte
	row := r.db.Raw("Select data from \"openTaskTime\" "+
		"left join project on projectId = project.id where project.title = ?", projectName).Row()
	if row.Err() != nil {
		if row.Err() == sql.ErrNoRows {
			return nil, nil
		}
		return nil, row.Err()
	}
	err := row.Scan(&request)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} else {
			return nil, err
		}
	}

	if err := json.Unmarshal(request, &data); err != nil {
		return nil, err
	}
	return data, nil
}

func (r *CompareGraphRepository) CheckExistenceOnTaskStateTimeTableOpen(projectName string) ([]models.GraphOutput, error) {
	var data []models.GraphOutput
	var request []byte
	row := r.db.Raw("Select data from \"taskStateTime\" "+
		"left join project on projectId = project.id where project.title = ? and state = 'Open'", projectName).Row()
	if row.Err() != nil {
		if row.Err() == sql.ErrNoRows {
			return nil, nil
		}
		return nil, row.Err()
	}
	err := row.Scan(&request)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} else {
			return nil, err
		}
	}

	if err := json.Unmarshal(request, &data); err != nil {
		return nil, err
	}
	return data, nil
}

func (r *CompareGraphRepository) CheckExistenceOnTaskStateTimeTableResolved(projectName string) ([]models.GraphOutput, error) {
	var data []models.GraphOutput
	var request []byte
	row := r.db.Raw("Select data from \"taskStateTime\" "+
		"left join project on projectId = project.id where project.title = ? and state = 'Resolved'", projectName).Row()
	if row.Err() != nil {
		if row.Err() == sql.ErrNoRows {
			return nil, nil
		}
		return nil, row.Err()
	}
	err := row.Scan(&request)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} else {
			return nil, err
		}
	}

	if err := json.Unmarshal(request, &data); err != nil {
		return nil, err
	}
	return data, nil
}

func (r *CompareGraphRepository) CheckExistenceOnTaskStateTimeTableReopened(projectName string) ([]models.GraphOutput, error) {
	var data []models.GraphOutput
	var request []byte
	row := r.db.Raw("Select data from \"taskStateTime\" "+
		"left join project on projectId = project.id where project.title = ? and state = 'Reopened'", projectName).Row()
	if row.Err() != nil {
		if row.Err() == sql.ErrNoRows {
			return nil, nil
		}
		return nil, row.Err()
	}
	err := row.Scan(&request)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} else {
			return nil, err
		}
	}

	if err := json.Unmarshal(request, &data); err != nil {
		return nil, err
	}
	return data, nil
}

func (r *CompareGraphRepository) CheckExistenceOnTaskStateTimeTableInProgress(projectName string) ([]models.GraphOutput, error) {
	var data []models.GraphOutput
	var request []byte
	row := r.db.Raw("Select data from \"taskStateTime\" "+
		"left join project on projectId = project.id where project.title = ? and state = 'In progress'", projectName).Row()
	if row.Err() != nil {
		if row.Err() == sql.ErrNoRows {
			return nil, nil
		}
		return nil, row.Err()
	}
	err := row.Scan(&request)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} else {
			return nil, err
		}
	}

	if err := json.Unmarshal(request, &data); err != nil {
		return nil, err
	}
	return data, nil
}

func (r *CompareGraphRepository) CheckExistenceOnActivityByTaskTableClose(projectName string) ([]models.GraphOutput, error) {
	var data []models.GraphOutput
	var request []byte
	row := r.db.Raw("Select data from \"activityByTask\" "+
		"left join project on projectId = project.id where project.title = ? and state = 'Closed'", projectName).Row()
	if row.Err() != nil {
		if row.Err() == sql.ErrNoRows {
			return nil, nil
		}
		return nil, row.Err()
	}
	err := row.Scan(&request)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} else {
			return nil, err
		}
	}

	if err := json.Unmarshal(request, &data); err != nil {
		return nil, err
	}
	return data, nil
}

func (r *CompareGraphRepository) CheckExistenceOnActivityByTaskTableOpen(projectName string) ([]models.GraphOutput, error) {
	var data []models.GraphOutput
	var request []byte
	row := r.db.Raw("Select data from \"activityByTask\" "+
		"left join project on projectId = project.id where project.title = ? and state = 'Open'", projectName).Row()
	if row.Err() != nil {
		if row.Err() == sql.ErrNoRows {
			return nil, nil
		}
		return nil, row.Err()
	}
	err := row.Scan(&request)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} else {
			return nil, err
		}
	}

	if err := json.Unmarshal(request, &data); err != nil {
		return nil, err
	}
	return data, nil
}

func (r *CompareGraphRepository) CheckExistenceOnComplexityTaskTimeTable(projectName string) ([]models.GraphOutput, error) {
	var data []models.GraphOutput
	var request []byte
	row := r.db.Raw("Select data from \"complexityTaskTime\" "+
		"left join project on projectId = project.id where project.title = ?", projectName).Row()
	if row.Err() != nil {
		if row.Err() == sql.ErrNoRows {
			return nil, nil
		}
		return nil, row.Err()
	}
	err := row.Scan(&request)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} else {
			return nil, err
		}
	}

	if err := json.Unmarshal(request, &data); err != nil {
		return nil, err
	}
	return data, nil
}

func (r *CompareGraphRepository) CheckExistenceOnTaskPriorityCountTableOpen(projectName string) ([]models.GraphOutput, error) {
	var data []models.GraphOutput
	var request []byte
	err := r.db.Raw("Select data from \"taskPriorityCount\" "+
		"left join project on projectId = project.id where project.title = ? and state = 'All'", projectName).Row().Scan(&request)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} else {
			return nil, err
		}
	}

	if err := json.Unmarshal(request, &data); err != nil {
		return nil, err
	}
	return data, nil
}

func (r *CompareGraphRepository) CheckExistenceOnTaskPriorityCountTableClose(projectName string) ([]models.GraphOutput, error) {
	var data []models.GraphOutput
	var request []byte
	query := fmt.Sprintf("Select data from \"taskPriorityCount\" left join project on projectId = project.id where project.title = '%s' and state = 'Closed'", projectName)
	err := r.db.Raw(query).Row().Scan(&request)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} else {
			return nil, err
		}
	}

	if err := json.Unmarshal(request, &data); err != nil {
		return nil, err
	}
	return data, nil
}
