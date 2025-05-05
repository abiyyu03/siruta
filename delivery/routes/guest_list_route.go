package routes

import (
	"github.com/abiyyu03/siruta/delivery/http"
	"github.com/abiyyu03/siruta/delivery/middleware"
	"github.com/gofiber/fiber/v2"
)

func RegisterGuestListRoutes(v1 fiber.Router, handler *http.GuestListHttp) {
	adminOnly := middleware.JWTMiddleware([]int{1})
	rtLeaderOnly := middleware.JWTMiddleware([]int{3})

	v1.Get("/guest-lists", adminOnly, handler.GetData)
	v1.Get("/guest-lists/:rt_profile_id/rt", adminOnly, rtLeaderOnly, handler.GetDataByRTProfileId)
	v1.Get("/guest-lists/:id", adminOnly, rtLeaderOnly, handler.GetDataById)
	v1.Put("/guest-lists/:id", adminOnly, rtLeaderOnly, handler.UpdateData)
	v1.Delete("/guest-lists/:id", adminOnly, rtLeaderOnly, handler.DeleteData)
}
