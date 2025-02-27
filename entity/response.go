package entity

import (
	"time"

	"github.com/abiyyu03/siruta/entity/model"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type AuthResponse struct {
	Username    string `json:"username"`
	FullName    string `json:"full_name"`
	Email       string `json:"email"`
	RoleName    string `json:"role_name"`
	AccessToken string `json:"accessToken"`
}

type UserResponse struct {
	ID        string     `json:"id"`
	RoleID    uint       `json:"role_id"`
	Role      model.Role `json:"role"`
	Email     string     `json:"email"`
	Username  string     `json:"username"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

//------------------------------------------------------------------------------
// Standard Response

type Response struct {
	Status      string      `json:"status"`
	Code        int         `json:"code"`
	Message     string      `json:"message"`
	Description string      `json:"description,omitempty"`
	Data        interface{} `json:"data,omitempty"`
}

func Success(c *fiber.Ctx, data interface{}, message string) error {
	response := Response{
		Status:  "Berhasil",
		Code:    fiber.StatusOK,
		Message: message,
		Data:    data,
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

func Error(c *fiber.Ctx, statusCode int, message string) error {
	response := Response{
		Status:      "Gagal",
		Code:        statusCode,
		Description: "-",
		Message:     message,
		// Data:    nil,
	}

	return c.Status(statusCode).JSON(response)
}
