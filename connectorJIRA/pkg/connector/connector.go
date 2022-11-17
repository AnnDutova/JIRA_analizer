package connector

import (
	"connectorJIRA/pkg/datatransformer"
	"connectorJIRA/pkg/properties"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/andygrunwald/go-jira"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Connection struct {
	client       jira.Client
	url          string
	maxTimeSleep uint
	minTimeSleep uint
}

func GetConnection(projectPath string) (*Connection, error) {
	config := properties.GetConfig(projectPath)
	jiraClient, err := jira.NewClient(nil, config.ProgramSettings.JiraUrl)
	if err != nil {
		return nil, errors.New("Incorrect url")
	}
	_, err = http.Get(config.ProgramSettings.JiraUrl)
	if err != nil {
		return nil, errors.New("Error when try to get request in GetConnection: " + err.Error())
	}
	con := new(Connection)
	con.client = *jiraClient
	con.url = config.ProgramSettings.JiraUrl
	con.maxTimeSleep = config.ProgramSettings.MaxTimeSleep
	con.minTimeSleep = config.ProgramSettings.MinTimeSleep
	return con, nil
}

func (con *Connection) GetTotalIssues(projectKey string) (int, error) {
	maxTimeSleepMillsec := con.maxTimeSleep
	timeSleepMillsec := con.minTimeSleep
	needWait := false
CON:
	if needWait {
		time.Sleep(time.Duration(timeSleepMillsec) * time.Millisecond)
	}
	res, err := http.Get(con.url + "/rest/api/2/search?jql=project=\"" + projectKey + "\"&expand=changelog")
	if err != nil {
		needWait = true
		if timeSleepMillsec < maxTimeSleepMillsec {
			timeSleepMillsec *= 2
			if timeSleepMillsec > maxTimeSleepMillsec {
				timeSleepMillsec = maxTimeSleepMillsec
			}
			goto CON
		}
		return 0, errors.New("Error when try to get request in GetTotalIssues: " + err.Error())
	}
	totalByte, err := io.ReadAll(res.Body)
	if err != nil {
		needWait = true
		if timeSleepMillsec < maxTimeSleepMillsec {
			timeSleepMillsec *= 2
			if timeSleepMillsec > maxTimeSleepMillsec {
				timeSleepMillsec = maxTimeSleepMillsec
			}
			goto CON
		}
		return 0, errors.New("Error when read response in GetTotalIssues: " + err.Error())
	}
	var body struct {
		Total         int      `json:"total,omitempty" structs:"total,omitempty"`
		ErrorMessages []string `json:"errorMessages,omitempty" structs:"errorMessages,omitempty"`
	}
	err = json.Unmarshal(totalByte, &body)
	if err != nil {
		return 0, errors.New("Error when unmarshaling json in GetTotalIssues: " + err.Error())
	}
	errorMessages := body.ErrorMessages
	if len(errorMessages) != 0 {
		return 0, errors.New(errorMessages[0])
	}
	total := body.Total
	return total, nil
}

func (con *Connection) GetExpandIssuesJSON(projectKey string, startAt int, maxResults int) ([]byte, error) {
	maxTimeSleepMillsec := con.maxTimeSleep
	timeSleepMillsec := con.minTimeSleep
	needWait := false
CON:
	if needWait {
		time.Sleep(time.Duration(timeSleepMillsec) * time.Millisecond)
	}
	res, err := http.Get(con.url + "/rest/api/2/search?jql=project=\"" + projectKey + "\"&expand=changelog&startAt=" +
		strconv.Itoa(startAt) + "&maxResults=" + strconv.Itoa(maxResults))
	if err != nil {
		needWait = true
		if timeSleepMillsec < maxTimeSleepMillsec {
			timeSleepMillsec *= 2
			if timeSleepMillsec > maxTimeSleepMillsec {
				timeSleepMillsec = maxTimeSleepMillsec
			}
			goto CON
		}
		return nil, errors.New("Error when try to get request in GetExpandIssuesJSON: " + err.Error())
	}
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		needWait = true
		if timeSleepMillsec < maxTimeSleepMillsec {
			timeSleepMillsec *= 2
			if timeSleepMillsec > maxTimeSleepMillsec {
				timeSleepMillsec = maxTimeSleepMillsec
			}
			goto CON
		}
		return nil, errors.New("Error when read response in GetExpandIssuesJSON: " + err.Error())
	}
	var body struct {
		ErrorMessages []string `json:"errorMessages,omitempty" structs:"errorMessages,omitempty"`
	}
	err = json.Unmarshal(resBody, &body)
	if err != nil {
		return nil, errors.New("Error when unmarshaling json in GetExpandIssuesJSON: " + err.Error())
	}
	errorMessages := body.ErrorMessages
	if len(errorMessages) != 0 {
		return nil, errors.New(errorMessages[0])
	}
	return resBody, nil
}

