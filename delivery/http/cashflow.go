package http

import (
	"strconv"

	"github.com/abiyyu03/siruta/entity/model"
	"github.com/abiyyu03/siruta/usecase/finance"
	"github.com/gofiber/fiber/v2"
)

type CashflowHttp struct {
	cashflowUsecase finance.CashflowUsecaseInterface
}

func NewCashflowHttp(cashflowUC finance.CashflowUsecaseInterface) *CashflowHttp {
	return &CashflowHttp{
		cashflowUsecase: cashflowUC,
	}
}

func (c *CashflowHttp) GetData(ctx *fiber.Ctx) error {
	query := ctx.Queries()

	return c.cashflowUsecase.Fetch(ctx, query["type"])
}

func (c *CashflowHttp) GetDataByRTProfileId(ctx *fiber.Ctx) error {
	rtProfileId := ctx.Params("rt_profile_id")
	query := ctx.Queries()

	return c.cashflowUsecase.FetchByRTProfileId(ctx, rtProfileId, query["type"])
}

func (c *CashflowHttp) GetDataById(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))
	query := ctx.Queries()

	return c.cashflowUsecase.FetchById(ctx, id, query["type"])
}

func (c *CashflowHttp) StoreData(ctx *fiber.Ctx) error {
	var cashflowExpense *model.Cashflow

	if err := ctx.BodyParser(&cashflowExpense); err != nil {
		return err
	}

	return c.cashflowUsecase.Store(ctx, cashflowExpense)
}

func (c *CashflowHttp) UpdateData(ctx *fiber.Ctx) error {
	var cashflowExpense *model.Cashflow
	id, _ := strconv.Atoi(ctx.Params("id"))

	if err := ctx.BodyParser(&cashflowExpense); err != nil {
		return err
	}

	return c.cashflowUsecase.Update(ctx, cashflowExpense, id)
}

func (c *CashflowHttp) DeleteData(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	return c.cashflowUsecase.Delete(ctx, id)
}
