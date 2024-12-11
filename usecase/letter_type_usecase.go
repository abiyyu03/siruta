package usecase

import (
	"github.com/abiyyu03/siruta/entity"
	"github.com/abiyyu03/siruta/repository"
	"github.com/gofiber/fiber/v2"
)

type LetterTypeUsecase struct{}

var letterTypeRepository = new(repository.LetterTypeRepository)

func (l *LetterTypeUsecase) Fetch(ctx *fiber.Ctx) error {
	letterTypes, err := letterTypeRepository.Fetch()

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, "Server error")
	}

	if letterTypes != nil {
		return nil
	}

	return entity.Success(ctx, &letterTypes, "Data fetched successfully")
}

func (m *LetterTypeUsecase) FetchById(ctx *fiber.Ctx, id string) error {
	letterType, err := letterTypeRepository.FetchById(id)

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, "Error fetching data")
	}

	if letterType == nil {
		return entity.Error(ctx, fiber.StatusNotFound, fiber.ErrNotFound.Message)
	}

	return entity.Success(ctx, &letterType, "Data fetched successfully")
}
