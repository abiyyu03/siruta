package routes

import (
	"github.com/abiyyu03/siruta/delivery/http"
	"github.com/abiyyu03/siruta/delivery/middleware"
	"github.com/gofiber/fiber/v2"
)

func RegisterReferalCodeRoutes(v1 fiber.Router, handler *http.ReferalCodeHttp) {
	v1.Get("/referal-codes", middleware.JWTMiddleware([]int{1}), handler.GetData)
	v1.Get("/referal-codes/:id", middleware.JWTMiddleware([]int{1}), handler.GetDataById)
	v1.Post("/referal-codes/validate", handler.ValidateReferalCode) //unauthenticated
	v1.Delete("/referal-codes/:id", middleware.JWTMiddleware([]int{1}), handler.GetDataById)
	v1.Get("/referal-codes/:profile_id/rt", middleware.JWTMiddleware([]int{1, 3}), handler.GetDataByRTProfileId)
	v1.Put("/referal-codes/:profile_id/rt/regenerate/:code", middleware.JWTMiddleware([]int{1, 3}), handler.RegenerateCode)
}
