package entity

import (
	"time"

	"github.com/abiyyu03/siruta/entity/model"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type AuthResponse struct {
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
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// ------------------------------------------------------------------------------
// Letter Preview Response
type OutcomeLetterResponse struct {
	ID           string            `json:"id"`
	LetterNumber int               `json:"letter_number"`
	Date         string            `json:"date"`
	Member       MemberPreview     `json:"member"`
	LetterType   LetterTypePreview `json:"letter_type"`
	RTProfile    RTProfilePreview  `json:"rt_profile"`
	IsRTApproved bool              `json:"is_rt_approved"`
	Description  string            `json:"description"`
}

type MemberPreview struct {
	ID              string  `json:"id"`
	Fullname        string  `json:"fullname"`
	NIKNumber       *string `json:"nik_number"`
	KKNumber        *string `json:"kk_number"`
	BornPlace       string  `json:"born_place"`
	BirthDate       string  `json:"birth_date"`
	Gender          string  `json:"gender"`
	HomeAddress     *string `json:"home_address"`
	MaritalStatus   *string `json:"marital_status"`
	ReligionID      uint    `json:"religion_id"`
	MemberStatusID  uint    `json:"member_status_id"`
	Occupation      *string `json:"occupation"`
	Status          string  `json:"status"`
	OutcomingLetter *any    `json:"OutcomingLetter"`
}

type LetterTypePreview struct {
	ID       int    `json:"id"`
	TypeName string `json:"type_name"`
	Code     string `json:"code"`
}

type RTProfilePreview struct {
	ID       string  `json:"id"`
	RTNumber string  `json:"rt_number"`
	RTLogo   *string `json:"rt_logo"` // nullable
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

func Error(c *fiber.Ctx, statusCode int, message string, desc string) error {
	response := Response{
		Status:      "Gagal",
		Code:        statusCode,
		Description: desc,
		Message:     message,
		// Data:    nil,
	}

	return c.Status(statusCode).JSON(response)
}
