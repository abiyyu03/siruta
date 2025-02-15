package model

import (
	"time"

	"gorm.io/gorm"
)

type LetterType struct {
	*gorm.Model
	ID                 int               `json:"id" gorm:"column:id;primaryKey;autoIncrement;not null"`
	TypeName           string            `json:"type_name" validate:"required" gorm:"not null"`
	Code               string            `json:"code" validate:"required;max=6" gorm:"not null"`
	OutcomingLetter    []OutcomingLetter `json:"-" gorm:"foreignKey:LetterTypeId"`
	IsForLocalResident bool              `json:"is_for_local_resident" gorm:"column:is_for_local_resident;default:true"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
	DeletedAt          gorm.DeletedAt `gorm:"index"`
}
