package routes

import (
	"github.com/abiyyu03/siruta/delivery/http"
	"github.com/abiyyu03/siruta/delivery/middleware"
	"github.com/gofiber/fiber/v2"
)

func RegisterOutcomingLetterRoutes(v1 fiber.Router, handler *http.OutcomingLetterHttp) {
	v1.Get("/outcoming-letters", middleware.JWTMiddleware([]int{1}), handler.GetData)
	v1.Get("/outcoming-letters/:id", middleware.JWTMiddleware([]int{1, 3}), handler.GetDataById)
	v1.Get("/outcoming-letters/:id/preview", middleware.JWTMiddleware([]int{1, 3, 4}), handler.GetPreview)
	v1.Get("/outcoming-letters/:rt_profile_id/rt", middleware.JWTMiddleware([]int{1, 3}), handler.GetDataByRTProfileId)
	v1.Delete("/outcoming-letters/:id", middleware.JWTMiddleware([]int{1, 3}), handler.DeleteData)
}
