package usecase

import (
	"crypto/rsa"
	"log"
	"os"
	"time"

	"github.com/abiyyu03/siruta/entity"
	"github.com/abiyyu03/siruta/entity/constant"
	"github.com/abiyyu03/siruta/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type AuthUsecase struct{}

var authRepository = new(repository.AuthRepository)

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
	claims["email"] = email
	claims["role_id"] = user.RoleID
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix()

	// Sign the token with the private RSA key
	generatedToken, err := token.SignedString(privateKey)

	if err != nil {
		return entity.Error(ctx, fiber.ErrForbidden.Code, constant.Errors["Unauthorized"].Message, constant.Errors["Unauthorized"].Clue)
	}

	finalResponse := &entity.AuthResponse{
		FullName:    member.Fullname,
		Email:       user.Email,
		RoleName:    user.Role.Name,
		AccessToken: generatedToken,
	}

	return entity.Success(ctx, finalResponse, "Login Berhasil")

}

// func (l *AuthUsecase) RevokeToken(ctx *fiber.Ctx) error {

// }
