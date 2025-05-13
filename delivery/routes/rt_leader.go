package routes

import (
	"github.com/abiyyu03/siruta/delivery/http"
	"github.com/abiyyu03/siruta/delivery/middleware"
	"github.com/gofiber/fiber/v2"
)

func RegisterRTLeaderRoutes(v1 fiber.Router, handler *http.RTLeaderHttp) {
	adminOnly := middleware.JWTMiddleware([]int{1})
	rtLeaderOnly := middleware.JWTMiddleware([]int{3})

	v1.Get("/rt-leaders", adminOnly, handler.GetData)
	v1.Get("/rt-leaders/:id", adminOnly, rtLeaderOnly, handler.GetDataById)
	v1.Put("/rt-leaders/:id", adminOnly, rtLeaderOnly, handler.UpdateData)
}
