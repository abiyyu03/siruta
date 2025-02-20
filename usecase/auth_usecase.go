package usecase

import (
	"crypto/rsa"
	"log"
	"os"
	"time"

	"github.com/abiyyu03/siruta/entity"
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

func (l *AuthUsecase) IssueAuthToken(ctx *fiber.Ctx, username string, password string) (*entity.AuthResponse, error) {
	user, member, err := authRepository.FetchLogin(username, password)
	if err != nil {
		return nil, entity.Error(ctx, fiber.StatusForbidden, "Username or password invalid")
	}

	token := jwt.New(jwt.SigningMethodRS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["email"] = user.Email
	claims["username"] = username
	claims["role_id"] = user.RoleID
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix()

	// Sign the token with the private RSA key
	generatedToken, err := token.SignedString(privateKey)

	if err != nil {
		return nil, entity.Error(ctx, fiber.ErrUnauthorized.Code, "Login failed")
	}

	return &entity.AuthResponse{
		Username:    user.Username,
		FullName:    member.Fullname,
		Email:       user.Email,
		RoleName:    user.Role.Name,
		AccessToken: generatedToken,
	}, nil

}

// func (l *AuthUsecase) RevokeToken(ctx *fiber.Ctx) error {

// }
