package model

import (
	"time"

	"gorm.io/gorm"
)

type Religion struct {
	*gorm.Model
	ID           int      `json:"id" gorm:"column:id;primaryKey;autoIncrement;not null"`
	ReligionName string   `json:"religion_name" validate:"column:religion_name;required" gorm:"not null"`
	Member       []Member `gorm:"foreignKey:ReligionId"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}
