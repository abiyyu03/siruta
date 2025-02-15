package model

import (
	"time"

	"gorm.io/gorm"
)

type Religion struct {
	*gorm.Model
	ID           int      `json:"id" gorm:"column:id;primaryKey;autoIncrement;not null"`
	ReligionName string   `json:"religion_name" validate:"required" gorm:"column:religion_name;not null"`
	Member       []Member `json:"-" gorm:"foreignKey:ReligionId"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}
