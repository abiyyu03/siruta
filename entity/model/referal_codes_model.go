package model

import (
	"time"

	"gorm.io/gorm"
)

type ReferalCode struct {
	*gorm.Model
	ID        int       `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	Code      string    `json:"code" gorm:"uniqueIndex;not null" validate:"required"`
	ExpiredAt time.Time `json:"expired_at"`
	IsExpired bool      `json:"is_expired" gorm:"default:false;not null"`
	ProfileId string    `json:"profile_id" gorm:"column:profile_id;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
