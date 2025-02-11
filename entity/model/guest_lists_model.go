package model

import (
	"time"

	"gorm.io/gorm"
)

type GuestList struct {
	*gorm.Model
	ID          int       `json:"id" gorm:"column:id;primaryKey;autoIncrement;not null"`
	FullName    string    `json:"full_name" validate:"required" gorm:"column:full_name;not null"`
	PhoneNumber string    `json:"phone_number" validate:"required" gorm:"column:phone_number;not null"`
	VisitAt     time.Time `json:"visit_at" gorm:"column:visit_at;not null"`
	RTProfileId string    `json:"rt_profile_id" gorm:"column:rt_profile_id;not null" validate:"required"`
	RTProfile   RTProfile `gorm:"foreignKey:RTProfileId"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
