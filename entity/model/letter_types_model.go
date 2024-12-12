package model

import (
	"time"

	"gorm.io/gorm"
)

type LetterType struct {
	*gorm.Model
	ID        int    `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	TypeName  string `json:"type_name" validate:"required" gorm:"not null"`
	Code      string `json:"code" validate:"required;max=6" gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
