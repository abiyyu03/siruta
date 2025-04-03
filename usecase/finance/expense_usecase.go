package finance

import (
	"github.com/abiyyu03/siruta/entity"
	"github.com/abiyyu03/siruta/entity/constant"
	"github.com/abiyyu03/siruta/entity/model"
	"github.com/abiyyu03/siruta/repository/finance"
	"github.com/gofiber/fiber/v2"
)

type ExpenseUsecase struct{}

var expenseRepository *finance.ExpenseRepository

func (e *ExpenseUsecase) Fetch(ctx *fiber.Ctx) error {
	expenses, err := expenseRepository.Fetch()

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	return entity.Success(ctx, &expenses, "Data fetched successfully")
}

func (e *ExpenseUsecase) FetchById(ctx *fiber.Ctx, id int) error {
	expense, err := expenseRepository.FetchById(id)

	if expense == nil {
		return entity.Error(ctx, fiber.StatusNotFound, constant.Errors["NotFound"].Message, constant.Errors["NotFound"].Clue)
	}

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	return entity.Success(ctx, &expense, "Data fetched successfully")
}
func (e *ExpenseUsecase) FetchByRTProfileId(ctx *fiber.Ctx, rtProfileId string) error {
	expenses, err := expenseRepository.FetchByRTProfileId(rtProfileId)

	if expenses == nil {
		return entity.Error(ctx, fiber.StatusNotFound, constant.Errors["NotFound"].Message, constant.Errors["NotFound"].Clue)
	}

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	return entity.Success(ctx, &expenses, "Data fetched successfully")
}

func (e *ExpenseUsecase) Store(ctx *fiber.Ctx, expenseData *model.Expense) error {
	expense, err := expenseRepository.Store(expenseData)

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	return entity.Success(ctx, expense, "Data stored successfully")
}

func (e *ExpenseUsecase) Update(ctx *fiber.Ctx, expenseData *model.Expense, id int) error {
	expense, err := expenseRepository.Update(expenseData, id)

	if expense == nil {
		return entity.Error(ctx, fiber.StatusNotFound, constant.Errors["NotFound"].Message, constant.Errors["NotFound"].Clue)
	}

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	return entity.Success(ctx, expense, "Data updated successfully")
}

func (e *ExpenseUsecase) Delete(ctx *fiber.Ctx, id int) error {
	_, err := expenseRepository.Delete(id)

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	return nil
}
