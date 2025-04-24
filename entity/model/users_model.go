package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	*gorm.Model
	ID           string `json:"id" gorm:"column:id;primaryKey;not null"`
	RoleID       uint   `json:"role_id" gorm:"column:role_id;not null"`
	Role         Role   `json:"-" gorm:"foreignKey:RoleID;references:ID"`
	Email        string `json:"email" gorm:"column:email;uniqueIndex;not null" validate:"required"`
	Password     string `json:"password,omitempty" gorm:"column:password;not null" validate:"required,min=8"`
	IsAuthorized bool   `json:"is_authorized" gorm:"column:is_authorized;default:true;not null"`
	// Member       Member `json:"-" gorm:"foreignKey:UserId,omitempty"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
