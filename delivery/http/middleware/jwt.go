package middleware

import (
	"crypto/rsa"
	"log"
	"os"

	"github.com/abiyyu03/siruta/entity"
	"github.com/abiyyu03/siruta/entity/constant"
	"github.com/abiyyu03/siruta/repository/role"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
)

// privateKey *rsa.PrivateKey

func loadPublicKey() (*rsa.PublicKey, error) {
	// publicKeyPath := os.Getenv("JWT_PUBLIC_KEY_PATH")
	publicKeyData, err := os.ReadFile("./keys/public.pem")
	if err != nil {
		return nil, err
	}

	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicKeyData)
	if err != nil {
		return nil, err
	}

	return publicKey, nil
}

// JWTMiddleware returns the configured Fiber JWT middleware
func JWTMiddleware(allowedRoles []int) fiber.Handler {
	publicKey, err := loadPublicKey()
	if err != nil {
		log.Fatalf("Failed to load public key: %v", err)
	}
	log.Println("JWT Middleware initialized with public key")
	return jwtware.New(jwtware.Config{
		SigningMethod: "RS256",
		SigningKey:    publicKey,
		ContextKey:    "user",
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			return entity.Error(ctx, fiber.StatusUnauthorized, constant.Errors["Unauthorized"].Message, constant.Errors["Unauthorized"].Clue)
		},
		SuccessHandler: func(ctx *fiber.Ctx) error {
			// Extract user claims from the token
			user := ctx.Locals("user").(*jwt.Token)
			claims := user.Claims.(jwt.MapClaims)

			if claims["role_id"] != 0 {
				log.Printf("role signed")
			}

			// Extract the role_id from the JWT claims
			roleID := claims["role_id"].(float64)

			castedRoleID := int(roleID)

			allowedRoleId, err := new(role.RoleRepository).FetchById(castedRoleID)

			if err != nil {
				return err
			}

			if !HasRequiredRole(allowedRoleId.ID, allowedRoles) {
				return entity.Error(ctx, fiber.StatusUnauthorized, constant.Errors["InvalidRole"].Message, constant.Errors["InvalidRole"].Clue)
			}

			return ctx.Next()
		},
	})
}

func HasRequiredRole(userRoles int, allowedRoles []int) bool {
	for _, allowed := range allowedRoles {
		if userRoles == allowed {
			return true
		}
	}
	return false
}
