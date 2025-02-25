package auth

import (
	"github.com/abiyyu03/siruta/entity"
	"github.com/abiyyu03/siruta/entity/request"
	"github.com/abiyyu03/siruta/usecase"
	"github.com/gofiber/fiber/v2"
)

var authUsecase = new(usecase.AuthUsecase)

func Login(ctx *fiber.Ctx) error {
	var request request.LoginRequest

	if err := ctx.BodyParser(&request); err != nil {
		return entity.Error(ctx, fiber.ErrBadRequest.Code, fiber.ErrBadRequest.Message)
	}

	response, err := authUsecase.IssueAuthToken(ctx, request.Username, request.Password)

	if err != nil {
		return entity.Error(ctx, fiber.ErrBadRequest.Code, fiber.ErrBadRequest.Message)
	}

	return entity.Success(ctx, response, "login successfully")
}
