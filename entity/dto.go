package entity

import (
	"time"

	"github.com/abiyyu03/siruta/entity/model"
	"gorm.io/gorm"
)

type CheckResidentMember struct {
	BirthDate    string `json:"birth_date" validate:"required"`
	NikNumber    string `json:"nik_number" validate:"required"`
	MemberStatus string `json:"member_status" validate:"required"`
}

// Register

type LeaderRegisterRequest struct {
	Email       string `json:"email" validate:"required"`
	Password    string `json:"password" validate:"required"`
	Fullname    string `json:"fullname" validate:"required"`
	NikNumber   string `json:"nik_number" validate:"required"`
	KKNumber    string `json:"kk_number" validate:"required"`
	FullAddress string `json:"full_address" validate:"required"`
}

type RTProfileRegisterRequest struct {
	RTNumber    string  `json:"rt_number" validate:"required"`
	VillageID   int     `json:"village_id" validate:"required"`
	RTEmail     string  `json:"rt_email" validate:"required" gorm:"uniqueIndex"`
	MobilePhone string  `json:"mobile_phone" validate:"required"`
	FullAddress string  `json:"full_address" validate:"required"`
	Description string  `json:"description"`
	Longitude   float64 `json:"longitude"`
	Latitude    float64 `json:"latitude"`
	ReferalCode string  `json:"referal_code"`
}

type MemberRegisterRequest struct {
	Email          string `json:"email" validate:"required"`
	Password       string `json:"password" validate:"required"`
	Fullname       string `json:"fullname" validate:"required"`
	NikNumber      string `json:"nik_number" validate:"required"`
	KKNumber       string `json:"kk_number" validate:"required"`
	BornPlace      string `json:"born_place" validate:"required"`
	BirthDate      string `json:"birth_date" validate:"required"`
	Gender         string `json:"gender" validate:"required"`
	HomeAddress    string `json:"home_address" validate:"required"`
	MaritalStatus  string `json:"marital_status" validate:"required"`
	ReligionId     uint   `json:"religion_id" validate:"required"`
	MemberStatusId uint   `json:"member_status_id" validate:"required"`
	Occupation     string `json:"occupation" validate:"required"`
	Status         string `json:"status" validate:"required"`
}

// Auth

type UpdateProfilePhoto struct {
	Photo string `json:"photo"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type ResetPassword struct {
	Password        string `json:"password" validate:"required"`
	ConfirmPassword string `json:"confirm_password" validate:"required"`
}

type LoginRepositoryResponse struct {
	FullName string `json:"fullname"`
	Email    string `json:"email"`
	RoleName string `json:"role_name"`
	RoleID   int    `json:"role_id"`
}

type AuthResponse struct {
	Data        LoginRepositoryResponse
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
