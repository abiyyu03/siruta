package model

import (
	"time"

	"gorm.io/gorm"
)

type Income struct {
	*gorm.Model
	ID            int        `json:"id" gorm:"column:id;primaryKey;autoIncrement;not null"`
	Amount        float64    `json:"amount" gorm:"column:amount;not null" validate:"required"`
	PaymentDate   time.Time  `json:"payment_date" gorm:"column:date;not null"`
	PlanId        string     `json:"plan_id" gorm:"column:plan_id;not null" validate:"required"`
	IncomePlan    IncomePlan `json:"-" gorm:"foreignKey:PlanId"`
	PaymentMethod string     `json:"payment_method" gorm:column:payment_method;not null" validate:"required"`
	PlanPeriod    string     `json:"plan_period" gorm:"column:plan_period;not null" validate:"required"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}
