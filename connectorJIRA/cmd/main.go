package main

import (
	"connectorJIRA/pkg/connector"
	"connectorJIRA/pkg/datatransformer"
	"connectorJIRA/pkg/properties"
	"os"
)

func main() {
	config := properties.GetConfig(os.Args[1])
	url := config.ProgramSettings.ApacheUrl
	connectionApache := connector.GetConnection(url)

	issuesRaw := connectionApache.GetExpandIssuesJSON(config.ProgramSettings.ProjectNames)
	issues := datatransformer.FormatIssues(issuesRaw)
	JSON := "["
	for _, issue := range issues {
		JSON += issue.ToJSON() + ","
		changelog := connectionApache.GetIssueChangelogJSON(issue.Key)
		statusChanges := datatransformer.FormatChangelog(changelog)
		JSON += statusChanges.ToJSON() + ","
	}
	JSON += "]"
	datatransformer.ToFile(JSON, "Agila_issues")

	//datapusher.PushIssues(issues)

}
