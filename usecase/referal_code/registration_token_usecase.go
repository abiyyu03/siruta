package referal_code

import (
	"encoding/base64"
	"strings"

	"github.com/abiyyu03/siruta/repository/register"
	"github.com/google/uuid"
)

type RegistrationTokenUsecase struct{}

var regTokenRepository *register.RegTokenRepository

func (r *RegistrationTokenUsecase) CreateToken(data string) (string, error) {
	uuid := uuid.New()
	parsedData := data + "***" + uuid.String()
	token := base64.URLEncoding.EncodeToString([]byte(parsedData))
	token = base64.URLEncoding.EncodeToString([]byte(token))

	return regTokenRepository.CreateToken(token)
}

func (r *RegistrationTokenUsecase) DecodeToken(token string) (string, error) {
	decodedFirst, err := base64.URLEncoding.DecodeString(token)

	if err != nil {
		return "", err
	}

	decodedSecond, err := base64.URLEncoding.DecodeString(string(decodedFirst))

	if err != nil {
		return "", err
	}

	data := strings.Split(string(decodedSecond), "***")

	return data[0], nil
}
