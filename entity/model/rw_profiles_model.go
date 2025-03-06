package model

import (
	"time"

	"gorm.io/gorm"
)

type RWProfile struct {
	*gorm.Model
	ID           string      `json:"id" gorm:"column:id;primaryKey;not null"`
	RWNumber     string      `json:"rw_number" gorm:"column:rw_number;not null" validate:"required"`
	VillageID    uint        `json:"village_id" gorm:"column:village_id;not null"`
	Village      Village     `json:"-" gorm:"foreignKey:VillageID;not null" validate:"omitempty,dive"`
	RWLogo       *string     `json:"rw_logo" gorm:"column:rw_logo"`
	Latitude     float64     `json:"latitude" gorm:"column:latitude;not null"`
	Longitude    float64     `json:"longitude" gorm:"column:longitude;not null"`
	IsAuthorized bool        `json:"is_authorized" gorm:"column:is_authorized;default:false;not null"`
	RwEmail      string      `json:"rw_email" gorm:"column:rw_email;uniqueIndex;not null" validate:"required,email,unique"`
	MobilePhone  string      `json:"mobile_phone" gorm:"column:mobile_phone;uniqueIndex;not null" validate:"required,unique"`
	RegencyLogo  *string     `json:"regency_logo" gorm:"column:regency_logo"`
	FullAddress  string      `json:"full_address" gorm:"column:full_address;type:text"`
	RTProfiles   []RTProfile `json:"rt_profiles"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}
