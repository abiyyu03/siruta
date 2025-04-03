package finance

import (
	"strconv"

	"github.com/abiyyu03/siruta/entity/model"
	"github.com/abiyyu03/siruta/usecase/finance"
	"github.com/gofiber/fiber/v2"
)

type IncomeHttp struct {
	incomeUsecase *finance.IncomeUsecase
}

func (i *IncomeHttp) GetData(ctx *fiber.Ctx) error {
	return i.incomeUsecase.Fetch(ctx)
}

func (i *IncomeHttp) GetDataByRTProfileId(ctx *fiber.Ctx) error {
	planId := ctx.Params("plan_id")

	return i.incomeUsecase.FetchByPlanId(ctx, planId)
}

func (i *IncomeHttp) GetDataById(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	return i.incomeUsecase.FetchById(ctx, id)
}

func (i *IncomeHttp) StoreData(ctx *fiber.Ctx) error {
	var plan *model.Income

	if err := ctx.BodyParser(&plan); err != nil {
		return err
	}

	return i.incomeUsecase.Store(ctx, plan)
}

func (i *IncomeHttp) UpdateData(ctx *fiber.Ctx) error {
	var incomePlan *model.Income
	id, _ := strconv.Atoi(ctx.Params("id"))

	if err := ctx.BodyParser(&incomePlan); err != nil {
		return err
	}

	return i.incomeUsecase.Update(ctx, incomePlan, id)
}

func (i *IncomeHttp) DeleteData(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	return i.incomeUsecase.Delete(ctx, id)
}
