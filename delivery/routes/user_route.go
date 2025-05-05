package routes

import (
	"github.com/abiyyu03/siruta/delivery/http"
	"github.com/abiyyu03/siruta/delivery/middleware"
	"github.com/gofiber/fiber/v2"
)

func RegisterUserRoutes(v1 fiber.Router, handler *http.UserHttp) {
	adminOnly := middleware.JWTMiddleware([]int{1})

	v1.Get("/users", adminOnly, handler.GetData)
	v1.Get("/users/:id", adminOnly, handler.GetDataById)
	v1.Put("/users/revoke/:id", adminOnly, handler.RevokeUser)
}
