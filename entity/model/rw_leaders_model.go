package model

import (
	"time"

	"gorm.io/gorm"
)

type RWLeader struct {
	*gorm.Model
	ID          string    `json:"id" gorm:"column:id;column:id;primaryKey;not null"`
	Fullname    string    `json:"fullname" gorm:"column:fullname"`
	NikNumber   string    `json:"nik_number" gorm:"column:nik_number;uniqueIndex" validate:"required,unique"`
	KKNumber    string    `json:"kk_number" gorm:"column:kk_number"`
	Photo       *string   `json:"photo" gorm:"column:photo"`
	RWProfileId string    `json:"rw_profile_id" gorm:"not null"`
	RWProfile   RWProfile `json:"-" gorm:"foreignKey:RWProfileId;not null"`
	UserId      string    `json:"user_id" gorm:"column:user_id"`
	User        User      `json:"-" gorm:"foreignKey:UserId;not null"`
	FullAddress string    `json:"full_address" gorm:"column:full_address;type:text"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
