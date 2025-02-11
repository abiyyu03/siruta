package model

import (
	"time"

	"gorm.io/gorm"
)

type Village struct {
	*gorm.Model
	ID         int     `json:"id" gorm:"column:id;primaryKey;autoIncrement;not null"`
	Name       string  `json:"name" gorm:"column:name;not null" validate:"required"`
	AltName    string  `json:"alt_name" gorm:"column:alt_name;not null" validate:"required"`
	Latitude   float64 `json:"latitude" gorm:"column:latitude;not null" validate:"required"`
	Longitude  float64 `json:"longitude" gorm:"column:longitude;not null" validate:"required"`
	CodePostal string  `json:"code_postal" gorm:"column:code_postal;not null" validate:"required"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}
