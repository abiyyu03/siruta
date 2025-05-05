package routes

import (
	"github.com/abiyyu03/siruta/delivery/http"
	"github.com/abiyyu03/siruta/delivery/middleware"
	"github.com/gofiber/fiber/v2"
)

func RegisterOutcomingLetterRoutes(v1 fiber.Router, handler *http.OutcomingLetterHttp) {
	adminOnly := middleware.JWTMiddleware([]int{1})
	rtLeaderOnly := middleware.JWTMiddleware([]int{3})

	v1.Get("/outcoming-letters", adminOnly, handler.GetData)
	v1.Get("/outcoming-letters/:rt_profile_id", adminOnly, rtLeaderOnly, handler.GetDataByRTProfileId)
	v1.Get("/outcoming-letters/:id", adminOnly, rtLeaderOnly, handler.GetDataById)
	v1.Get("/outcoming-letters/:id/preview", handler.GetPreview)
	v1.Get("/outcoming-letters/:rt_profile_id/rt", adminOnly, rtLeaderOnly, handler.GetDataByRTProfileId)
	v1.Delete("/outcoming-letters/:id", adminOnly, rtLeaderOnly, handler.DeleteData)
}
