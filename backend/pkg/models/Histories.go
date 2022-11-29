package models

type Histories struct {
	IssueId   int            `json:"Id" validate:"required"`
	Histories []StatusChange `json:"Histories" validate:"required"`
}
