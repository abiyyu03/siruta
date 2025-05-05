package routes

import (
	"github.com/abiyyu03/siruta/delivery/http"
	"github.com/abiyyu03/siruta/delivery/middleware"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoleRoutes(v1 fiber.Router, handler *http.RoleHttp) {
	adminOnly := middleware.JWTMiddleware([]int{1})

	v1.Get("/roles", adminOnly, handler.GetData)
	v1.Get("/roles/:id", adminOnly, handler.GetDataById)
	v1.Post("/roles", adminOnly, handler.StoreData)
	v1.Put("/roles/:id", adminOnly, handler.UpdateData)
	v1.Delete("/roles/:id", adminOnly, handler.DeleteData)
}
