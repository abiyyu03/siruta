package request

type RegisterRequest struct {
	Email          string `json:"email" validate:"required"`
	Password       string `json:"password" validate:"required"`
	Fullname       string `json:"fullname" validate:"required"`
	NikNumber      string `json:"nik_number" validate:"required"`
	KKNumber       string `json:"kk_number" validate:"required"`
	BornPlace      string `json:"born_place" validate:"required"`
	BirthDate      string `json:"birth_place" validate:"required"`
	Gender         string `json:"gender" validate:"required"`
	HomeAddress    string `json:"home_address" validate:"required"`
	MaritalStatus  string `json:"marital_status" validate:"required"`
	ReligionId     uint   `json:"religion_id" validate:"required"`
	MemberStatusId uint   `json:"member_status_id" validate:"required"`
	Occupation     string `json:"occupation" validate:"required"`
	Status         string `json:"role_id" validate:"required"`
}
