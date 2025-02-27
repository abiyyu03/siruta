package register

import (
	"github.com/abiyyu03/siruta/delivery/http/middleware"
	"github.com/abiyyu03/siruta/entity"
	"github.com/abiyyu03/siruta/entity/request"
	"github.com/abiyyu03/siruta/usecase"
	"github.com/gofiber/fiber/v2"
)

type RTProfileRegisterHttp struct{}

var rtProfileRegisterUsecase = new(usecase.RTProfileRegisterUsecase)

func (r *RTProfileRegisterHttp) RegisterRTProfile(ctx *fiber.Ctx) error {
	var rtProfile *request.RTProfileRegisterRequest

	if err := ctx.BodyParser(&rtProfile); err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, err.Error())
	}

	return rtProfileRegisterUsecase.RegisterRTProfile(rtProfile, ctx)
}

func (r *RTProfileRegisterHttp) ApproveRegistration(ctx *fiber.Ctx) error {
	rtProfileId := ctx.Params("rtProfileId")
	queryParam := ctx.Queries()

	if queryParam["email"] != "" {
		return rtProfileRegisterUsecase.Approve(queryParam["email"], rtProfileId, ctx)
	}

	return entity.Error(ctx, fiber.StatusBadRequest, "Email query parameter is required")
}

func (r *RTProfileRegisterHttp) RegisterUserRt(ctx *fiber.Ctx) error {
	params := ctx.Queries()
	//token verif
	isTokenValid := middleware.TokenRegisterValidator(params["token"])

	if !isTokenValid {
		return entity.Error(ctx, fiber.StatusForbidden, "Token tidak valid")
	}

	var userRt *request.RegisterRTRequest

	if err := ctx.BodyParser(&userRt); err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, err.Error())
	}

	return rtProfileRegisterUsecase.RegisterUserRt(userRt, ctx, params["token"])
}
