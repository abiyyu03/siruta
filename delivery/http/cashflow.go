package http

import (
	"strconv"

	"github.com/abiyyu03/siruta/entity/model"
	"github.com/abiyyu03/siruta/usecase/finance"
	"github.com/gofiber/fiber/v2"
)

type CashflowHttp struct{}

var cashflowUsecase *finance.CashflowUsecase

func (c *CashflowHttp) GetData(ctx *fiber.Ctx) error {
	query := ctx.Queries()

	return cashflowUsecase.Fetch(ctx, query["type"])
}

func (c *CashflowHttp) GetDataByRTProfileId(ctx *fiber.Ctx) error {
	rtProfileId := ctx.Params("rt_profile_id")
	query := ctx.Queries()

	return cashflowUsecase.FetchByRTProfileId(ctx, rtProfileId, query["type"])
}

func (c *CashflowHttp) GetDataById(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))
	query := ctx.Queries()

	return cashflowUsecase.FetchById(ctx, id, query["type"])
}

func (c *CashflowHttp) StoreData(ctx *fiber.Ctx) error {
	var cashflowExpense *model.Cashflow

	if err := ctx.BodyParser(&cashflowExpense); err != nil {
		return err
	}

	return cashflowUsecase.Store(ctx, cashflowExpense)
}

func (c *CashflowHttp) UpdateData(ctx *fiber.Ctx) error {
	var cashflowExpense *model.Cashflow
	id, _ := strconv.Atoi(ctx.Params("id"))

	if err := ctx.BodyParser(&cashflowExpense); err != nil {
		return err
	}

	return cashflowUsecase.Update(ctx, cashflowExpense, id)
}

func (c *CashflowHttp) DeleteData(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	return cashflowUsecase.Delete(ctx, id)
}
