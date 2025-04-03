package finance

import (
	"strconv"

	"github.com/abiyyu03/siruta/entity/model"
	"github.com/abiyyu03/siruta/usecase/finance"
	"github.com/gofiber/fiber/v2"
)

type ExpenseHttp struct {
	expenseUsecase *finance.ExpenseUsecase
}

func (e *ExpenseHttp) GetData(ctx *fiber.Ctx) error {
	return e.expenseUsecase.Fetch(ctx)
}

func (e *ExpenseHttp) GetDataByRTProfileId(ctx *fiber.Ctx) error {
	rtProfileId := ctx.Params("rt_profile_id")

	return e.expenseUsecase.FetchByRTProfileId(ctx, rtProfileId)
}

func (e *ExpenseHttp) GetDataById(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	return e.expenseUsecase.FetchById(ctx, id)
}

func (e *ExpenseHttp) StoreData(ctx *fiber.Ctx) error {
	var plan *model.Expense

	if err := ctx.BodyParser(&plan); err != nil {
		return err
	}

	return e.expenseUsecase.Store(ctx, plan)
}

func (e *ExpenseHttp) UpdateData(ctx *fiber.Ctx) error {
	var incomePlan *model.Expense
	id, _ := strconv.Atoi(ctx.Params("id"))

	if err := ctx.BodyParser(&incomePlan); err != nil {
		return err
	}

	return e.expenseUsecase.Update(ctx, incomePlan, id)
}

func (e *ExpenseHttp) DeleteData(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	return e.expenseUsecase.Delete(ctx, id)
}
