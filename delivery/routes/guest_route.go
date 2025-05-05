package routes

import (
	"github.com/abiyyu03/siruta/delivery/http/auth"
	"github.com/abiyyu03/siruta/delivery/middleware"
	"github.com/abiyyu03/siruta/entity/request"
	"github.com/gofiber/fiber/v2"
)

func RegisterResetPasswordRoutes(v1 fiber.Router, handlerReset *auth.ResetPasswordHttp) {
	v1.Post("/forgot-password", handlerReset.SendForgotPasswordLink)
	v1.Put("/reset-password", middleware.ValidateField[request.ResetPassword](), handlerReset.ResetPassword)
}

func RegisterAuthRoutes(v1 fiber.Router, handlerReset *auth.AuthHttp) {
	v1.Post("/login", middleware.ValidateField[request.LoginRequest](), auth.Login)
}
