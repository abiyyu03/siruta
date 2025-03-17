package http

import (
	"github.com/abiyyu03/siruta/entity"
	"github.com/abiyyu03/siruta/entity/constant"
	"github.com/abiyyu03/siruta/entity/model"
	"github.com/abiyyu03/siruta/usecase/rt_profile"
	"github.com/gofiber/fiber/v2"
)

type RTLeaderHttp struct{}

var rtLeaderUsecase *rt_profile.RTLeaderUsecase

func (r *RTLeaderHttp) GetData(ctx *fiber.Ctx) error {
	rtProfileId := ctx.Params("rt_profile_id")

	if rtProfileId == "" {
		return rtLeaderUsecase.Fetch(ctx)
	}

	return rtLeaderUsecase.FetchByRTProfileId(ctx, rtProfileId)
}

func (r *RTLeaderHttp) GetDataById(ctx *fiber.Ctx) error {
	rtProfileId := ctx.Params("id")

	return rtLeaderUsecase.FetchByRTProfileId(ctx, rtProfileId)
}

func (r *RTLeaderHttp) UpdateData(ctx *fiber.Ctx) error {
	rtProfileId := ctx.Params("id")
	var rtLeaderData *model.RTLeader

	if err := ctx.BodyParser(&rtLeaderData); err != nil {
		return entity.Error(ctx, fiber.StatusUnprocessableEntity, constant.Errors["UnprocessableEntity"].Message, constant.Errors["UnprocessableEntity"].Clue)
	}

	return rtLeaderUsecase.Update(ctx, rtProfileId, rtLeaderData)
}
