package model

import (
	"time"

	"gorm.io/gorm"
)

type RTLeader struct {
	*gorm.Model
	ID          int       `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	MemberId    uint      `json:"member_id" gorm:"column:member_id;not null"`
	Member      Member    `gorm:"foreignKey:MemberId;not null"`
	RTProfileId uint      `json:"rt_profile_id" gorm:"column:rt_profile_id;not null"`
	RTProfile   RTProfile `gorm:"foreignKey:RTProfileId;not null"`
	IsActive    bool      `json:"is_active" gorm:"column:is_active;not null"`
	StartPeriod int       `json:"start_period" gorm:"column:start_period;not null"`
	EndPeriod   int       `json:"end_period" gorm:"column:end_period;not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
