package datatransformer

type JiraIssue struct {
	Id        string       `json:"id,omitempty" structs:"id,omitempty"`
	Key       string       `json:"key,omitempty" structs:"key,omitempty"`
	Fields    *IssueFields `json:"fields,omitempty" structs:"fields,omitempty"`
	Changelog *Changelog   `json:"changelog,omitempty" structs:"changelog,omitempty"`
}

type IssueFields struct {
	Type        IssueType   `json:"issuetype,omitempty" structs:"issuetype,omitempty"`
	Priority    *Priority   `json:"priority,omitempty" structs:"priority,omitempty"`
	Project     JiraProject `json:"project,omitempty" structs:"project,omitempty"`
	Description string      `json:"description,omitempty" structs:"description,omitempty"`
	Created     string      `json:"created,omitempty" structs:"created,omitempty"`
	Updated     string      `json:"updated,omitempty" structs:"updated,omitempty"`
	Summary     string      `json:"summary,omitempty" structs:"summary,omitempty"`
	Creator     *User       `json:"Creator,omitempty" structs:"Creator,omitempty"`
	Assignee    *User       `json:"assignee,omitempty" structs:"assignee,omitempty"`
	Status      *Status     `json:"status,omitempty" structs:"status,omitempty"`
	TimeSpent   int         `json:"timespent,omitempty" structs:"timespent,omitempty"`
}

type JiraProject struct {
	Expand       string            `json:"expand,omitempty" structs:"expand,omitempty"`
	Self         string            `json:"self,omitempty" structs:"self,omitempty"`
	ID           string            `json:"id,omitempty" structs:"id,omitempty"`
	Key          string            `json:"key,omitempty" structs:"key,omitempty"`
	Description  string            `json:"description,omitempty" structs:"description,omitempty"`
	Lead         User              `json:"lead,omitempty" structs:"lead,omitempty"`
	IssueTypes   []IssueType       `json:"issueTypes,omitempty" structs:"issueTypes,omitempty"`
	URL          string            `json:"url,omitempty" structs:"url,omitempty"`
	Email        string            `json:"email,omitempty" structs:"email,omitempty"`
	AssigneeType string            `json:"assigneeType,omitempty" structs:"assigneeType,omitempty"`
	Name         string            `json:"name,omitempty" structs:"name,omitempty"`
	Roles        map[string]string `json:"roles,omitempty" structs:"roles,omitempty"`
}

type Status struct {
	Self        string `json:"self" structs:"self"`
	Description string `json:"description" structs:"description"`
	IconURL     string `json:"iconUrl" structs:"iconUrl"`
	Name        string `json:"name" structs:"name"`
	ID          string `json:"id" structs:"id"`
}

type Priority struct {
	Self        string `json:"self,omitempty" structs:"self,omitempty"`
	IconURL     string `json:"iconUrl,omitempty" structs:"iconUrl,omitempty"`
	Name        string `json:"name,omitempty" structs:"name,omitempty"`
	ID          string `json:"id,omitempty" structs:"id,omitempty"`
	StatusColor string `json:"statusColor,omitempty" structs:"statusColor,omitempty"`
	Description string `json:"description,omitempty" structs:"description,omitempty"`
}

type IssueType struct {
	Self        string `json:"self,omitempty" structs:"self,omitempty"`
	ID          string `json:"id,omitempty" structs:"id,omitempty"`
	Description string `json:"description,omitempty" structs:"description,omitempty"`
	IconURL     string `json:"iconUrl,omitempty" structs:"iconUrl,omitempty"`
	Name        string `json:"name,omitempty" structs:"name,omitempty"`
	Subtask     bool   `json:"subtask,omitempty" structs:"subtask,omitempty"`
	AvatarID    int    `json:"avatarId,omitempty" structs:"avatarId,omitempty"`
}

type ChangelogHistory struct {
	Id      string           `json:"id" structs:"id"`
	Author  User             `json:"author" structs:"author"`
	Created string           `json:"created" structs:"created"`
	Items   []ChangelogItems `json:"items" structs:"items"`
}

type ChangelogItems struct {
	Field      string      `json:"field" structs:"field"`
	FieldType  string      `json:"fieldtype" structs:"fieldtype"`
	From       interface{} `json:"from" structs:"from"`
	FromString string      `json:"fromString" structs:"fromString"`
	To         interface{} `json:"to" structs:"to"`
	ToString   string      `json:"toString" structs:"toString"`
}

type User struct {
	Self            string   `json:"self,omitempty" structs:"self,omitempty"`
	AccountID       string   `json:"accountId,omitempty" structs:"accountId,omitempty"`
	AccountType     string   `json:"accountType,omitempty" structs:"accountType,omitempty"`
	Name            string   `json:"name,omitempty" structs:"name,omitempty"`
	Key             string   `json:"key,omitempty" structs:"key,omitempty"`
	Password        string   `json:"-"`
	EmailAddress    string   `json:"emailAddress,omitempty" structs:"emailAddress,omitempty"`
	DisplayName     string   `json:"displayName,omitempty" structs:"displayName,omitempty"`
	Active          bool     `json:"active,omitempty" structs:"active,omitempty"`
	TimeZone        string   `json:"timeZone,omitempty" structs:"timeZone,omitempty"`
	Locale          string   `json:"locale,omitempty" structs:"locale,omitempty"`
	ApplicationKeys []string `json:"applicationKeys,omitempty" structs:"applicationKeys,omitempty"`
}

type Changelog struct {
	Histories []ChangelogHistory `json:"histories,omitempty"`
}
