package models

type StatusChange struct {
	Author     string `json:"Author" validate:"required"`
	ChangeTime string `json:"ChangeTime" gorm:"column:changetime" validate:"required"`
	FromStatus string `json:"FromStatus" gorm:"column:fromstatus"`
	ToStatus   string `json:"ToStatus" gorm:"column:tostatus" validate:"required"`
}
