package rt_profile

import (
	"github.com/abiyyu03/siruta/entity"
	"github.com/abiyyu03/siruta/entity/constant"
	"github.com/abiyyu03/siruta/repository/rt_profile"
	"github.com/gofiber/fiber/v2"
)

type RTProfileUsecase struct{}

var rtProfileRepository *rt_profile.RTProfileRepository

func (r *RTProfileUsecase) Fetch(ctx *fiber.Ctx) error {
	rtProfiles, err := rtProfileRepository.Fetch()

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	return entity.Success(ctx, &rtProfiles, "Data fetched successfully")
}

func (r *RTProfileUsecase) FetchById(ctx *fiber.Ctx, id string) error {
	rtProfile, err := rtProfileRepository.FetchById(id)

	if err != nil {
		return entity.Error(ctx, fiber.StatusNotFound, constant.Errors["NotFound"].Message, constant.Errors["NotFound"].Clue)
	}

	return entity.Success(ctx, &rtProfile, "Data fetched successfully")
}

func (r *RTProfileUsecase) FetchByRtProfileId(ctx *fiber.Ctx, rtProfileId string) error {
	rtProfiles, err := rtProfileRepository.FetchByRtProfileId(rtProfileId)

	if rtProfiles == nil {
		return entity.Error(ctx, fiber.StatusNotFound, constant.Errors["NotFound"].Message, constant.Errors["NotFound"].Clue)
	}

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	return entity.Success(ctx, &rtProfiles, "Data fetched successfully")
}
