package middleware

import (
	"github.com/abiyyu03/siruta/repository/register"
)

var RegTokenRepository = new(register.RegTokenRepository)

func TokenRegisterValidator(token string) bool {
	if token == "" {
		return false
	}

	fetchedToken, err := RegTokenRepository.Validate(token)
	if err != nil {
		return false
	}

	if fetchedToken != token {
		return false
	}

	return true
}
