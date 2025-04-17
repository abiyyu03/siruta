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
	outcomeLetters, err := outcomingLetterRepository.Fetch()

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	return entity.Success(ctx, &outcomeLetters, "Data fetched successfully")
}

func (u *OutcomingLetterUsecase) FetchByRtProfileId(ctx *fiber.Ctx, rtProfileId string) error {
	outcomeLetters, err := outcomingLetterRepository.FetchByRtProfileId(rtProfileId)

	if outcomeLetters == nil {
		return entity.Error(ctx, fiber.StatusNotFound, constant.Errors["NotFound"].Message, constant.Errors["NotFound"].Clue)
	}

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	return entity.Success(ctx, &outcomeLetters, "Data fetched successfully")
}
