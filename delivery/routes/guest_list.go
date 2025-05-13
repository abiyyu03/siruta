package routes

import (
	"github.com/abiyyu03/siruta/delivery/http"
	"github.com/abiyyu03/siruta/delivery/middleware"
	"github.com/gofiber/fiber/v2"
)

func RegisterGuestListRoutes(v1 fiber.Router, handler *http.GuestListHttp) {
	v1.Get("/guest-lists", middleware.JWTMiddleware([]int{1}), handler.GetData)
	v1.Get("/guest-lists/:rt_profile_id/rt", middleware.JWTMiddleware([]int{1, 3}), handler.GetDataByRTProfileId)
	v1.Get("/guest-lists/:id", middleware.JWTMiddleware([]int{1, 3}), handler.GetDataById)
	v1.Put("/guest-lists/:id", middleware.JWTMiddleware([]int{1, 3}), handler.UpdateData)
	v1.Delete("/guest-lists/:id", middleware.JWTMiddleware([]int{1, 3}), handler.DeleteData)
}
