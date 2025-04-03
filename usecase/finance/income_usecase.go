package finance

import (
	"time"

	"github.com/abiyyu03/siruta/entity"
	"github.com/abiyyu03/siruta/entity/constant"
	"github.com/abiyyu03/siruta/entity/model"
	"github.com/abiyyu03/siruta/repository/finance"
	"github.com/gofiber/fiber/v2"
)

type IncomeUsecase struct{}

var incomeRepository *finance.IncomeRepository

func (i *IncomeUsecase) Fetch(ctx *fiber.Ctx) error {
	incomeLogs, err := incomeRepository.Fetch()

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	return entity.Success(ctx, &incomeLogs, "Data fetched successfully")
}

func (i *IncomeUsecase) FetchById(ctx *fiber.Ctx, id int) error {
	incomeLog, err := incomeRepository.FetchById(id)

	if incomeLog == nil {
		return entity.Error(ctx, fiber.StatusNotFound, constant.Errors["NotFound"].Message, constant.Errors["NotFound"].Clue)
	}

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	return entity.Success(ctx, &incomeLog, "Data fetched successfully")
}
func (i *IncomeUsecase) FetchByPlanId(ctx *fiber.Ctx, planId string) error {
	incomeLogs, err := incomeRepository.FetchByPlanId(planId)

	if incomeLogs == nil {
		return entity.Error(ctx, fiber.StatusNotFound, constant.Errors["NotFound"].Message, constant.Errors["NotFound"].Clue)
	}

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	return entity.Success(ctx, &incomeLogs, "Data fetched successfully")
}

func (i *IncomeUsecase) Store(ctx *fiber.Ctx, incomeLogData *model.Income) error {
	newIncome := &model.Income{
		Amount:        incomeLogData.Amount,
		PaymentDate:   time.Now(),
		PlanId:        incomeLogData.PlanId,
		PaymentMethod: incomeLogData.PaymentMethod,
		PlanPeriod:    incomeLogData.PlanPeriod,
	}

	incomeLog, err := incomeRepository.Store(newIncome)

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	return entity.Success(ctx, incomeLog, "Data stored successfully")
}

func (i *IncomeUsecase) Update(ctx *fiber.Ctx, incomeLogData *model.Income, id int) error {
	updatedIncome := &model.Income{
		Amount:        incomeLogData.Amount,
		PlanId:        incomeLogData.PlanId,
		PaymentMethod: incomeLogData.PaymentMethod,
		PlanPeriod:    incomeLogData.PlanPeriod,
	}

	incomeLog, err := incomeRepository.Update(updatedIncome, id)

	if incomeLog == nil {
		return entity.Error(ctx, fiber.StatusNotFound, constant.Errors["NotFound"].Message, constant.Errors["NotFound"].Clue)
	}

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	return entity.Success(ctx, incomeLog, "Data updated successfully")
}

func (i *IncomeUsecase) Delete(ctx *fiber.Ctx, id int) error {
	_, err := incomeRepository.Delete(id)

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	return nil
}
