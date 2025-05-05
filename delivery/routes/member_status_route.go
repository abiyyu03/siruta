package routes

import (
	"github.com/abiyyu03/siruta/delivery/http"
	"github.com/abiyyu03/siruta/delivery/middleware"
	"github.com/gofiber/fiber/v2"
)

func RegisterMemberStatusRoutes(v1 fiber.Router, handler *http.MemberStatusHttp) {
	adminOnly := middleware.JWTMiddleware([]int{1})

	v1.Get("/member-status", adminOnly, handler.GetData)
	v1.Get("/member-status/:id", adminOnly, handler.GetDataById)
	v1.Post("/member-status", adminOnly, handler.StoreData)
	v1.Put("/member-status/:id", adminOnly, handler.UpdateData)
	v1.Delete("/member-status/:id", adminOnly, handler.DeleteData)
}
