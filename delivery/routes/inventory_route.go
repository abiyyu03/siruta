package routes

import (
	"github.com/abiyyu03/siruta/delivery/http"
	"github.com/abiyyu03/siruta/delivery/middleware"
	"github.com/gofiber/fiber/v2"
)

func RegisterInventoryRoutes(v1 fiber.Router, handler *http.InventoryHttp) {
	adminOnly := middleware.JWTMiddleware([]int{1})
	rtLeaderOnly := middleware.JWTMiddleware([]int{3})

	v1.Get("/inventories", adminOnly, rtLeaderOnly, handler.GetData)
	v1.Get("/inventories/:id", adminOnly, rtLeaderOnly, handler.GetDataById)
	v1.Get("/inventories/:rt_profile_id/rt", adminOnly, rtLeaderOnly, handler.GetDataByRTProfileId)
	v1.Post("/inventories", adminOnly, rtLeaderOnly, handler.StoreData)
	v1.Put("/inventories/:id", rtLeaderOnly, handler.UpdateData)
	v1.Delete("/inventories/:id", rtLeaderOnly, handler.DeleteData)
}
