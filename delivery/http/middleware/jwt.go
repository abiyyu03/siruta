package middleware

import (
	"crypto/rsa"
	"log"
	"os"
	"reflect"

	"github.com/abiyyu03/siruta/entity"
	"github.com/abiyyu03/siruta/repository"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
)

// var privateKey *rsa.PrivateKey

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
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return entity.Error(c, fiber.StatusUnauthorized, err.Error())
		},
		SuccessHandler: func(ctx *fiber.Ctx) error {
			// Extract user claims from the token
			user := ctx.Locals("user").(*jwt.Token)
			claims := user.Claims.(jwt.MapClaims)

			if claims["role_id"] != 0 {
				log.Printf("signed role id is : %v", reflect.TypeOf((claims["role_id"])))
			}

			// Extract the role_id from the JWT claims
			roleID := claims["role_id"].(float64)

			castedRoleID := int(roleID)

			allowedRoleId, err := new(repository.RoleRepository).FetchById(castedRoleID)

			if err != nil {
				return err
			}

			if !HasRequiredRole(allowedRoleId.ID, allowedRoles) {
				return entity.Error(ctx, fiber.StatusForbidden, "Insufficient role")
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
