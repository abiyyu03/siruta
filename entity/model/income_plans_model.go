package model

import (
	"time"

	"gorm.io/gorm"
)

type IncomePlan struct {
	*gorm.Model
	ID            string    `json:"id" gorm:"column:id;column:id;primaryKey;not null" validate:"required"`
	PlanName      string    `json:"plan_name" gorm:"column:plan_name;not null" validate:"required"`
	RTProfileId   string    `json:"rt_profile_id" gorm:"column:rt_profile_id;not null" validate:"required"`
	RTProfile     RTProfile `json:"-" gorm:"foreignKey:RTProfileId"`
	StartPlan     time.Time `json:"start_plan" gorm:"column:start_plan"`
	EndPlan       time.Time `json:"end_plan" gorm:"column:end_plan"`
	IsClosed      bool      `json:"is_closed" gorm:"column:is_closed;default:false"`
	Description   string    `json:"description" gorm:"column:description"`
	IsSetDeadline bool      `json:"is_set_deadline" gorm:"column:is_set_deadline;default:false"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}
