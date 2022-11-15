package connector

import (
	"connectorJIRA/pkg/datatransformer"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/andygrunwald/go-jira"
	"io"
	"net/http"
	"os"
	"strconv"
)

type Connection struct {
	client jira.Client
	url    string
}

func (con *Connection) getTotalIssues(projectName string) (int, error) {
	res, err := http.Get(con.url + "/rest/api/2/search?jql=project=" + projectName + "&expand=changelog")
	if err != nil {
		return 0, errors.New("Error when try to get request")
	}
	totalByte, err := io.ReadAll(res.Body)
	if err != nil {
		return 0, errors.New("Error when read response")
	}
	var body struct {
		Total         int      `json:"total,omitempty" structs:"total,omitempty"`
		ErrorMessages []string `json:"errorMessages,omitempty" structs:"errorMessages,omitempty"`
	}
	err = json.Unmarshal(totalByte, &body)
	if err != nil {
		return 0, errors.New("Error when unmarshaling json")
	}
	errorMessages := body.ErrorMessages
	if len(errorMessages) != 0 {
		return 0, errors.New(errorMessages[0])
	}
	total := body.Total
	return total, nil
}

func (con *Connection) getExpandIssuesJSON(projectName string, startAt int) ([]byte, error) {
	res, err := http.Get(con.url + "/rest/api/2/search?jql=project=" + projectName + "&expand=changelog&startAt=" + strconv.Itoa(startAt))
	if err != nil {
		return nil, errors.New("Error when try to get request")
	}
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, errors.New("Error when read response")
	}
	var body struct {
		ErrorMessages []string `json:"errorMessages,omitempty" structs:"errorMessages,omitempty"`
	}
	err = json.Unmarshal(resBody, &body)
	if err != nil {
		return nil, errors.New("Error when unmarshaling json")
	}
	errorMessages := body.ErrorMessages
	if len(errorMessages) != 0 {
		return nil, errors.New(errorMessages[0])
	}
	return resBody, nil
}

func (con *Connection) getAllProjectsJSON() ([]byte, error) {
	res, err := http.Get(con.url + "/rest/api/2/project")
	if err != nil {
		return nil, errors.New("Error when try to get request")
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, errors.New("Error when read response")
	}

	return resBody, nil
}

func GetConnection(baseURL string) (*Connection, error) {
	jiraClient, err := jira.NewClient(nil, baseURL)
	if err != nil {
		return nil, errors.New("Incorrect url")
	}
	_, err = http.Get(baseURL)
	if err != nil {
		return nil, errors.New("Error when try to get request")
	}
	con := new(Connection)
	con.client = *jiraClient
	con.url = baseURL
	return con, nil
}

func GetConnectionWithAuth(baseURL string, username string, password string) *Connection {
	tp := jira.BasicAuthTransport{
		Username: username,
		Password: password,
	}
	jiraClient, _ := jira.NewClient(tp.Client(), baseURL)
	con := new(Connection)
	con.client = *jiraClient
	con.url = baseURL
	return con
}

func (con *Connection) GetProjectJSON(projectName string) []byte {
	res, err := http.Get(con.url + "/rest/api/2/project/" + projectName + "?expand=description")
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		os.Exit(1)
	}
	return resBody
}

func (con *Connection) GetIssueChangelogJSON(key string) []byte {
	res, err := http.Get(con.url + "/rest/api/2/issue/" + key + "?expand=changelog&fields=key")
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		os.Exit(1)
	}
	return resBody
}

func (con *Connection) GetAllFormattedProjects() ([]datatransformer.Project, error) {
	projectsByte, err := con.getAllProjectsJSON()
	if err != nil {
		return nil, err
	}
	project, err := datatransformer.FormatProjects(projectsByte)
	if err != nil {
		return nil, err
	}

	return project, nil
}

func (con *Connection) GetFormattedIssues(projectName string) ([]datatransformer.Issue, error) {
	startAt := 0
	total, err := con.getTotalIssues(projectName)
	if err != nil {
		return nil, err
	}
	var issues []datatransformer.Issue
	for ; startAt < total; startAt += 50 {
		issuesRaw, err := con.getExpandIssuesJSON(projectName, startAt)
		if err != nil {
			return nil, err
		}
		formattedIssues, err := datatransformer.FormatIssues(issuesRaw)
		if err != nil {
			return nil, err
		}
		issues = append(issues, formattedIssues...)
	}
	return issues, nil
}
