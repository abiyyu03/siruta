package finance

import (
	"strconv"

	"github.com/abiyyu03/siruta/entity/model"
	"github.com/abiyyu03/siruta/usecase/finance"
	"github.com/gofiber/fiber/v2"
)

type IncomePlanHttp struct {
	incomePlanUsecase *finance.IncomePlanUsecase
}

func (i *IncomePlanHttp) GetData(ctx *fiber.Ctx) error {
	return i.incomePlanUsecase.Fetch(ctx)
}

func (i *IncomePlanHttp) GetDataByRTProfileId(ctx *fiber.Ctx) error {
	rtProfileId := ctx.Params("rt_profile_id")

	return i.incomePlanUsecase.FetchByRTProfileId(ctx, rtProfileId)
}

func (i *IncomePlanHttp) GetDataById(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	return i.incomePlanUsecase.FetchById(ctx, id)
}

func (i *IncomePlanHttp) StoreData(ctx *fiber.Ctx) error {
	var plan *model.IncomePlan

	if err := ctx.BodyParser(&plan); err != nil {
		return err
	}

	return i.incomePlanUsecase.Store(ctx, plan)
}

func (i *IncomePlanHttp) UpdateData(ctx *fiber.Ctx) error {
	var incomePlan *model.IncomePlan
	id, _ := strconv.Atoi(ctx.Params("id"))

	if err := ctx.BodyParser(&incomePlan); err != nil {
		return err
	}

	return i.incomePlanUsecase.Update(ctx, incomePlan, id)
}

func (i *IncomePlanHttp) DeleteData(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	return i.incomePlanUsecase.Delete(ctx, id)
}
