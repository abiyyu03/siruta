package routes

import (
	"github.com/abiyyu03/siruta/delivery/http"
	"github.com/abiyyu03/siruta/delivery/middleware"
	"github.com/gofiber/fiber/v2"
)

func RegisterReferalCodeRoutes(v1 fiber.Router, handler *http.ReferalCodeHttp) {
	adminOnly := middleware.JWTMiddleware([]int{1})
	rtLeaderOnly := middleware.JWTMiddleware([]int{3})

	v1.Get("/referal-codes", adminOnly, handler.GetData)
	v1.Get("/referal-codes/:id", adminOnly, handler.GetDataById)
	v1.Get("/referal-codes/:rt_profile_id", adminOnly, handler.GetDataByRTProfileId)
	v1.Post("/referal-codes/validate", handler.ValidateReferalCode) //unauthenticated
	v1.Delete("/referal-codes/:id", adminOnly, handler.GetDataById)
	v1.Get("/referal-codes/:profile_id/rt", rtLeaderOnly, handler.GetDataByRTProfileId)
	v1.Put("/referal-codes/:profile_id/rt/regenerate/:code", rtLeaderOnly, handler.RegenerateCode)
}
