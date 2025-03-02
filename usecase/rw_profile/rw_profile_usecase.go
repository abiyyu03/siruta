package rw_profile

import (
	"github.com/abiyyu03/siruta/entity"
	"github.com/abiyyu03/siruta/entity/constant"
	"github.com/abiyyu03/siruta/repository/rw_profile"
	"github.com/gofiber/fiber/v2"
)

type RWProfileUsecase struct {
	rwProfileRepository *rw_profile.RWProfileRepository
}

func (r *RWProfileUsecase) Fetch(ctx *fiber.Ctx) error {
	rwProfiles, err := r.rwProfileRepository.Fetch()

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	return entity.Success(ctx, &rwProfiles, "Data fetched successfully")
}

func (r *RWProfileUsecase) FetchById(ctx *fiber.Ctx, id string) error {
	rwProfile, err := r.rwProfileRepository.FetchById(id)

	if err != nil {
		return entity.Error(ctx, fiber.StatusNotFound, constant.Errors["NotFound"].Message, constant.Errors["NotFound"].Clue)
	}

	return entity.Success(ctx, &rwProfile, "Data fetched successfully")
}
