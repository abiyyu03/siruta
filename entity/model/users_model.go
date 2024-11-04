package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        string `json:"id" gorm:"primaryKey"`
	FirstName string `json:"first_name" validate:"required, max=32"`
	LastName  string `json:"last_name" validate:"required, max=32"`
	RoleID    uint   `json:"role_id"`           // Karena one-to-many, gunakan uint sebagai foreign key
	Role      Role   `gorm:"foreignKey:RoleID"` // Menghubungkan dengan model Role
	Email     string `json:"email" gorm:"uniqueIndex" validate:"required"`
	Username  string `json:"username" gorm:"uniqueIndex" validate:"required"`
	Password  string `json:"-" validate:"required, min=8"` // Jangan tampilkan password dalam JSON
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
