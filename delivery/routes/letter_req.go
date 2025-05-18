package routes

import (
	"github.com/abiyyu03/siruta/delivery/http"
	"github.com/abiyyu03/siruta/delivery/middleware"
	"github.com/gofiber/fiber/v2"
)

func RegisterLetterReqRoutes(v1 fiber.Router, handler *http.LetterReqHttp) {
	v1.Post("/request-letters", middleware.JWTMiddleware([]int{3, 4}), handler.CreateData)
	v1.Put("/request-letters/approve/:letter_req_id", middleware.JWTMiddleware([]int{1, 3}), handler.UpdateApprovalStatus)
}
