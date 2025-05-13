package routes

import (
	"github.com/abiyyu03/siruta/delivery/http"
	"github.com/abiyyu03/siruta/delivery/middleware"
	"github.com/gofiber/fiber/v2"
)

func RegisterRTLeaderRoutes(v1 fiber.Router, handler *http.RTLeaderHttp) {
	v1.Get("/rt-leaders", middleware.JWTMiddleware([]int{1}), handler.GetData)
	v1.Get("/rt-leaders/:id", middleware.JWTMiddleware([]int{1, 3}), handler.GetDataById)
	v1.Put("/rt-leaders/:id", middleware.JWTMiddleware([]int{1, 3}), handler.UpdateData)
}
