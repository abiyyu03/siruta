package routes

import (
	"github.com/abiyyu03/siruta/delivery/http"
	"github.com/abiyyu03/siruta/delivery/middleware"
	"github.com/gofiber/fiber/v2"
)

func RegisterRWProfileRoutes(v1 fiber.Router, handler *http.RWProfileHttp) {
	adminOnly := middleware.JWTMiddleware([]int{1})
	rtLeaderOnly := middleware.JWTMiddleware([]int{2})

	v1.Get("/rw-profiles", adminOnly, handler.GetData)
	v1.Get("/rw-profiles/:id", adminOnly, rtLeaderOnly, handler.GetDataById)
}
