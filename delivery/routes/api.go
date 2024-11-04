package routes

import (
	"github.com/abiyyu03/siruta/delivery/http"
	"github.com/abiyyu03/siruta/delivery/http/auth"
	"github.com/abiyyu03/siruta/delivery/http/register"
	"github.com/gofiber/fiber/v2"
)

func HttpRoutes(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")
	// v1Auth := api.Group("/v1")

	memberHttp := new(http.MemberHttp)

	// publicKeyBytes, err := os.ReadFile("./keys/public.pem")
	// if err != nil {
	// 	log.Fatal("Error reading public key")
	// }

	// pubKey, err := jwt.ParseRSAPublicKeyFromPEM(publicKeyBytes)
	// if err != nil {
	// 	log.Fatal("Error reading public key")
	// }

	// v1Auth.Use(middleware.ValidateToken(pubKey))

	v1.Post("/login", auth.Login)
	v1.Post("/register", register.Register)

	v1.Get("/members", memberHttp.GetData)
}
