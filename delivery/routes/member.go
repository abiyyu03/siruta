package routes

import (
	"github.com/abiyyu03/siruta/delivery/http"
	"github.com/abiyyu03/siruta/delivery/middleware"
	"github.com/gofiber/fiber/v2"
)

func RegisterMemberRoutes(v1 fiber.Router, handler *http.MemberHttp) {
	adminOnly := middleware.JWTMiddleware([]int{1})
	rtLeaderOnly := middleware.JWTMiddleware([]int{3})

	v1.Get("/members", adminOnly, handler.GetData)
	v1.Get("/members/:id", adminOnly, handler.GetDataById)
	v1.Put("/members/:id", adminOnly, handler.UpdateData)
	v1.Get("/members/:rt_profile_id/rt", adminOnly, rtLeaderOnly, handler.GetDataByRTProfileId)
}
