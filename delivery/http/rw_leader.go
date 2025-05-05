package http

import (
	"github.com/abiyyu03/siruta/entity"
	"github.com/abiyyu03/siruta/entity/constant"
	"github.com/abiyyu03/siruta/entity/model"
	"github.com/abiyyu03/siruta/usecase/rw_profile"
	"github.com/gofiber/fiber/v2"
)

type RWLeaderHttp struct {
	rwLeaderUsecase rw_profile.RWLeaderUsecaseInterface
}

func NewRWLeaderHttp(rwLeaderUC rw_profile.RWLeaderUsecaseInterface) *RWLeaderHttp {
	return &RWLeaderHttp{
		rwLeaderUsecase: rwLeaderUC,
	}
}

func (r *RWLeaderHttp) GetData(ctx *fiber.Ctx) error {
	return r.rwLeaderUsecase.Fetch(ctx)
}

func (r *RWLeaderHttp) GetDataById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	return r.rwLeaderUsecase.FetchById(ctx, id)
}

func (r *RWLeaderHttp) UpdateData(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	var rwLeaderData *model.RWLeader

	if err := ctx.BodyParser(&rwLeaderData); err != nil {
		return entity.Error(ctx, fiber.StatusUnprocessableEntity, constant.Errors["UnprocessableEntity"].Message, constant.Errors["UnprocessableEntity"].Clue)
	}

	return r.rwLeaderUsecase.Update(ctx, id, rwLeaderData)
}
