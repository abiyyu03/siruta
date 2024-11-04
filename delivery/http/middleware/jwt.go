package middleware

import (
	"crypto/rsa"
	"fmt"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// func ValidateJWT(ctx *fiber.Ctx) error {
// 	publicKey := usecase.GoDotEnv("PUBLIC_KEY")
// 	token, err := getToken(publicKey)

// 	if err != nil {
// 		return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{
// 			"error": err,
// 		})
// 	}

// 	_, ok = token.Claims(jwt.)
// 	if
// }

func ValidateToken(publicKey *rsa.PublicKey) fiber.Handler { //publicKey *rsa.PublicKey
	return func(ctx *fiber.Ctx) error {
		tokenString, _ := getTokenFromHeader(ctx)
		if tokenString == "" {
			return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"error": "Token not found",
			})
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, fmt.Errorf("invalid token")
			}
			return publicKey, nil
		})

		if err != nil || !token.Valid {
			return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return ctx.Next()
	}
}

func getTokenFromHeader(ctx *fiber.Ctx) (string, error) {
	bearerToken := ctx.Get("Authorization")

	return strings.Replace(bearerToken, "Bearer ", "", -1), nil
}
