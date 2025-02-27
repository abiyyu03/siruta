package register

import (
	"github.com/abiyyu03/siruta/delivery/http/middleware"
	"github.com/abiyyu03/siruta/entity"
	"github.com/abiyyu03/siruta/entity/constant"
	"github.com/abiyyu03/siruta/entity/request"
	"github.com/abiyyu03/siruta/usecase"
	"github.com/gofiber/fiber/v2"
)

type RTProfileRegisterHttp struct{}

var rtProfileRegisterUsecase = new(usecase.RTProfileRegisterUsecase)

func (r *RTProfileRegisterHttp) RegisterRTProfile(ctx *fiber.Ctx) error {
	var rtProfile *request.RTProfileRegisterRequest

	if err := ctx.BodyParser(&rtProfile); err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	return rtProfileRegisterUsecase.RegisterRTProfile(rtProfile, ctx)
}

func (r *RTProfileRegisterHttp) ApproveRegistration(ctx *fiber.Ctx) error {
	rtProfileId := ctx.Params("rtProfileId")
	queryParam := ctx.Queries()

	if queryParam["email"] != "" {
		return rtProfileRegisterUsecase.Approve(queryParam["email"], rtProfileId, ctx)
	}

	return entity.Error(ctx, fiber.StatusUnprocessableEntity, constant.Errors["EmailQueryRequired"].Message, constant.Errors["EmailQueryRequired"].Clue)
}

func (r *RTProfileRegisterHttp) RegisterUserRt(ctx *fiber.Ctx) error {
	params := ctx.Queries()
	//token verif
	isTokenValid := middleware.TokenRegisterValidator(params["token"])

	if !isTokenValid {
		return entity.Error(ctx, fiber.StatusForbidden, constant.Errors["InvalidToken"].Message, constant.Errors["InvalidToken"].Clue)
	}

	var userRt *request.RegisterRTRequest

	if err := ctx.BodyParser(&userRt); err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	return rtProfileRegisterUsecase.RegisterUserRt(userRt, ctx, params["token"])
}
