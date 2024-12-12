package model

import (
	"time"

	"gorm.io/gorm"
)

type MemberStatus struct {
	*gorm.Model
	ID     int    `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	Status string `json:"status" validate:"required" gorm:"column:status"`
	// Members   Member `gorm:"foreignKey:MemberStatusId"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
