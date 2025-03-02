package http

import (
	"strconv"

	"github.com/abiyyu03/siruta/entity/model"
	"github.com/abiyyu03/siruta/usecase/village"
	"github.com/gofiber/fiber/v2"
)

type VillageHttp struct {
	villageUsecase *village.VillageUsecase
}

func (v *VillageHttp) GetData(ctx *fiber.Ctx) error {
	return v.villageUsecase.Fetch(ctx)
}

func (v *VillageHttp) GetDataById(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	return v.villageUsecase.FetchById(ctx, id)
}

func (v *VillageHttp) StoreData(ctx *fiber.Ctx) error {
	var village *model.Village

	if err := ctx.BodyParser(&village); err != nil {
		return err
	}

	return v.villageUsecase.Store(village, ctx)
}

func (v *VillageHttp) UpdateData(ctx *fiber.Ctx) error {
	var village *model.Village
	id, _ := strconv.Atoi(ctx.Params("id"))

	if err := ctx.BodyParser(&village); err != nil {
		return err
	}

	return v.villageUsecase.Update(village, ctx, id)
}

func (v *VillageHttp) DeleteData(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	return v.villageUsecase.Delete(ctx, id)
}
