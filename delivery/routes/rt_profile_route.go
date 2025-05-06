package routes

import (
	"github.com/abiyyu03/siruta/delivery/http"
	"github.com/abiyyu03/siruta/delivery/middleware"
	"github.com/gofiber/fiber/v2"
)

func RegisterRTProfileRoutes(v1 fiber.Router, handler *http.RTProfileHttp) {
	adminOnly := middleware.JWTMiddleware([]int{1})
	rwLeaderOnly := middleware.JWTMiddleware([]int{2})

	v1.Get("/rt-profiles", adminOnly, handler.GetData)
	v1.Get("/rt-profiles/:id", adminOnly, rwLeaderOnly, handler.GetDataById)
	v1.Get("/rt-profiles/:rw_profile_id/rw", adminOnly, rwLeaderOnly, handler.GetDataByRWProfileId)
}
