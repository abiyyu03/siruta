package routes

import (
	"github.com/abiyyu03/siruta/delivery/http"
	"github.com/abiyyu03/siruta/delivery/middleware"
	"github.com/gofiber/fiber/v2"
)

func RegisterVillageRoutes(v1 fiber.Router, handler *http.VillageHttp) {
	adminOnly := middleware.JWTMiddleware([]int{1})
	rtLeaderOnly := middleware.JWTMiddleware([]int{3})
	memberOnly := middleware.JWTMiddleware([]int{4})

	v1.Get("/villages", adminOnly, rtLeaderOnly, memberOnly, handler.GetData)
	v1.Get("/villages/:id", adminOnly, handler.GetDataById)
	v1.Post("/villages", adminOnly, handler.StoreData)
	v1.Put("/villages/:id", adminOnly, handler.UpdateData)
	v1.Delete("/villages/:id", adminOnly, handler.DeleteData)
}
