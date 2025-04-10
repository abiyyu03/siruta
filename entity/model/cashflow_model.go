package model

import (
	"time"

	"gorm.io/gorm"
)

type Cashflow struct {
	*gorm.Model
	ID                 int       `json:"id" gorm:"column:id;primaryKey;autoIncrement;not null"`
	Description        string    `json:"description" gorm:"column:description"`
	LogType            string    `json:"log_type" gorm:"column:log_type"`
	Amount             float64   `json:"amount" gorm:"column:amount;not null" validate:"required"`
	PaymentDate        time.Time `json:"payment_date" gorm:"column:payment_date;not null"`
	RTProfileId        string    `json:"rt_profile_id" gorm:"column:rt_profile_id;not null" validate:"required"`
	RTProfile          RTProfile `json:"-" gorm:"foreignKey:RTProfileId"`
	PaymentPeriodYear  string    `json:"payment_period_year" gorm:"column:payment_period_year;not null" validate:"required"`
	PaymentPeriodMonth string    `json:"payment_period_month" gorm:"column:payment_period_month;not null" validate:"required"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
	DeletedAt          gorm.DeletedAt `gorm:"index"`
}
