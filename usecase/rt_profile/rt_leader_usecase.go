package rt_profile

import (
	"log"

	"github.com/abiyyu03/siruta/config"
	"github.com/abiyyu03/siruta/entity"
	"github.com/abiyyu03/siruta/entity/constant"
	"github.com/abiyyu03/siruta/entity/model"
	"github.com/abiyyu03/siruta/repository/rt_profile"
	"github.com/gofiber/fiber/v2"
)

type RTLeaderUsecase struct {
}

var rtLeaderRepository *rt_profile.RTLeaderRepository

type RTLeaderUsecaseInterface interface {
	Fetch(ctx *fiber.Ctx) error
	FetchByRTProfileId(ctx *fiber.Ctx, rtProfileId string) error
	FetchById(ctx *fiber.Ctx, id string) error
	Update(ctx *fiber.Ctx, id string, rtLeaderData *model.RTLeader) error
}

func (r *RTLeaderUsecase) Fetch(ctx *fiber.Ctx) error {
	rtLeaders, err := rtLeaderRepository.Fetch()

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	return entity.Success(ctx, rtLeaders, "Data Fetched successfully")
}

func (r *RTLeaderUsecase) FetchByRTProfileId(ctx *fiber.Ctx, rtProfileId string) error {
	rtLeaders, err := rtLeaderRepository.FetchByRTProfileId(rtProfileId)

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	return entity.Success(ctx, rtLeaders, "Data Fetched successfully")
}

func (r *RTLeaderUsecase) FetchById(ctx *fiber.Ctx, id string) error {
	rtLeader, err := rtLeaderRepository.FetchById(id)

	if rtLeader == nil {
		return entity.Error(ctx, fiber.StatusNotFound, constant.Errors["NotFound"].Message, constant.Errors["NotFound"].Clue)
	}

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	result := &model.RTLeader{
		ID:          rtLeader.ID,
		Fullname:    rtLeader.Fullname,
		NikNumber:   rtLeader.NikNumber,
		KKNumber:    rtLeader.KKNumber,
		RTProfileId: rtLeader.RTProfileId,
		Photo:       rtLeader.Photo,
		UserId:      rtLeader.UserId,
		FullAddress: rtLeader.FullAddress,
	}

	return entity.Success(ctx, result, "Data Fetched successfully")
}

func (r *RTLeaderUsecase) Update(ctx *fiber.Ctx, id string, rtLeaderData *model.RTLeader) error {
	rtLeader := &model.RTLeader{
		Fullname:    rtLeaderData.Fullname,
		NikNumber:   rtLeaderData.NikNumber,
		KKNumber:    rtLeaderData.KKNumber,
		Photo:       rtLeaderData.Photo,
		FullAddress: rtLeaderData.FullAddress,
	}

	err := rtLeaderRepository.Update(config.DB, rtLeader, id)

	if err != nil {
		log.Print(err)
		return entity.Error(ctx, fiber.StatusNotFound, constant.Errors["NotFound"].Message, constant.Errors["NotFound"].Clue)
	}

	return entity.Success(ctx, nil, "Data updated successfully")
}
