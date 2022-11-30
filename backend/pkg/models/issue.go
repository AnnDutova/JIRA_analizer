package models

import (
	"time"
)

type Tabler interface {
	TableName() string
}

type Issue struct {
	Id          int       `json:"Id" gorm:"primarykey" validate:"required"`
	Project     *Project  `json:"Project" gorm:"embedded" validate:"required"`
	Key         string    `gorm:"unique" json:"Key" validate:"required"`
	CreatedTime time.Time `json:"CreatedTime" gorm:"column:createdtime" validate:"required"`
	ClosedTime  time.Time `json:"ClosedTime" gorm:"column:closedtime"`
	UpdatedTime time.Time `json:"UpdatedTime" gorm:"column:updatedtime" validate:"required"`
	Summary     string    `json:"Summary" validate:"required"`
	Description string    `json:"Description"`
	Type        string    `json:"Type" validate:"required"`
	Priority    string    `json:"Priority" validate:"required"`
	Status      string    `json:"Status" validate:"required"`
	Creator     string    `json:"Creator" validate:"required"`
	Assignee    string    `json:"Assignee"`
	TimeSpent   int       `json:"TimeSpent" gorm:"column:timespent"`
}
