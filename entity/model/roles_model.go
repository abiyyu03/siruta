package model

import (
	"time"

	"gorm.io/gorm"
)

type Role struct {
	*gorm.Model
	ID        int    `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	Name      string `json:"name" gorm:"column:name;not null"`
	Users     []User `json:"-" gorm:"foreignKey:RoleID"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
