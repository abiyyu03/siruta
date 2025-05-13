package routes

import (
	"github.com/abiyyu03/siruta/delivery/http"
	"github.com/abiyyu03/siruta/delivery/middleware"
	"github.com/gofiber/fiber/v2"
)

func RegisterLetterTypeRoutes(v1 fiber.Router, handler *http.LetterTypeHttp) {
	v1.Get("/letter-types", middleware.JWTMiddleware([]int{1, 3, 4}), handler.GetData)
	v1.Get("/letter-types/:id", middleware.JWTMiddleware([]int{1}), handler.GetDataById)
	v1.Post("/letter-types", middleware.JWTMiddleware([]int{1}), handler.StoreData)
	v1.Put("/letter-types/:id", middleware.JWTMiddleware([]int{1}), handler.UpdateData)
	v1.Delete("/letter-types/:id", middleware.JWTMiddleware([]int{1}), handler.DeleteData)
}
