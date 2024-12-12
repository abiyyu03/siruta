package model

import (
	"time"

	"gorm.io/gorm"
)

type RWLeader struct {
	*gorm.Model
	ID          string    `json:"id" gorm:"column:id;column:id;primaryKey;not null"`
	MemberId    uint      `json:"member_id" gorm:"not null"`
	Member      Member    `json:"-" gorm:"foreignKey:MemberId;not null"`
	RWProfileId uint      `json:"rw_profile_id" gorm:"not null"`
	RWProfile   RWProfile `json:"-" gorm:"foreignKey:RWProfileId;not null"`
	IsActive    bool      `json:"is_active" gorm:"not null"`
	StartPeriod int       `json:"start_period" gorm:"not null"`
	EndPeriod   int       `json:"end_period" gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
