package auth

import (
	"crypto/rsa"
	"log"
	"os"
	"time"

	"github.com/abiyyu03/siruta/entity"
	"github.com/abiyyu03/siruta/entity/request"
	"github.com/abiyyu03/siruta/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

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

func Login(ctx *fiber.Ctx) error {
	var request request.LoginRequest

	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	userRepository := new(repository.UserRepository)
	user, err := userRepository.FetchLogin(request.Username, request.Password)
	if err != nil {
		return entity.Error(ctx, fiber.StatusForbidden, "User not found")
	}

	token := jwt.New(jwt.SigningMethodRS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["email"] = user.Email
	claims["username"] = request.Username
	claims["role_id"] = user.RoleID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Sign the token with the private RSA key
	generatedToken, err := token.SignedString(privateKey)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return entity.Success(ctx, generatedToken, "login succcessfully")
}
