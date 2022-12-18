package models

type Project struct {
	Id    int    `json:"Id" gorm:"primarykey"  validate:"required"`
	Key   string `json:"Key" default:""`
	Title string `json:"Name" gorm:"unique" validate:"required"`
	Url   string `json:"Url"`
}

type ProjectWithExistence struct {
	Id        int    `json:"Id" gorm:"primarykey"  validate:"required"`
	Key       string `json:"Key" default:""`
	Title     string `json:"Name" gorm:"unique" validate:"required"`
	Url       string `json:"Url"`
	Existence bool   `json:"Existence"`
}

type Projects struct {
	Projects []*ProjectWithExistence `json:"Projects"`
	PageInfo PageInfo                `json:"PageInfo"`
}

type ProjectAnalytic struct {
	Id                    int    `json:"Id"`
	Key                   string `json:"Key"`
	Title                 string `json:"Name"`
	AllIssuesCount        int    `json:"allIssuesCount"`
	OpenIssuesCount       int    `json:"openIssuesCount"`
	CloseIssuesCount      int    `json:"closeIssuesCount"`
	ResolvedIssuesCount   int    `json:"resolvedIssuesCount"`
	ReopenedIssuesCount   int    `json:"reopenedIssuesCount"`
	InProgressIssuesCount int    `json:"progressIssuesCount"`
	AverageTime           int    `json:"averageTime"`
	AverageIssuesCount    int    `json:"averageIssuesCount"`
}

type PageCount struct {
	PageCount int `json:"PageCount"`
}

type PageInfo struct {
	CurrentPage   uint `json:"currentPage"`
	PageCount     uint `json:"pageCount"`
	ProjectsCount uint `json:"projectsCount"`
}
