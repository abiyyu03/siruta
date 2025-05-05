package routes

import (
	"github.com/abiyyu03/siruta/delivery/http"
	"github.com/abiyyu03/siruta/delivery/middleware"
	"github.com/gofiber/fiber/v2"
)

func RegisterIncomingLetterRoutes(v1 fiber.Router, handler *http.IncomingLetterHttp) {
	adminOnly := middleware.JWTMiddleware([]int{1})
	rtLeaderOnly := middleware.JWTMiddleware([]int{3})

	v1.Get("/incoming-letters", adminOnly, handler.GetData)
	v1.Get("/incoming-letters/:id", adminOnly, handler.GetDataById)
	v1.Post("/incoming-letters", adminOnly, handler.StoreData)
	v1.Put("/incoming-letters/:id", adminOnly, handler.UpdateData)
	v1.Delete("/incoming-letters/:id", adminOnly, handler.DeleteData)
	v1.Get("/incoming-letters/:rt_profile_id/rt", rtLeaderOnly, handler.GetDataByRTProfileId)
}
