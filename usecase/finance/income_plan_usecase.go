package finance

import (
	"github.com/abiyyu03/siruta/entity"
	"github.com/abiyyu03/siruta/entity/constant"
	"github.com/abiyyu03/siruta/entity/model"
	"github.com/abiyyu03/siruta/repository/finance"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type IncomePlanUsecase struct{}

var IncomePlanRepository *finance.IncomePlanRepository

func (i *IncomePlanUsecase) Fetch(ctx *fiber.Ctx) error {
	incomePlans, err := IncomePlanRepository.Fetch()

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	return entity.Success(ctx, &incomePlans, "Data fetched successfully")
}

func (i *IncomePlanUsecase) FetchById(ctx *fiber.Ctx, id int) error {
	incomePlan, err := IncomePlanRepository.FetchById(id)

	if incomePlan == nil {
		return entity.Error(ctx, fiber.StatusNotFound, constant.Errors["NotFound"].Message, constant.Errors["NotFound"].Clue)
	}

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	return entity.Success(ctx, &incomePlan, "Data fetched successfully")
}
func (i *IncomePlanUsecase) FetchByRTProfileId(ctx *fiber.Ctx, rtProfileId string) error {
	incomePlan, err := IncomePlanRepository.FetchByRTProfileId(rtProfileId)

	if incomePlan == nil {
		return entity.Error(ctx, fiber.StatusNotFound, constant.Errors["NotFound"].Message, constant.Errors["NotFound"].Clue)
	}

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	return entity.Success(ctx, &incomePlan, "Data fetched successfully")
}

func (i *IncomePlanUsecase) Store(ctx *fiber.Ctx, incomePlanData *model.IncomePlan) error {
	id, _ := uuid.NewV7()

	newIncomePlan := &model.IncomePlan{
		ID:            id.String(),
		PlanName:      incomePlanData.PlanName,
		RTProfileId:   incomePlanData.RTProfileId,
		StartPlan:     incomePlanData.StartPlan,
		EndPlan:       incomePlanData.EndPlan,
		Description:   incomePlanData.Description,
		IsSetDeadline: incomePlanData.IsSetDeadline,
	}

	incomePlan, err := IncomePlanRepository.Store(newIncomePlan)

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	return entity.Success(ctx, incomePlan, "Data stored successfully")
}

func (i *IncomePlanUsecase) Update(ctx *fiber.Ctx, incomePlanData *model.IncomePlan, id int) error {
	updatedIncomePlan := &model.IncomePlan{
		PlanName:      incomePlanData.PlanName,
		StartPlan:     incomePlanData.StartPlan,
		EndPlan:       incomePlanData.EndPlan,
		Description:   incomePlanData.Description,
		IsSetDeadline: incomePlanData.IsSetDeadline,
		IsClosed:      incomePlanData.IsClosed,
	}

	incomePlan, err := IncomePlanRepository.Update(updatedIncomePlan, id)

	if incomePlan == nil {
		return entity.Error(ctx, fiber.StatusNotFound, constant.Errors["NotFound"].Message, constant.Errors["NotFound"].Clue)
	}

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	return entity.Success(ctx, incomePlan, "Data updated successfully")
}

func (i *IncomePlanUsecase) Delete(ctx *fiber.Ctx, id int) error {
	_, err := IncomePlanRepository.Delete(id)

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	return nil
}
