package http

import (
	"strconv"

	"github.com/abiyyu03/siruta/entity/model"
	"github.com/abiyyu03/siruta/usecase/religion"
	"github.com/gofiber/fiber/v2"
)

type ReligionHttp struct {
	religionUsecase religion.ReligionUsecaseInterface
}

func NewReligionHttp(religionUC religion.ReligionUsecaseInterface) *ReligionHttp {
	return &ReligionHttp{
		religionUsecase: religionUC,
	}
}

func (r *ReligionHttp) GetData(ctx *fiber.Ctx) error {
	return r.religionUsecase.Fetch(ctx)
}

func (r *ReligionHttp) GetDataById(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	return r.religionUsecase.FetchById(ctx, id)
}

func (r *ReligionHttp) StoreData(ctx *fiber.Ctx) error {
	var religion *model.Religion

	if err := ctx.BodyParser(&religion); err != nil {
		return err
	}

	return r.religionUsecase.Store(religion, ctx)
}

func (r *ReligionHttp) UpdateData(ctx *fiber.Ctx) error {
	var religion *model.Religion
	id, _ := strconv.Atoi(ctx.Params("id"))

	if err := ctx.BodyParser(&religion); err != nil {
		return err
	}

	return r.religionUsecase.Update(religion, ctx, id)
}

func (r *ReligionHttp) DeleteData(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	return r.religionUsecase.Delete(ctx, id)
}
