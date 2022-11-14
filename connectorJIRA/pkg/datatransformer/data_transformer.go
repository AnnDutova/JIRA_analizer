package datatransformer

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

type Project struct {
	Id   int
	Name string
}

type Issue struct {
	Id          int
	Key         string
	Project     Project
	CreatedTime time.Time
	ClosedTime  time.Time
	UpdatedTime time.Time
	Summary     string
	Description string
	Type        string
	Priority    string
	Status      string
	Creator     string
	Assignee    string
	TimeSpent   int
}

type Page struct {
	Issues []JiraIssue `json:"issues" structs:"issues"`
}

type IssueWithChangelog struct {
	Id      string    `json:"id" structs:"id"`
	Changes Changelog `json:"changelog" structs:"changelog"`
}

type IssueStatusChanges struct {
	Id        int
	Histories []StatusChange
}

type StatusChange struct {
	Author     string
	ChangeTime time.Time
	FromStatus string
	ToStatus   string
}

func (is *Issue) ToJSON() string {
	str, _ := json.Marshal(is)
	return string(str)
}

func (stCh *IssueStatusChanges) ToJSON() string {
	str, _ := json.Marshal(stCh)
	return string(str)
}

func FormatIssues(issues []byte) ([]Issue, error) {
	var body Page
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
		priority := issue.Fields.Priority.Name
		timeSpent := issue.Fields.TimeSpent
		description := issue.Fields.Description
		creator := issue.Fields.Creator.DisplayName
		assignee := ""
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

func FormatChangelog(changelog []byte) IssueStatusChanges {
	var body IssueWithChangelog
	err := json.Unmarshal(changelog, &body)
	if err != nil {
		fmt.Printf("Error when unmarshal JSON: %s\n", err)
		os.Exit(1)
	}
	issueStatusChanges := IssueStatusChanges{}
	issueStatusChanges.Id, err = strconv.Atoi(body.Id)
	if err != nil {
		panic(err)
	}
	for _, history := range body.Changes.Histories {
		author := history.Author.DisplayName
		change := history.Created
		change = change[:len(change)-5] + "Z" //replace timezone for parsing
		changeTime, err := time.Parse(time.RFC3339, change)
		if err != nil {
			fmt.Printf("Error when parsing time: %s\n", err)
			os.Exit(1)
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

	return issueStatusChanges
}
