package model

import (
	"time"

	"gorm.io/gorm"
)

type RegistrationToken struct {
	*gorm.Model
	ID        int    `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	Token     string `json:"token" gorm:"column:token;unique;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
