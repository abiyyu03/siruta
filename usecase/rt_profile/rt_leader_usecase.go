package rt_profile

import (
	"github.com/abiyyu03/siruta/config"
	"github.com/abiyyu03/siruta/entity"
	"github.com/abiyyu03/siruta/entity/constant"
	"github.com/abiyyu03/siruta/entity/model"
	"github.com/abiyyu03/siruta/repository/rt_profile"
	"github.com/gofiber/fiber/v2"
)

type RTLeaderUsecase struct {
	rtLeaderRepository rt_profile.RTLeaderRepository
}

type RTLeaderUsecaseInterface interface {
	Fetch(ctx *fiber.Ctx) error
	FetchByRTProfileId(ctx *fiber.Ctx, rtProfileId string) error
	FetchById(ctx *fiber.Ctx, id string) error
	Update(ctx *fiber.Ctx, id string, rtLeaderData *model.RTLeader) error
}

func (r *RTLeaderUsecase) Fetch(ctx *fiber.Ctx) error {
	rtLeaders, err := r.rtLeaderRepository.Fetch()

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	return entity.Success(ctx, rtLeaders, "Data Fetched successfully")
}

func (r *RTLeaderUsecase) FetchByRTProfileId(ctx *fiber.Ctx, rtProfileId string) error {
	rtLeaders, err := r.rtLeaderRepository.FetchByRTProfileId(rtProfileId)

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	return entity.Success(ctx, rtLeaders, "Data Fetched successfully")
}

func (r *RTLeaderUsecase) FetchById(ctx *fiber.Ctx, id string) error {
	rtLeader, err := r.rtLeaderRepository.FetchById(id)

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	if rtLeader == nil {
		return entity.Error(ctx, fiber.StatusNotFound, constant.Errors["NotFound"].Message, constant.Errors["NotFound"].Clue)
	}

	return entity.Success(ctx, rtLeader, "Data Fetched successfully")
}

func (r *RTLeaderUsecase) Update(ctx *fiber.Ctx, id string, rtLeaderData *model.RTLeader) error {
	rtLeader := &model.RTLeader{
		Fullname:    rtLeaderData.Fullname,
		NikNumber:   rtLeaderData.NikNumber,
		KKNumber:    rtLeaderData.KKNumber,
		Photo:       rtLeaderData.Photo,
		FullAddress: rtLeaderData.FullAddress,
	}

	updatedRtLeader := r.rtLeaderRepository.Update(config.DB, rtLeader, id)

	if updatedRtLeader == nil {
		return entity.Error(ctx, fiber.StatusNotFound, constant.Errors["NotFound"].Message, constant.Errors["NotFound"].Clue)
	}

	return entity.Success(ctx, updatedRtLeader, "Data updated successfully")
}
