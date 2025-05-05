package routes

import (
	"github.com/abiyyu03/siruta/delivery/http"
	"github.com/abiyyu03/siruta/delivery/middleware"
	"github.com/gofiber/fiber/v2"
)

func RegisterLetterReqRoutes(v1 fiber.Router, handler *http.LetterReqHttp) {
	adminOnly := middleware.JWTMiddleware([]int{1})
	rtLeaderOnly := middleware.JWTMiddleware([]int{3})
	memberOnly := middleware.JWTMiddleware([]int{4})

	v1.Post("/request-letters", rtLeaderOnly, memberOnly, handler.CreateData)
	v1.Put("/request-letters/approve/:letter_req_id", rtLeaderOnly, handler.UpdateApprovalStatus)
	v1.Post("/request-letters", adminOnly, memberOnly, handler.CreateData)
}
