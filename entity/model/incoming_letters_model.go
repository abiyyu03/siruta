package model

import (
	"time"

	"gorm.io/gorm"
)

type IncomingLetter struct {
	*gorm.Model
	ID           int       `json:"id" gorm:"column:id;primaryKey;autoIncrement;not null"`
	Title        string    `json:"title" gorm:"column:title;not null" validate:"required"`
	LetterDate   time.Time `json:"letter_date" gorm:"column:letter_date;not null" validate:"required"`
	OriginLetter string    `json:"origin_letter" gorm:"column:origin_letter;not null" validate:"required"`
	RTProfileId  string    `json:"rt_profile_id" gorm:"column:rt_profile_id;not null" validate:"required"`
	RTProfile    RTProfile `gorm:"foreignKey:RTProfileId;not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}
