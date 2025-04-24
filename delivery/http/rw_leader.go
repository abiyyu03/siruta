package http

import (
	"github.com/abiyyu03/siruta/entity"
	"github.com/abiyyu03/siruta/entity/constant"
	"github.com/abiyyu03/siruta/entity/model"
	"github.com/abiyyu03/siruta/usecase/rw_profile"
	"github.com/gofiber/fiber/v2"
)

type RWLeaderHttp struct{}

var rwLeaderUsecase *rw_profile.RWLeaderUsecase

func (r *RWLeaderHttp) GetData(ctx *fiber.Ctx) error {
	return rwLeaderUsecase.Fetch(ctx)
}

func (r *RWLeaderHttp) GetDataById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	return rwLeaderUsecase.FetchById(ctx, id)
}

func (r *RWLeaderHttp) UpdateData(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	var rwLeaderData *model.RWLeader

	if err := ctx.BodyParser(&rwLeaderData); err != nil {
		return entity.Error(ctx, fiber.StatusUnprocessableEntity, constant.Errors["UnprocessableEntity"].Message, constant.Errors["UnprocessableEntity"].Clue)
	}

	return rwLeaderUsecase.Update(ctx, id, rwLeaderData)
}
