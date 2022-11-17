package datatransformer

import "time"

type ProjectsRespond struct {
	Projects  []Project
	PageCount int
}

type Project struct {
	Id   int
	Key  string
	Name string
	Url  string
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
