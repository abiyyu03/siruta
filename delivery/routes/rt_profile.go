package routes

import (
	"github.com/abiyyu03/siruta/delivery/http"
	"github.com/abiyyu03/siruta/delivery/middleware"
	"github.com/gofiber/fiber/v2"
)

func RegisterRTProfileRoutes(v1 fiber.Router, handler *http.RTProfileHttp) {
	v1.Get("/rt-profiles", middleware.JWTMiddleware([]int{1}), handler.GetData)
	v1.Get("/rt-profiles/:id", middleware.JWTMiddleware([]int{1, 2}), handler.GetDataById)
	v1.Get("/rt-profiles/:rw_profile_id/rw", middleware.JWTMiddleware([]int{1, 2}), handler.GetDataByRWProfileId)
}
