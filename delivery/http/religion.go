package http

import (
	"strconv"

	"github.com/abiyyu03/siruta/entity/model"
	"github.com/abiyyu03/siruta/usecase"
	"github.com/gofiber/fiber/v2"
)

type ReligionHttp struct{}

var religionUsecase = new(usecase.ReligionUsecase)

func (v *ReligionHttp) GetData(ctx *fiber.Ctx) error {
	return religionUsecase.Fetch(ctx)
}

func (v *ReligionHttp) GetDataById(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	return religionUsecase.FetchById(ctx, id)
}

func (v *ReligionHttp) StoreData(ctx *fiber.Ctx) error {
	var religion *model.Religion

	if err := ctx.BodyParser(&religion); err != nil {
		return err
	}

	return religionUsecase.Store(religion, ctx)
}

func (v *ReligionHttp) UpdateData(ctx *fiber.Ctx) error {
	var religion *model.Religion
	id, _ := strconv.Atoi(ctx.Params("id"))

	if err := ctx.BodyParser(&religion); err != nil {
		return err
	}

	return religionUsecase.Update(religion, ctx, id)
}

func (v *ReligionHttp) DeleteData(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	return religionUsecase.Delete(ctx, id)
}
