package routes

import (
	"github.com/abiyyu03/siruta/delivery/http"
	"github.com/abiyyu03/siruta/delivery/middleware"
	"github.com/gofiber/fiber/v2"
)

func RegisterInventoryRoutes(v1 fiber.Router, handler *http.InventoryHttp) {
	v1.Get("/inventories", middleware.JWTMiddleware([]int{1}), handler.GetData)
	v1.Get("/inventories/:id", middleware.JWTMiddleware([]int{1, 3}), handler.GetDataById)
	v1.Get("/inventories/:rt_profile_id/rt", middleware.JWTMiddleware([]int{1, 3}), handler.GetDataByRTProfileId)
	v1.Post("/inventories", middleware.JWTMiddleware([]int{1, 3}), handler.StoreData)
	v1.Put("/inventories/:id", middleware.JWTMiddleware([]int{1, 3}), handler.UpdateData)
	v1.Delete("/inventories/:id", middleware.JWTMiddleware([]int{1, 3}), handler.DeleteData)
}
