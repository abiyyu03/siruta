package entity

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
