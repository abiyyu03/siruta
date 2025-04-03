package model

import (
	"time"

	"gorm.io/gorm"
)

type Expense struct {
	*gorm.Model
	ID             int       `json:"id" gorm:"column:id;primaryKey;autoIncrement;not null"`
	Amount         float64   `json:"amount_total" gorm:"column:amount;not null" validate:"required"`
	Description    string    `json:"description" gorm:"column:description;not null" validate:"required"`
	PaymentDate    time.Time `json:"payment_date" gorm:"column:date;not null"`
	AdditionalNote *string   `json:"additional_note" gorm:"column:additional_note"`
	Quantity       int       `json:"quantity" gorm:"column:quantity;default:1"`
	RTProfileId    string    `json:"rt_profile_id" gorm:"column:rt_profile_id;not null" validate:"required"`
	RTProfile      RTProfile `json:"-" gorm:"foreignKey:RTProfileId"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}
