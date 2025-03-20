package model

import (
	"time"

	"gorm.io/gorm"
)

type Member struct {
	*gorm.Model
	ID              string            `json:"id" gorm:"column:id;column:id;primaryKey;not null" validate:"required"`
	Fullname        string            `json:"fullname" gorm:"column:fullname;not null" validate:"required,max=40"`
	NikNumber       *string           `json:"nik_number" gorm:"column:nik_number;uniqueIndex" validate:"required,unique"`
	KKNumber        *string           `json:"kk_number" gorm:"column:kk_number"`
	BornPlace       string            `json:"born_place" gorm:"column:born_place" validate:"required"`
	BirthDate       string            `json:"birth_date" gorm:"column:birth_date" validate:"required"`
	Gender          string            `json:"gender" gorm:"column:gender"`
	HomeAddress     *string           `json:"home_address" gorm:"column:home_address"`
	MaritalStatus   *string           `json:"marital_status" gorm:"column:marital_status"`
	ReligionId      uint              `json:"religion_id" gorm:"column:religion_id;not null" validate:"required"`
	Religion        Religion          `json:"-" gorm:"foreignKey:ReligionId;not null"`
	MemberStatusId  uint              `json:"member_status_id" gorm:"column:member_status_id;not null" validate:"required"`
	MemberStatus    MemberStatus      `json:"-" gorm:"foreignKey:MemberStatusId;not null"`
	UserId          *string           `json:"user_id" gorm:"column:user_id"`
	User            User              `json:"-" gorm:"foreignKey:UserId;not null"`
	Occupation      *string           `json:"occupation" gorm:"column:occupation"`
	Status          string            `json:"status" gorm:"column:status;default:resident"` //resident or guest
	RTProfileId     string            `json:"rt_profile_id" gorm:"column:rt_profile_id;not null"`
	RTProfile       RTProfile         `gorm:"foreignKey:RTProfileId;not null"`
	OutcomingLetter []OutcomingLetter `gorm:"foreignKey:MemberId"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`
}
