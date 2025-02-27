package usecase

import (
	"github.com/abiyyu03/siruta/entity"
	"github.com/abiyyu03/siruta/entity/constant"
	"github.com/abiyyu03/siruta/repository"
	"github.com/gofiber/fiber/v2"
)

type RWProfileUsecase struct{}

var rwProfileRepository = new(repository.RWProfileRepository)

func (u *RWProfileUsecase) Fetch(ctx *fiber.Ctx) error {
	rwProfiles, err := rwProfileRepository.Fetch()

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	return entity.Success(ctx, &rwProfiles, "Data fetched successfully")
}

func (u *RWProfileUsecase) FetchById(ctx *fiber.Ctx, id string) error {
	rwProfile, err := rwProfileRepository.FetchById(id)

	if err != nil {
		return entity.Error(ctx, fiber.StatusNotFound, constant.Errors["NotFound"].Message, constant.Errors["NotFound"].Clue)
	}

	return entity.Success(ctx, &rwProfile, "Data fetched successfully")
}
