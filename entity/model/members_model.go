package model

import (
	"time"

	"gorm.io/gorm"
)

type Member struct {
	*gorm.Model
	ID            int       `json:"id" gorm:"primaryKey"`
	FirstName     string    `json:"first_name" validate:"required, max=32"`
	LastName      string    `json:"last_name" validate:"required, max=32"`
	NikNumber     string    `json:"nik_number"`
	KKNumber      string    `json:"kk_number"`
	BornPlace     string    `json:"born_place"`
	Birthdate     time.Time `json:"birthdate"`
	Gender        string    `json:"gender"`
	HomeAddress   string    `json:"home_address"`
	MartialStatus string    `json:"martial_status"`
	ReligionId    uint      `json:"religion_id"`
	Religion      Religion  `gorm:"foreignKey:ReligionId"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}
