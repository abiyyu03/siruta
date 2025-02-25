package register

import (
	"github.com/abiyyu03/siruta/delivery/http/middleware"
	"github.com/abiyyu03/siruta/entity"
	"github.com/abiyyu03/siruta/entity/model"
	"github.com/abiyyu03/siruta/entity/request"
	"github.com/abiyyu03/siruta/usecase"
	"github.com/gofiber/fiber/v2"
)

type RWProfileRegisterHttp struct{}

var rwProfileRegisterUsecase = new(usecase.RWProfileRegisterUsecase)

func (r *RWProfileRegisterHttp) RegisterRWProfile(ctx *fiber.Ctx) error {
	var rwProfile *model.RWProfile

	if err := ctx.BodyParser(&rwProfile); err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, err.Error())
	}
	return rwProfileRegisterUsecase.RegisterProfileRW(rwProfile, ctx)
}

func (r *RWProfileRegisterHttp) RegisterUserRw(ctx *fiber.Ctx) error {
	params := ctx.Queries()
	//token verif
	isTokenValid := middleware.TokenRegisterValidator(params["token"])

	if !isTokenValid {
		return entity.Error(ctx, fiber.StatusForbidden, "Token tidak valid")
	}

	var userRw *request.RegisterRWRequest

	if err := ctx.BodyParser(&userRw); err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, err.Error())
	}

	return rwProfileRegisterUsecase.RegisterUserRW(userRw, ctx, params["token"])
}

func (r *RWProfileRegisterHttp) ApproveRegistration(ctx *fiber.Ctx) error {
	rwProfileId := ctx.Params("id")
	queryParam := ctx.Queries()

	if queryParam["email"] != "" {
		return rwProfileRegisterUsecase.Approve(queryParam["email"], rwProfileId, ctx)
	}

	return entity.Error(ctx, fiber.StatusBadRequest, "Email query parameter is required")
}
