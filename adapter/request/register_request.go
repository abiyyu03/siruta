package request

import (
	"github.com/go-playground/validator/v10"
)

type RegisterRequest struct {
	Username  string `json:"username" validate:"required"`
	Password  string `json:"password" validate:"required,min=8"`
	FirstName string `json:"first_name" validate:"required, max=32"`
	LastName  string `json:"last_name" validate:"required, max=32"`
	RoleID    uint   `json:"role_id" validate:"required"`
	Email     string `json:"email" validate:"required, unique"`
}

// Fungsi untuk validasi RegisterRequest
func (r *RegisterRequest) Validate() error {
	var validate = validator.New()

	return validate.Struct(r)
}
