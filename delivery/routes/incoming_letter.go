package routes

import (
	"github.com/abiyyu03/siruta/delivery/http"
	"github.com/abiyyu03/siruta/delivery/middleware"
	"github.com/gofiber/fiber/v2"
)

func RegisterIncomingLetterRoutes(v1 fiber.Router, handler *http.IncomingLetterHttp) {
	v1.Get("/incoming-letters", middleware.JWTMiddleware([]int{1}), handler.GetData)
	v1.Get("/incoming-letters/:id", middleware.JWTMiddleware([]int{1}), handler.GetDataById)
	v1.Post("/incoming-letters", middleware.JWTMiddleware([]int{1}), handler.StoreData)
	v1.Put("/incoming-letters/:id", middleware.JWTMiddleware([]int{1}), handler.UpdateData)
	v1.Delete("/incoming-letters/:id", middleware.JWTMiddleware([]int{1}), handler.DeleteData)
	v1.Get("/incoming-letters/:rt_profile_id/rt", middleware.JWTMiddleware([]int{1, 3}), handler.GetDataByRTProfileId)
}
