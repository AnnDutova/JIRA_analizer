package connector

import (
	"fmt"
	"github.com/andygrunwald/go-jira"
	"io"
	"net/http"
	"os"
)

type Connection struct {
	client jira.Client
	url    string
}

func GetConnection(baseURL string) *Connection {
	jiraClient, _ := jira.NewClient(nil, baseURL)
	con := new(Connection)
	con.client = *jiraClient
	con.url = baseURL
	return con
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

func (con *Connection) GetAllIssuesFromProject(projectName string) []jira.Issue {
	jql := "project = " + projectName
	issues, _, _ := con.client.Issue.Search(jql, nil)
	return issues
}

func (con *Connection) GetClosedIssuesFromProject(projectName string) []jira.Issue {
	jql := "project = " + projectName + " AND status = Closed"
	issues, _, _ := con.client.Issue.Search(jql, nil)
	return issues
}

func (con *Connection) GetExpandIssuesJSON(projectName string) []byte {
	res, err := http.Get(con.url + "/rest/api/2/search?jql=project=" + projectName + "&expand=changelog")
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
