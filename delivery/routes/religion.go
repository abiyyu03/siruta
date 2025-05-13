package routes

import (
	"github.com/abiyyu03/siruta/delivery/http"
	"github.com/abiyyu03/siruta/delivery/middleware"
	"github.com/gofiber/fiber/v2"
)

func RegisterReligionRoutes(v1 fiber.Router, handler *http.ReligionHttp) {
	v1.Get("/religions", handler.GetData)
	v1.Get("/religions/:id", middleware.JWTMiddleware([]int{1, 3, 4}), handler.GetDataById)
	v1.Post("/religions", middleware.JWTMiddleware([]int{1}), handler.StoreData)
	v1.Put("/religions/:id", middleware.JWTMiddleware([]int{1}), handler.UpdateData)
	v1.Delete("/religions/:id", middleware.JWTMiddleware([]int{1}), handler.DeleteData)
}
