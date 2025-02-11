package model

import (
	"time"

	"gorm.io/gorm"
)

type OutcomingLetter struct {
	*gorm.Model
	ID           string     `json:"id" gorm:"column:id;primaryKey" validate:"unique"`
	LetterNumber int        `json:"letter_number" gorm:"column:letter_number;default:0;not null"`
	Date         string     `json:"date" gorm:"column:date"`
	MemberId     string     `json:"member_id" gorm:"column:member_id;not null"`
	Member       Member     `gorm:"foreignKey:MemberId;not null"`
	LetterTypeId uint       `json:"letter_type_id" gorm:"column:letter_type_id;not null"`
	LetterType   LetterType `gorm:"foreignKey:LetterTypeId;not null"`
	RTProfileId  string     `json:"rt_profile_id" gorm:"column:rt_profile_id;not null"`
	RTProfile    RTProfile  `gorm:"foreignKey:RTProfileId;not null"`
	IsRTApproved bool       `json:"is_rt_approved" gorm:"column:is_rt_approved;default:false"`
	// IsRWApproved bool      `json:"is_rw_approved" gorm:"column:is_rw_approved;default:false"`
	Description string `json:"description" gorm:"column:description"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
