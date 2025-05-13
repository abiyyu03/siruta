package routes

import (
	"github.com/abiyyu03/siruta/delivery/http"
	"github.com/abiyyu03/siruta/delivery/middleware"
	"github.com/gofiber/fiber/v2"
)

func RegisterVillageRoutes(v1 fiber.Router, handler *http.VillageHttp) {
	v1.Get("/villages", middleware.JWTMiddleware([]int{1, 3, 4}), handler.GetData)
	v1.Get("/villages/:id", middleware.JWTMiddleware([]int{1}), handler.GetDataById)
	v1.Post("/villages", middleware.JWTMiddleware([]int{1}), handler.StoreData)
	v1.Put("/villages/:id", middleware.JWTMiddleware([]int{1}), handler.UpdateData)
	v1.Delete("/villages/:id", middleware.JWTMiddleware([]int{1}), handler.DeleteData)
}
