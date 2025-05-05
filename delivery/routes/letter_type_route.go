package routes

import (
	"github.com/abiyyu03/siruta/delivery/http"
	"github.com/abiyyu03/siruta/delivery/middleware"
	"github.com/gofiber/fiber/v2"
)

func RegisterLetterTypeRoutes(v1 fiber.Router, handler *http.LetterTypeHttp) {
	adminOnly := middleware.JWTMiddleware([]int{1})
	memberOnly := middleware.JWTMiddleware([]int{4})
	rtLeaderOnly := middleware.JWTMiddleware([]int{4})

	v1.Get("/letter-types", adminOnly, rtLeaderOnly, memberOnly, handler.GetData)
	v1.Get("/letter-types/:id", adminOnly, handler.GetDataById)
	v1.Post("/letter-types", adminOnly, handler.StoreData)
	v1.Put("/letter-types/:id", adminOnly, handler.UpdateData)
	v1.Delete("/letter-types/:id", adminOnly, handler.DeleteData)
}
