package model

import (
	"time"

	"gorm.io/gorm"
)

type RTProfile struct {
	*gorm.Model
	ID               string            `json:"id" gorm:"column:id;column:id;primaryKey;not null"`
	RtNumber         int               `json:"rt_number" gorm:"column:rt_number;not null" validate:"required"`
	RtLogo           *string           `json:"rt_logo" gorm:"column:rt_logo"`
	Latitude         float64           `json:"latitude" gorm:"column:latitude;not null"`
	Longitude        float64           `json:"longitude" gorm:"column:longitude;not null"`
	IsAuthorized     bool              `json:"is_authorized" gorm:"column:is_authorized;default:false;not null"`
	RtEmail          string            `json:"rt_email" gorm:"column:rt_email;uniqueIndex;not null" validate:"required,unique,email"`
	MobilePhone      string            `json:"mobile_phone" gorm:"column:mobile_phone;uniqueIndex;not null" validate:"required,unique"`
	Description      *string           `json:"description" gorm:"column:description;type:text"`
	RWProfileId      string            `json:"rw_profile_id" gorm:"column:rw_profile_id;not null"`
	RWProfile        RWProfile         `json:"-" gorm:"foreignKey:RWProfileId;not null" validate:"omitempty"`
	IncomingLetters  []IncomingLetter  `json:"-" gorm:"foreignKey:RTProfileId"`
	OutcomingLetters []OutcomingLetter `json:"-" gorm:"foreignKey:RTProfileId"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt `gorm:"index"`
}
