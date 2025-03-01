package register

import (
	"github.com/abiyyu03/siruta/delivery/http/middleware"
	"github.com/abiyyu03/siruta/entity"
	"github.com/abiyyu03/siruta/entity/constant"
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
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}
	return rwProfileRegisterUsecase.RegisterProfileRW(rwProfile, ctx)
}

func (r *RWProfileRegisterHttp) RegisterUserRw(ctx *fiber.Ctx) error {
	params := ctx.Queries()
	//token verif
	isTokenValid := middleware.TokenRegisterValidator(params["token"])

	if !isTokenValid {
		return entity.Error(ctx, fiber.StatusForbidden, constant.Errors["InvalidToken"].Message, constant.Errors["InvalidToken"].Clue)
	}

	var userRw *request.RegisterRequest

	if err := ctx.BodyParser(&userRw); err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	return rwProfileRegisterUsecase.RegisterUserRw(userRw, ctx, params["token"])
}

func (r *RWProfileRegisterHttp) ApproveRegistration(ctx *fiber.Ctx) error {
	rwProfileId := ctx.Params("rwProfileId")
	queryParam := ctx.Queries()

	if queryParam["email"] != "" {
		return rwProfileRegisterUsecase.Approve(queryParam["email"], rwProfileId, ctx)
	}

	return entity.Error(ctx, fiber.StatusUnprocessableEntity, constant.Errors["EmailQueryRequired"].Message, constant.Errors["EmailQueryRequired"].Clue)
}
