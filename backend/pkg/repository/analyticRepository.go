package repository

import (
	"Backend/pkg/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"sort"
	"time"
)

type AnalyticRepository struct {
	db *gorm.DB
}

func NewAnalyticRepository(db_ *gorm.DB) *AnalyticRepository {
	return &AnalyticRepository{
		db: db_,
	}
}

func (r *AnalyticRepository) ReturnTheMostActiveCreators(projectName string) ([]models.GraphOutput, error) {
	var graph []models.GraphOutput
	rows, err := r.db.Raw("Select author.name as creator, count(author.name) as count "+
		"from issues as i join project on project.id = projectId join author on author.id = authorId "+
		"where project.title = ? group by author.name order by count desc", projectName).Rows()

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		nameCount := models.GraphOutput{}
		if err = rows.Scan(&nameCount.Title, &nameCount.Count); err != nil {
			return nil, err
		}
		graph = append(graph, nameCount)
	}
	return graph, nil
}

func (r *AnalyticRepository) ReturnTimeCountOfIssuesInCloseState(projectName string) ([]models.GraphOutput, error) {
	var graph []models.GraphOutput
	var creationTime time.Time
	var request []byte

	if row := r.db.Raw("Select createdTime,"+" data from \"openTaskTime\" "+
		"left join project on projectId = project.id where project.title = ?", projectName).Row(); row.Err() != nil {
		return nil, row.Err()
	} else {
		err := row.Scan(&creationTime, &request)
		if err != nil {
			if err != sql.ErrNoRows {
				return nil, err
			}
		}
		if time.Now().Sub(creationTime).Hours() < 1.0 {
			if err := json.Unmarshal(request, &graph); err != nil {
				return nil, err
			}
			return graph, nil
		} else {
			graph, err := r.returnOpenTimeInClose(projectName)
			if err != nil {
				return nil, err
			}

			id, err := r.returnProjectId(projectName)
			if err != nil {
				return nil, err
			}

			if res, err := json.Marshal(graph); err != nil {
				return nil, err
			} else {
				if err = r.db.Exec("call addOpenTaskTime(?, ?, ?)", id, time.Now(), res).Error; err != nil {
					return nil, err
				}
			}
			return graph, nil
		}
	}
}

func (r *AnalyticRepository) returnProjectId(projectName string) (int, error) {
	var id int
	row := r.db.Raw("Select project.id from project where project.title = ?", projectName).Row()
	if row.Err() != nil {
		return 0, row.Err()
	}
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AnalyticRepository) returnOpenTimeInClose(projectName string) ([]models.GraphOutput, error) {
	var graph []models.TimeCount
	var resultSet []models.GraphOutput
	mapTimeCount := make(map[int]int)

	rows, err := r.db.Raw("Select i.id, i.createdtime from issues as i"+
		" left join project on i.projectId = project.id "+
		" where project.title = ? and i.status = 'Closed'", projectName).Rows()

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var issueId int
		var prevChangeTime time.Time
		if err = rows.Scan(&issueId, &prevChangeTime); err != nil {
			return nil, err
		}

		changes, err := r.db.Raw("Select sc.ChangeTime, sc.fromStatus, sc.toStatus from  \"statusChange\" as sc "+
			"where issueId = ?", issueId).Rows()

		if err != nil {
			return nil, err
		}

		var count time.Duration
		for changes.Next() {
			var change time.Time
			var fromStatus string
			var toStatus string
			if err = changes.Scan(&change, &fromStatus, &toStatus); err != nil {
				return nil, err
			}
			if fromStatus == "Reopen" {
				prevChangeTime = change
			} else if fromStatus == "In progress" {
				count += change.Sub(prevChangeTime)
				prevChangeTime = change
			}

			if toStatus == "Resolved" {
				count += change.Sub(prevChangeTime)
				//prevChangeTime = change
			} else if toStatus == "Closed" && fromStatus == "Open" {
				count += change.Sub(prevChangeTime)
				//prevChangeTime = change
			}
		}

		if val, exist := mapTimeCount[int(count.Hours())]; exist {
			mapTimeCount[int(count.Hours())] = val + 1
		} else {
			mapTimeCount[int(count.Hours())] = 1
		}
	}

	for key, val := range mapTimeCount {
		graph = append(graph, models.TimeCount{
			Time:  key,
			Count: val,
		})
	}

	sort.SliceStable(graph, func(i, j int) bool {
		return graph[i].Time < graph[j].Time
	})

	if len(graph) > 50 {
		resultSet = r.convertTimeWithYears(graph)
	} else {
		resultSet = r.convertTimeWithHours(graph)
	}

	return resultSet, nil
}

func (r *AnalyticRepository) convertTimeWithHours(graph []models.TimeCount) []models.GraphOutput {
	var resultSet []models.GraphOutput
	day := 1
	month := 1
	count := 0
	for i := 0; i < len(graph); i++ {
		if graph[i].Time/24 < 1 {
			resultSet = append(resultSet, models.GraphOutput{
				Title: fmt.Sprintf("%d hours", graph[i].Time),
				Count: graph[i].Count,
			})
		} else if graph[i].Time/24 <= day {
			count += graph[i].Count
		} else if graph[i].Time/24 > day && day <= 30 {
			count += graph[i].Count
			ans := models.GraphOutput{Title: fmt.Sprintf("%d day", day), Count: count}
			resultSet = append(resultSet, ans)
			day += 1
			count = 0
		} else if graph[i].Time/(24*30) <= month {
			count += graph[i].Count
		} else if graph[i].Time/(24*30) > month && month < 12 {
			count += graph[i].Count
			ans := models.GraphOutput{Title: fmt.Sprintf("%d month", month), Count: count}
			resultSet = append(resultSet, ans)
			month += 1
			count = 0
		} else {
			month += 1
			count += 1
		}
	}
	if month > 13 {
		ans := models.GraphOutput{Title: "1+year", Count: count}
		resultSet = append(resultSet, ans)
	}
	return resultSet
}

