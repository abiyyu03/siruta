package usecase

import (
	"github.com/abiyyu03/siruta/entity"
	"github.com/abiyyu03/siruta/repository"
	"github.com/gofiber/fiber/v2"
)

type RTProfileUsecase struct{}

var rtProfileRepository = new(repository.RTProfileRepository)

func (r *RTProfileUsecase) Fetch(ctx *fiber.Ctx) error {
	rtProfiles, err := rtProfileRepository.Fetch()

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, "Server error")
	}

	return entity.Success(ctx, &rtProfiles, "Data fetched successfully")
}

func (r *RTProfileUsecase) FetchById(ctx *fiber.Ctx, id string) error {
	rtProfile, err := rtProfileRepository.FetchById(id)

	if err != nil {
		return entity.Error(ctx, fiber.StatusNotFound, "RT Profile not found")
	}

	return entity.Success(ctx, &rtProfile, "Data fetched successfully")
}
