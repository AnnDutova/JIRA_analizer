package datatransformer

import (
	"connectorJIRA/pkg/properties"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func (is *Issue) ToJSON() string {
	str, _ := json.Marshal(is)
	return string(str)
}

func (stCh *IssueStatusChanges) ToJSON() string {
	str, _ := json.Marshal(stCh)
	return string(str)
}

func FormatIssues(issues []byte) ([]Issue, error) {
	var body struct {
		Issues []JiraIssue `json:"issues" structs:"issues"`
	}
	err := json.Unmarshal(issues, &body)
	if err != nil {
		log.Printf("Error when unmarshal JSON: %s\n", err)
		return nil, errors.New("cannot unmarshal JSON with issues")
	}
	var arrNewIssues []Issue
	for _, issue := range body.Issues {
		key := issue.Key
		id, err := strconv.Atoi(issue.Id)
		if err != nil {
			log.Printf("Error when parsing issue id: %s\n", err)
			return nil, errors.New("cannot unmarshal issueID of issue: " + key)
		}
		issueType := issue.Fields.Type.Name
		status := issue.Fields.Status.Name
		summary := issue.Fields.Summary
		priority := "No"
		if issue.Fields.Priority != nil {
			priority = issue.Fields.Priority.Name
		}
		timeSpent := issue.Fields.TimeSpent
		description := issue.Fields.Description
		creator := "Unknown"
		if issue.Fields.Creator != nil {
			creator = issue.Fields.Creator.DisplayName
		}
		assignee := "Unknown"
		if issue.Fields.Assignee != nil {
			assignee = issue.Fields.Assignee.DisplayName
		}
		created := issue.Fields.Created
		created = created[:len(created)-5] + "Z" //replace timezone for parsing
		createdTime, err := time.Parse(time.RFC3339, created)
		if err != nil {
			log.Printf("Error when parsing time: %s\n", err)
			return nil, errors.New("cannot unmarshal createdTime of issue: " + key)
		}
		updated := issue.Fields.Updated
		updated = updated[:len(created)-5] + "Z" //replace timezone for parsing
		updatedTime, err := time.Parse(time.RFC3339, updated)
		if err != nil {
			log.Printf("Error when parsing time: %s\n", err)
			return nil, errors.New("cannot unmarshal updatedTime of issue: " + key)
		}
		histories := issue.Changelog.Histories
		var closedTime time.Time
		if status == "Closed" {
			length := len(histories)
			closed := histories[length-1].Created
			closed = closed[:len(closed)-5] + "Z"
			closedTime, err = time.Parse(time.RFC3339, closed)
			if err != nil {
				log.Printf("Error when parsing time: %s\n", err)
				return nil, errors.New("cannot unmarshal closedTime of issue: " + key)
			}
		}

		projectId, err := strconv.Atoi(issue.Fields.Project.ID)
		if err != nil {
			log.Printf("Error when parsing project id: %s\n", err)
			return nil, errors.New("cannot unmarshal projectID of issue: " + key)
		}
		project := Project{
			Id:   projectId,
			Key:  issue.Fields.Project.Key,
			Name: issue.Fields.Project.Name,
		}

		myIssue := Issue{
			Id:          id,
			Project:     project,
			Key:         key,
			CreatedTime: createdTime,
			ClosedTime:  closedTime,
			UpdatedTime: updatedTime,
			Summary:     summary,
			Description: description,
			Type:        issueType,
			Priority:    priority,
			Status:      status,
			Creator:     creator,
			Assignee:    assignee,
			TimeSpent:   timeSpent,
		}
		arrNewIssues = append(arrNewIssues, myIssue)
	}
	return arrNewIssues, nil
}

func ToFile(str string, name string) {
	file, err := os.Create(name + ".json")

	if err != nil {
		fmt.Println("Unable to create file:", err.Error())
	}
	defer file.Close()
	_, err = file.WriteString(str)
	if err != nil {
		fmt.Println("Unable to create file:", err.Error())
	}
}

func FormatChangelog(changelog []byte) (IssueStatusChanges, error) {
	var body struct {
		Id      string    `json:"id" structs:"id"`
		Changes Changelog `json:"changelog" structs:"changelog"`
	}
	err := json.Unmarshal(changelog, &body)
	if err != nil {
		return IssueStatusChanges{}, errors.New("Error when unmarshal JSON: " + err.Error())
	}
	issueStatusChanges := IssueStatusChanges{}
	issueStatusChanges.Id, err = strconv.Atoi(body.Id)
	if err != nil {
		return IssueStatusChanges{}, errors.New("Error when parse id: " + err.Error())
	}
	for _, history := range body.Changes.Histories {
		author := "Unknown"
		if history.Author != nil {
			author = history.Author.DisplayName
		}
		change := history.Created
		change = change[:len(change)-5] + "Z" //replace timezone for parsing
		changeTime, err := time.Parse(time.RFC3339, change)
		if err != nil {
			return IssueStatusChanges{}, errors.New("Error when parse time: " + err.Error())
		}
		for _, item := range history.Items {
			if item.Field == "status" {
				fromStatus := item.FromString
				toStatus := item.ToString

				statusChange := StatusChange{
					Author:     author,
					ChangeTime: changeTime,
					FromStatus: fromStatus,
					ToStatus:   toStatus,
				}
				issueStatusChanges.Histories = append(issueStatusChanges.Histories, statusChange)
			}
		}
	}

	return issueStatusChanges, nil
}

func FormatProjectsRespond(projects []byte, limit int, page int, search string) (*ProjectsRespond, error) {
	var body []JiraProject
	err := json.Unmarshal(projects, &body)
	if err != nil {
		return nil, errors.New("Error when unmarshaling json in FormatProjectsRespond: " + err.Error())
	}
	search = strings.ToLower(search)
	var searchProjects []JiraProject
	for _, project := range body {
		name := strings.ToLower(project.Name)
		key := strings.ToLower(project.Key)
		if strings.Contains(name, search) || strings.Contains(key, search) {
			searchProjects = append(searchProjects, project)
		}
	}
	projectCount := len(searchProjects)
	var pageCount int
	if projectCount%limit == 0 {
		pageCount = projectCount / limit
	} else {
		pageCount = projectCount/limit + 1
	}
	var projectsArr []Project
	for i := limit * (page - 1); i < limit*page && i < projectCount; i++ {
		project := searchProjects[i]
		id, err := strconv.Atoi(project.ID)
		if err != nil {
			return nil, errors.New("Cannot convert id to string in FormatProjectsRespond: " + err.Error())
		}
		config := properties.GetConfig(os.Args[1])
		url := config.ProgramSettings.JiraUrl
		projectUrl := url + "/projects/" + project.Key
		myProject := Project{
			Id:   id,
			Key:  project.Key,
			Name: project.Name,
			Url:  projectUrl,
		}
		projectsArr = append(projectsArr, myProject)
	}
	pageInfo := PageInfo{
		PageCount:     pageCount,
		ProjectsCount: projectCount,
		CurrentPage:   page,
	}
	projectsRespond := ProjectsRespond{
		Projects: projectsArr,
		PageInfo: &pageInfo,
	}

	return &projectsRespond, nil
}
