package routes

import (
	"github.com/abiyyu03/siruta/delivery/http"
	"github.com/abiyyu03/siruta/delivery/middleware"
	"github.com/gofiber/fiber/v2"
)

func RegisterRWLeaderRoutes(v1 fiber.Router, handler *http.RWLeaderHttp) {
	v1.Get("/rw-leaders", middleware.JWTMiddleware([]int{1}), handler.GetData)
	v1.Get("/rw-leaders/:id", middleware.JWTMiddleware([]int{1, 2}), handler.GetDataById)
	v1.Put("/rw-leaders/:id", middleware.JWTMiddleware([]int{1, 2}), handler.UpdateData)
}
