package routes

import (
	"github.com/abiyyu03/siruta/delivery/http"
	"github.com/abiyyu03/siruta/delivery/middleware"
	"github.com/gofiber/fiber/v2"
)

func RegisterMemberRoutes(v1 fiber.Router, handler *http.MemberHttp) {
	v1.Get("/members", middleware.JWTMiddleware([]int{1}), handler.GetData)
	v1.Get("/members/:id", middleware.JWTMiddleware([]int{1}), handler.GetDataById)
	v1.Put("/members/:id", middleware.JWTMiddleware([]int{1}), handler.UpdateData)
	v1.Get("/members/:rt_profile_id/rt", middleware.JWTMiddleware([]int{1, 3}), handler.GetDataByRTProfileId)
}
