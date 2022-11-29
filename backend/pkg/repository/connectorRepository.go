package repository

import (
	"Backend/pkg/models"
	"encoding/json"
	"gorm.io/gorm"
	"io"
	"net/http"
)

type ConnectorRepository struct {
	db *gorm.DB
}

func NewConnectorRepository(db_ *gorm.DB) *ConnectorRepository {
	return &ConnectorRepository{
		db: db_,
	}
}

func (r *ConnectorRepository) AddProjectToDB(project string) (*models.Project, error) {
	httpResp, err := http.Get("http://localhost:8050/updateProject?project=" + project)

	if err != nil {
		return nil, err
	}

	if httpResp.StatusCode != http.StatusOK {
		var connectorError models.ConnectorError
		json.NewDecoder(httpResp.Body).Decode(&connectorError)
		return nil, connectorError
	}

	defer httpResp.Body.Close()

	return nil, nil
}

func (r *ConnectorRepository) ReturnAllProjectsFromConnector(limit, page, search string) (*models.Projects, error) {
	httpResp, err := http.Get("http://localhost:8050/projects?limit=" + limit + "&page=" + page + "&search=" + search)
	projects := &models.Projects{}

	if err != nil {
		return nil, err
	}

	if httpResp.StatusCode != http.StatusOK {
		var connectorError models.ConnectorError
		json.NewDecoder(httpResp.Body).Decode(&connectorError)
		return nil, connectorError
	}

	body, _ := io.ReadAll(httpResp.Body)
	defer httpResp.Body.Close()

	err = json.Unmarshal(body, &projects)

	for _, project := range projects.Projects {
		err = r.db.Raw("select exists(select "+
			"from project where id =?)", project.Id).Scan(&project.Existence).Error
	}

	return projects, nil
}
