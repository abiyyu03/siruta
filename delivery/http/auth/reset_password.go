package auth

import (
	"github.com/abiyyu03/siruta/entity/request"
	"github.com/abiyyu03/siruta/usecase/auth"
	"github.com/gofiber/fiber/v2"
)

type ResetPasswordHttp struct{}

var resetPasswordUsecase *auth.ResetPasswordUsecase

type ForgotPasswordLink struct {
	Email string `json:"email" validate:"required;email"`
}

func (r *ResetPasswordHttp) SendForgotPasswordLink(ctx *fiber.Ctx) error {
	var reset *ForgotPasswordLink

	if err := ctx.BodyParser(&reset); err != nil {
		return err
	}

	return resetPasswordUsecase.SendForgotPasswordLink(ctx, reset.Email)
}

func (r *ResetPasswordHttp) ResetPassword(ctx *fiber.Ctx) error {
	query := ctx.Queries()
	var reset *request.ResetPassword

	if err := ctx.BodyParser(&reset); err != nil {
		return err
	}

	return resetPasswordUsecase.ResetPassword(ctx, reset, query["token"])
}
