package http

import (
	"github.com/abiyyu03/siruta/entity"
	"github.com/abiyyu03/siruta/entity/constant"
	"github.com/abiyyu03/siruta/entity/model"
	"github.com/abiyyu03/siruta/usecase/rt_profile"
	"github.com/gofiber/fiber/v2"
)

type RTLeaderHttp struct {
	rtLeaderUsecase rt_profile.RTLeaderUsecaseInterface
}

func NewRTLeaderHttp(rtLeaderUC rt_profile.RTLeaderUsecaseInterface) *RTLeaderHttp {
	return &RTLeaderHttp{
		rtLeaderUsecase: rtLeaderUC,
	}
}

func (r *RTLeaderHttp) GetData(ctx *fiber.Ctx) error {
	rtProfileId := ctx.Params("rt_profile_id")

	if rtProfileId == "" {
		return r.rtLeaderUsecase.Fetch(ctx)
	}

	return r.rtLeaderUsecase.FetchByRTProfileId(ctx, rtProfileId)
}

func (r *RTLeaderHttp) GetDataById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	return r.rtLeaderUsecase.FetchById(ctx, id)
}

func (r *RTLeaderHttp) UpdateData(ctx *fiber.Ctx) error {
	rtProfileId := ctx.Params("id")
	var rtLeaderData *model.RTLeader

	if err := ctx.BodyParser(&rtLeaderData); err != nil {
		return entity.Error(ctx, fiber.StatusUnprocessableEntity, constant.Errors["UnprocessableEntity"].Message, constant.Errors["UnprocessableEntity"].Clue)
	}

	return r.rtLeaderUsecase.Update(ctx, rtProfileId, rtLeaderData)
}
