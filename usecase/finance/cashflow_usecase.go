package finance

import (
	"github.com/abiyyu03/siruta/entity"
	"github.com/abiyyu03/siruta/entity/constant"
	"github.com/abiyyu03/siruta/entity/model"
	"github.com/abiyyu03/siruta/repository/finance"
	"github.com/gofiber/fiber/v2"
)

type CashflowUsecase struct{}

var cashflowRepository *finance.CashflowRepository

func (c *CashflowUsecase) Fetch(ctx *fiber.Ctx, logType string) error {
	cashlog, err := cashflowRepository.Fetch(logType)

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	return entity.Success(ctx, &cashlog, "Data fetched successfully")
}

func (c *CashflowUsecase) FetchById(ctx *fiber.Ctx, id int, logType string) error {
	cashlog, err := cashflowRepository.FetchById(id, logType)

	if cashlog == nil {
		return entity.Error(ctx, fiber.StatusNotFound, constant.Errors["NotFound"].Message, constant.Errors["NotFound"].Clue)
	}

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	return entity.Success(ctx, &cashlog, "Data fetched successfully")
}
func (c *CashflowUsecase) FetchByRTProfileId(ctx *fiber.Ctx, rtProfileId string, logType string) error {
	cashlog, err := cashflowRepository.FetchByRTProfileId(rtProfileId, logType)

	if cashlog == nil {
		return entity.Error(ctx, fiber.StatusNotFound, constant.Errors["NotFound"].Message, constant.Errors["NotFound"].Clue)
	}

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	return entity.Success(ctx, &cashlog, "Data fetched successfully")
}

func (c *CashflowUsecase) Store(ctx *fiber.Ctx, cashlogData *model.Cashflow) error {
	cashlog, err := cashflowRepository.Store(cashlogData)

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	return entity.Success(ctx, cashlog, "Data stored successfully")
}

func (c *CashflowUsecase) Update(ctx *fiber.Ctx, cashlogData *model.Cashflow, id int) error {
	cashlog, err := cashflowRepository.Update(cashlogData, id)

	if cashlog == nil {
		return entity.Error(ctx, fiber.StatusNotFound, constant.Errors["NotFound"].Message, constant.Errors["NotFound"].Clue)
	}

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	return entity.Success(ctx, cashlog, "Data updated successfully")
}

func (c *CashflowUsecase) Delete(ctx *fiber.Ctx, id int) error {
	err := cashflowRepository.Delete(id)

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	return nil
}
