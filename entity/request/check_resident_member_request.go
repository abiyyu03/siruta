package request

type CheckResidentMember struct {
	BirthDate    string `json:"birth_date" validate:"required"`
	NikNumber    string `json:"nik_number" validate:"required"`
	MemberStatus string `json:"member_status" validate:"required"`
}
