package auth

import (
	"crypto/rsa"
	"log"
	"os"
	"time"

	"github.com/abiyyu03/siruta/entity"
	"github.com/abiyyu03/siruta/entity/constant"
	"github.com/abiyyu03/siruta/repository/auth"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type AuthUsecaseInterface interface {
	IssueAuthToken(ctx *fiber.Ctx, email string, password string) error
}

type AuthUsecase struct{}

var authRepository *auth.AuthRepository

var privateKey *rsa.PrivateKey

func init() {
	var err error
	keyData, err := os.ReadFile("./keys/private.pem")
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(keyData)
	if err != nil {
		log.Fatal("Error parsing private key:", err)
	}
}

func (l *AuthUsecase) IssueAuthToken(ctx *fiber.Ctx, email string, password string) error {
	var fullname string
	user, member, err := authRepository.FetchLogin(email, password)

	if err != nil {
		return entity.Error(ctx,
			fiber.StatusUnauthorized,
			constant.Errors["AccountInputError"].Message,
			constant.Errors["AccountInputError"].Clue,
		)
	}

	token := jwt.New(jwt.SigningMethodRS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["email"] = user.Email
	claims["role_id"] = user.RoleID
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix()

	// Sign the token with the private RSA key
	generatedToken, err := token.SignedString(privateKey)

	if err != nil {
		return entity.Error(ctx, fiber.ErrForbidden.Code, constant.Errors["Unauthorized"].Message, constant.Errors["Unauthorized"].Clue)
	}

	if user.Role.Name == "Super Admin" {
		fullname = "Super Administrator"
	} else {
		fullname = member.Fullname
	}

	finalResponse := &entity.AuthResponse{
		FullName:    fullname,
		Email:       user.Email,
		RoleName:    user.Role.Name,
		AccessToken: generatedToken,
	}

	return entity.Success(ctx, finalResponse, "Login Berhasil")

}
