package model

import (
	"time"

	"gorm.io/gorm"
)

type ResetPassword struct {
	*gorm.Model
	Token     string    `json:"token" gorm:"column:token;primaryKey;unique;not null"`
	UserID    string    `json:"user_id" gorm:"not null"`
	User      User      `json:"-" gorm:"foreignKey:UserID;not null"`
	ExpiredAt time.Time `json:"expired_at" gorm:"expired_at;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
