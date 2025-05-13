package routes

import (
	"github.com/abiyyu03/siruta/delivery/http"
	"github.com/abiyyu03/siruta/delivery/middleware"
	"github.com/gofiber/fiber/v2"
)

func RegisterUserRoutes(v1 fiber.Router, handler *http.UserHttp) {
	v1.Get("/users", middleware.JWTMiddleware([]int{1}), handler.GetData)
	v1.Get("/users/:id", middleware.JWTMiddleware([]int{1}), handler.GetDataById)
	v1.Put("/users/photo/:id", middleware.JWTMiddleware([]int{1, 2, 3, 4}), handler.UpdateProfilePhoto)
	v1.Put("/users/revoke/:id", middleware.JWTMiddleware([]int{1}), handler.RevokeUser)
}
