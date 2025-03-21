package letter

import (
	"github.com/abiyyu03/siruta/entity"
	"github.com/abiyyu03/siruta/entity/constant"
	"github.com/abiyyu03/siruta/repository/letter"
	"github.com/gofiber/fiber/v2"
)

type OutcomingLetterUsecase struct{}

var outcomingLetterRepository *letter.OutcomingLetterRepository

func (u *OutcomingLetterUsecase) Fetch(ctx *fiber.Ctx) error {
	letters, err := outcomingLetterRepository.Fetch()

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	return entity.Success(ctx, &letters, "Data fetched successfully")
}

func (u *OutcomingLetterUsecase) FetchById(ctx *fiber.Ctx, id string) error {
	letter, err := outcomingLetterRepository.FetchById(id)

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	if letter == nil {
		return nil
	}

	return entity.Success(ctx, &letter, "Data fetched successfully")
}

func (u *OutcomingLetterUsecase) FetchByRTProfileId(ctx *fiber.Ctx, rtProfileId string) error {
	letters, err := outcomingLetterRepository.FetchByRTProfileId(rtProfileId)

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	return entity.Success(ctx, &letters, "Data fetched successfully")
}
