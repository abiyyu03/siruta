package routes

import (
	"github.com/abiyyu03/siruta/delivery/http"
	"github.com/abiyyu03/siruta/delivery/middleware"
	"github.com/gofiber/fiber/v2"
)

func RegisterRWLeaderRoutes(v1 fiber.Router, handler *http.RWLeaderHttp) {
	adminOnly := middleware.JWTMiddleware([]int{1})
	rwLeaderOnly := middleware.JWTMiddleware([]int{2})

	v1.Get("/rw-leaders", adminOnly, handler.GetData)
	v1.Get("/rw-leaders/:id", adminOnly, rwLeaderOnly, handler.GetDataById)
	v1.Put("/rw-leaders/:id", adminOnly, rwLeaderOnly, handler.UpdateData)
}
