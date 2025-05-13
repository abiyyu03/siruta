package routes

import (
	"github.com/abiyyu03/siruta/delivery/http"
	"github.com/abiyyu03/siruta/delivery/middleware"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoleRoutes(v1 fiber.Router, handler *http.RoleHttp) {
	v1.Get("/roles", middleware.JWTMiddleware([]int{1}), handler.GetData)
	v1.Get("/roles/:id", middleware.JWTMiddleware([]int{1}), handler.GetDataById)
	v1.Post("/roles", middleware.JWTMiddleware([]int{1}), handler.StoreData)
	v1.Put("/roles/:id", middleware.JWTMiddleware([]int{1}), handler.UpdateData)
	v1.Delete("/roles/:id", middleware.JWTMiddleware([]int{1}), handler.DeleteData)
}