func (con *Connection) getAllProjectsJSON() ([]byte, error) {
	maxTimeSleepMillsec := con.maxTimeSleep
	timeSleepMillsec := con.minTimeSleep
	needWait := false
CON:
	if needWait {
		time.Sleep(time.Duration(timeSleepMillsec) * time.Millisecond)
	}
	res, err := http.Get(con.url + "/rest/api/2/project")
	if err != nil {
		needWait = true
		if timeSleepMillsec < maxTimeSleepMillsec {
			timeSleepMillsec *= 2
			if timeSleepMillsec > maxTimeSleepMillsec {
				timeSleepMillsec = maxTimeSleepMillsec
			}
			goto CON
		}
		return nil, errors.New("Error when try to get request in getAllProjectsJSON: " + err.Error())
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		needWait = true
		if timeSleepMillsec < maxTimeSleepMillsec {
			timeSleepMillsec *= 2
			if timeSleepMillsec > maxTimeSleepMillsec {
				timeSleepMillsec = maxTimeSleepMillsec
			}
			goto CON
		}
		return nil, errors.New("Error when read response in getAllProjectsJSON: " + err.Error())
	}

	return resBody, nil
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

func (con *Connection) GetProjectJSON(projectKey string) []byte {
	res, err := http.Get(con.url + "/rest/api/2/project/" + projectKey + "?expand=description")
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

func (con *Connection) GetIssueChangelogJSON(key string) ([]byte, error) {
	maxTimeSleepMillsec := con.maxTimeSleep
	timeSleepMillsec := con.minTimeSleep
	needWait := false
CON:
	if needWait {
		time.Sleep(time.Duration(timeSleepMillsec) * time.Millisecond)
	}
	res, err := http.Get(con.url + "/rest/api/2/issue/" + key + "?expand=changelog&fields=key")
	if err != nil {
		needWait = true
		if timeSleepMillsec < maxTimeSleepMillsec {
			timeSleepMillsec *= 2
			if timeSleepMillsec > maxTimeSleepMillsec {
				timeSleepMillsec = maxTimeSleepMillsec
			}
			goto CON
		}
		return nil, errors.New("Error while try to get request in GetIssueChangelogJSON: " + err.Error())
	}
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		needWait = true
		if timeSleepMillsec < maxTimeSleepMillsec {
			timeSleepMillsec *= 2
			if timeSleepMillsec > maxTimeSleepMillsec {
				timeSleepMillsec = maxTimeSleepMillsec
			}
			goto CON
		}
		return nil, errors.New("Error while try to read response body GetIssueChangelogJSON: " + err.Error())
	}
	return resBody, nil
}

func (con *Connection) GetAllFormattedProjects(limit int, page int, search string) (*datatransformer.ProjectsRespond, error) {
	projectsByte, err := con.getAllProjectsJSON()
	if err != nil {
		return nil, err
	}
	projects, err := datatransformer.FormatProjectsRespond(projectsByte, limit, page, search)
	if err != nil {
		return nil, err
	}

	return projects, nil
}

func (con *Connection) GetFormattedIssues(projectName string) ([]datatransformer.Issue, error) {
	startAt := 0
	total, err := con.GetTotalIssues(projectName)
	if err != nil {
		return nil, err
	}
	var issues []datatransformer.Issue
	for ; startAt < total; startAt += 50 {
		issuesRaw, err := con.GetExpandIssuesJSON(projectName, startAt, 50)
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
