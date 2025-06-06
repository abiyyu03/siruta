package auth

import (
	"github.com/abiyyu03/siruta/entity"
	"github.com/abiyyu03/siruta/entity/constant"
	"github.com/abiyyu03/siruta/usecase/auth"
	"github.com/gofiber/fiber/v2"
)

var authUsecase *auth.AuthUsecase

type AuthHttp struct {
	authUsecase auth.AuthUsecaseInterface
}

func NewAuthHttp(authUC auth.AuthUsecaseInterface) *AuthHttp {
	return &AuthHttp{
		authUsecase: authUC,
	}
}

func Login(ctx *fiber.Ctx) error {
	var request entity.LoginRequest

	if err := ctx.BodyParser(&request); err != nil {
		return entity.Error(
			ctx,
			fiber.ErrBadRequest.Code,
			constant.Errors["UnprocessableEntity"].Message,
			constant.Errors["UnprocessableEntity"].Clue,
		)
	}

	return authUsecase.IssueAuthToken(ctx, request.Email, request.Password)

}
