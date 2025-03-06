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
	members, err := outcomingLetterRepository.Fetch()

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	if members == nil {
		return nil
	}

	return entity.Success(ctx, &members, "Data fetched successfully")
}
