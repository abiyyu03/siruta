package routes

import (
	"github.com/abiyyu03/siruta/delivery/http"
	"github.com/abiyyu03/siruta/delivery/middleware"
	"github.com/gofiber/fiber/v2"
)

func RegisterRWProfileRoutes(v1 fiber.Router, handler *http.RWProfileHttp) {
	v1.Get("/rw-profiles", middleware.JWTMiddleware([]int{1}), handler.GetData)
	v1.Get("/rw-profiles/:id", middleware.JWTMiddleware([]int{1, 3}), handler.GetDataById)
}
