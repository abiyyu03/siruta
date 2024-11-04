package request

import "github.com/go-playground/validator/v10"

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required,min=8"`
}

// Fungsi untuk validasi RegisterRequest
func (r *LoginRequest) Validate() error {
	var validate = validator.New()

	return validate.Struct(r)
}