func (r *AnalyticRepository) convertTimeWithYears(graph []models.TimeCount) []models.GraphOutput {
	var resultSet []models.GraphOutput
	var count int
	day := 1
	month := 1
	year := 1
	for i := 0; i < len(graph); i++ {
		if graph[i].Time/24 < 1 {
			resultSet = append(resultSet, models.GraphOutput{
				Title: fmt.Sprintf("%d hours", graph[i].Time),
				Count: graph[i].Count,
			})
		} else if graph[i].Time/24 <= day {
			count += graph[i].Count
		} else if graph[i].Time/24 > day && day <= 30 {
			count += graph[i].Count
			ans := models.GraphOutput{Title: fmt.Sprintf("%d day", day), Count: count}
			resultSet = append(resultSet, ans)
			day += 1
			count = 0
		} else if graph[i].Time/(24*30) <= month {
			count += graph[i].Count
		} else if graph[i].Time/(24*30) > month && month < 12 {
			count += graph[i].Count
			ans := models.GraphOutput{Title: fmt.Sprintf("%d month", month), Count: count}
			resultSet = append(resultSet, ans)
			month += 1
			count = 0
		} else if graph[i].Time/(24*30*12) <= year {
			count += graph[i].Count
		} else if graph[i].Time/(24*30*12) > year && year <= 7 {
			count += graph[i].Count
			ans := models.GraphOutput{Title: fmt.Sprintf("%d year", year), Count: count}
			resultSet = append(resultSet, ans)
			year += 1
			count = 0
		} else {
			year += 1
			count += graph[i].Count
		}
	}
	if year > 7 {
		ans := models.GraphOutput{Title: "8+years", Count: count}
		resultSet = append(resultSet, ans)
	}
	return resultSet
}

func (r *AnalyticRepository) ReturnPriorityCountOfProjectOpen(projectName string) ([]models.GraphOutput, error) {
	return nil, nil
}

func (r *AnalyticRepository) ReturnPriorityCountOfProjectClose(projectName string) ([]models.GraphOutput, error) {
	return nil, nil
}

func (r *AnalyticRepository) ReturnTimeSpentOnAllTasks(projectName string) ([]models.GraphOutput, error) {
	var graph []models.GraphOutput
	var creationTime time.Time
	var request []byte

	if row := r.db.Raw("Select createdTime,"+" data from \"complexityTaskTime\" "+
		"left join project on projectId = project.id where project.title = ?", projectName).Row(); row.Err() != nil {
		return nil, row.Err()
	} else {
		err := row.Scan(&creationTime, &request)
		if err != nil {
			if err != sql.ErrNoRows {
				return nil, err
			}
		}
		if time.Now().Sub(creationTime).Hours() < 1.0 {
			if err := json.Unmarshal(request, &graph); err != nil {
				return nil, err
			}
			return graph, nil
		} else {
			graph, err = r.returnTimeSpentOnAllTasks(projectName)
			if err != nil {
				return nil, err
			}

			id, err := r.returnProjectId(projectName)
			if err != nil {
				return nil, err
			}

			if res, err := json.Marshal(graph); err != nil {
				return nil, err
			} else {
				if err = r.db.Exec("call addComplexityTaskTime(?, ?, ?)", id, time.Now(), res).Error; err != nil {
					return nil, err
				}
			}
			return graph, nil
		}
	}
}

func (r *AnalyticRepository) returnTimeSpentOnAllTasks(projectName string) ([]models.GraphOutput, error) {
	var graph []models.GraphOutput
	rows, err := r.db.Raw("Select i.timespent,"+" count(i.timespent) as count from issues as i"+
		" left join project on i.projectId = project.id "+
		" where i.timespent > 0 and project.title = ? "+
		" group by i.timespent order by i.timespent ", projectName).Rows()

	/*
		Select i.status,  i.timespent, count(i.timespent) as count from issues as i
		left join project on i.projectId = project.id
		where i.timespent > 0 and project.title = 'Airavata'
		group by i.timespent, i.status order by i.timespent ;
	*/

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		el := models.GraphOutput{}
		var seconds int
		if err = rows.Scan(&seconds, &el.Count); err != nil {
			return nil, err
		}
		el.Title = fmt.Sprintf("%dm", seconds/60)
		graph = append(graph, el)
	}
	return graph, nil
}

func (r *AnalyticRepository) ReturnCountOpenTaskInDay(projectName string) ([]models.GraphOutput, error) {
	return nil, nil
}

func (r *AnalyticRepository) ReturnCountCloseTaskInDay(projectName string) ([]models.GraphOutput, error) {
	return nil, nil
}

func (r *AnalyticRepository) ReturnCountTimeOfOpenStateInCloseTask(projectName string) ([]models.GraphOutput, error) {
	return nil, nil
}

func (r *AnalyticRepository) ReturnCountTimeOfResolvedStateInCloseTask(projectName string) ([]models.GraphOutput, error) {
	return nil, nil
}

func (r *AnalyticRepository) ReturnCountTimeOfReopenedStateInCloseTask(projectName string) ([]models.GraphOutput, error) {
	return nil, nil
}

func (r *AnalyticRepository) ReturnCountTimeOfInProgressStateInCloseTask(projectName string) ([]models.GraphOutput, error) {
	return nil, nil
}
