package datapusher

import (
	"connectorJIRA/pkg/datatransformer"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"strings"
	"time"
)

type PSQLConnector struct {
	db *sql.DB
}

func New(db_ *sql.DB) (*PSQLConnector, error) {
	return &PSQLConnector{
		db: db_,
	}, nil
}

func (c *PSQLConnector) UpdateData(issues []datatransformer.Issue, statuses []datatransformer.IssueStatusChanges) error {
	if err := c.pushIssues(issues); err != nil {
		return errors.New("Error when pushIssues: " + err.Error())
	}
	if err := c.pushStatusChanges(statuses); err != nil {
		return errors.New("Error when pushStatusChanges: " + err.Error())
	}
	return nil
}

// UpdateIssue redundant function
func (c *PSQLConnector) UpdateIssue(issue datatransformer.Issue, statuses datatransformer.IssueStatusChanges) error {
	command := fmt.Sprintf("select updateIssue('%d','%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s')",
		issue.Id, issue.Assignee, issue.UpdatedTime, issue.ClosedTime,
		strings.ReplaceAll(issue.Summary, "'", " "), strings.ReplaceAll(issue.Description, "'", " "), issue.Type, issue.Priority, issue.Status)
	var oldDate time.Time
	if err := c.db.QueryRow(command).Scan(&oldDate); err != nil {
		return errors.New("Error when exec command \"" + command + "\": " + err.Error())
	}
	for _, val := range statuses.Histories {
		if oldDate.After(val.ChangeTime) {
			command = fmt.Sprintf("call insertStatusChange('%s','%d','%s','%s','%s')", val.Author, issue.Id,
				val.ChangeTime.Format(time.RFC3339), val.FromStatus, val.ToStatus)
			_, err := c.db.Exec(command)
			if err != nil {
				return errors.New("(loop) Error when exec command \"" + command + "\": " + err.Error())
			}
		}
	}
	return nil
}

func (c *PSQLConnector) pushIssues(issues []datatransformer.Issue) error {
	for _, val := range issues {

		command := fmt.Sprintf("call insertOrUpdateIssue('%d','%s','%d','%s','%s','%s','%s','%s','%s', '%s', '%s', '%s', '%s', '%s', '%d')",
			val.Id, strings.ReplaceAll(val.Project.Name, "'", " "), val.Project.Id, strings.ReplaceAll(val.Creator, "'", " "), strings.ReplaceAll(val.Assignee, "'", " "), val.Key,
			val.CreatedTime.Format(time.RFC3339), val.ClosedTime.Format(time.RFC3339), val.UpdatedTime.Format(time.RFC3339),
			strings.ReplaceAll(val.Summary, "'", " "), strings.ReplaceAll(val.Description, "'", " "), val.Type, val.Priority, val.Status,
			val.TimeSpent)
		_, err := c.db.Exec(command)
		if err != nil {
			return errors.New("Error when exec command \"" + command + "\": " + err.Error())
		}
	}
	return nil
}

func (c *PSQLConnector) pushStatusChanges(statuses []datatransformer.IssueStatusChanges) error {
	for _, val := range statuses {
		command := fmt.Sprintf("select getLastChangeTime('%d')", val.Id)
		var oldDate time.Time
		if err := c.db.QueryRow(command).Scan(&oldDate); err != nil {
			return errors.New("Error when exec command \"" + command + "\": " + err.Error())
		}
		for _, history := range val.Histories {
			if history.ChangeTime.After(oldDate) {
				command := fmt.Sprintf("call insertStatusChange('%s','%d','%s','%s','%s')", strings.ReplaceAll(history.Author, "'", " "), val.Id,
					history.ChangeTime.Format(time.RFC3339), history.FromStatus, history.ToStatus)
				_, err := c.db.Exec(command)
				if err != nil {
					return errors.New("(loop) Error when exec command \"" + command + "\": " + err.Error())
				}
			}
		}
	}
	return nil
}
