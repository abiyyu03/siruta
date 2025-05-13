package auth

import (
	"github.com/abiyyu03/siruta/entity"
	"github.com/abiyyu03/siruta/usecase/auth"
	"github.com/gofiber/fiber/v2"
)

type ResetPasswordHttp struct {
	resetPasswordUsecase auth.ResetPasswordUsecaseInterface
}

func NewResetPasswordHttp(resetUC auth.ResetPasswordUsecaseInterface) *ResetPasswordHttp {
	return &ResetPasswordHttp{
		resetPasswordUsecase: resetUC,
	}
}

type ForgotPasswordLink struct {
	Email string `json:"email" validate:"required;email"`
}

func (r *ResetPasswordHttp) SendForgotPasswordLink(ctx *fiber.Ctx) error {
	var reset *ForgotPasswordLink

	if err := ctx.BodyParser(&reset); err != nil {
		return err
	}

	return r.resetPasswordUsecase.SendForgotPasswordLink(ctx, reset.Email)
}

func (r *ResetPasswordHttp) ResetPassword(ctx *fiber.Ctx) error {
	query := ctx.Queries()
	var reset *entity.ResetPassword

	if err := ctx.BodyParser(&reset); err != nil {
		return err
	}

	return r.resetPasswordUsecase.ResetPassword(ctx, reset, query["token"])
}
