package request

type LeaderRegisterRequest struct {
	Email       string `json:"email" validate:"required"`
	Password    string `json:"password" validate:"required"`
	Fullname    string `json:"fullname" validate:"required"`
	NikNumber   string `json:"nik_number" validate:"required"`
	KKNumber    string `json:"kk_number" validate:"required"`
	FullAddress string `json:"full_address" validate:"required"`
}
