package model

import (
	"time"

	"gorm.io/gorm"
)

type RTLeader struct {
	*gorm.Model
	ID          string    `json:"id" gorm:"column:id;primaryKey;not null"`
	Fullname    string    `json:"fullname" gorm:"column:fullname"`
	NikNumber   string    `json:"nik_number" gorm:"column:nik_number;uniqueIndex" validate:"required,unique"`
	KKNumber    string    `json:"kk_number" gorm:"column:kk_number"`
	RTProfileId string    `json:"rt_profile_id" gorm:"column:rt_profile_id;not null"`
	RTProfile   RTProfile `gorm:"foreignKey:RTProfileId;not null"`
	UserId      string    `json:"user_id" gorm:"column:user_id"`
	User        User      `json:"-" gorm:"foreignKey:UserId;not null"`
	FullAddress string    `json:"full_address" gorm:"column:full_address;type:text"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
