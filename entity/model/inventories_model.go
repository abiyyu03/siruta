package model

import (
	"time"

	"gorm.io/gorm"
)

type Inventory struct {
	*gorm.Model
	ID          int       `json:"id" gorm:"column:id;primaryKey;autoIncrement;not null"`
	Name        string    `json:"name" gorm:"column:name;not null"`
	Quantity    int       `json:"quantity" gorm:"column:quantity; not null"`
	Image       *string   `json:"image" gorm:"column:image"`
	RTProfileId string    `json:"rt_profile_id" gorm:"column:rt_profile_id;not null" validate:"required"`
	RTProfile   RTProfile `gorm:"foreignKey:RTProfileId"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
