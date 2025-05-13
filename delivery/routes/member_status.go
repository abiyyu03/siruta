package routes

import (
	"github.com/abiyyu03/siruta/delivery/http"
	"github.com/abiyyu03/siruta/delivery/middleware"
	"github.com/gofiber/fiber/v2"
)

func RegisterMemberStatusRoutes(v1 fiber.Router, handler *http.MemberStatusHttp) {
	v1.Get("/member-status", handler.GetData) // unauthenticated
	v1.Get("/member-status/:id", middleware.JWTMiddleware([]int{1}), handler.GetDataById)
	v1.Post("/member-status", middleware.JWTMiddleware([]int{1}), handler.StoreData)
	v1.Put("/member-status/:id", middleware.JWTMiddleware([]int{1}), handler.UpdateData)
	v1.Delete("/member-status/:id", middleware.JWTMiddleware([]int{1}), handler.DeleteData)
}
