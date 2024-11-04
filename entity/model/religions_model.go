package model

import (
	"time"

	"gorm.io/gorm"
)

type Religion struct {
	*gorm.Model
	ID           int    `json:"id" gorm:"primaryKey"`
	ReligionName string `json:"religion_name"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}
