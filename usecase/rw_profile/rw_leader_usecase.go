package rw_profile

import (
	"github.com/abiyyu03/siruta/config"
	"github.com/abiyyu03/siruta/entity"
	"github.com/abiyyu03/siruta/entity/constant"
	"github.com/abiyyu03/siruta/entity/model"
	"github.com/abiyyu03/siruta/repository/rw_profile"
	"github.com/gofiber/fiber/v2"
)

type RWLeaderUsecase struct{}

var rwLeaderRepository *rw_profile.RWLeaderRepository

func (r *RWLeaderUsecase) Fetch(ctx *fiber.Ctx) error {
	rwLeaders, err := rwLeaderRepository.Fetch()

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	return entity.Success(ctx, rwLeaders, "Data fetched successfully")
}

func (r *RWLeaderUsecase) FetchById(ctx *fiber.Ctx, id string) error {
	rwLeader, err := rwLeaderRepository.FetchById(id)

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	if rwLeader == nil {
		return entity.Error(ctx, fiber.StatusNotFound, constant.Errors["NotFound"].Message, constant.Errors["NotFound"].Clue)
	}

	return entity.Success(ctx, rwLeader, "Data fetched successfully")
}

func (r *RWLeaderUsecase) Update(ctx *fiber.Ctx, id string, rwLeaderData *model.RWLeader) error {
	rwLeader := &model.RWLeader{
		Fullname:    rwLeaderData.Fullname,
		NikNumber:   rwLeaderData.NikNumber,
		KKNumber:    rwLeaderData.KKNumber,
		Photo:       rwLeaderData.Photo,
		FullAddress: rwLeaderData.FullAddress,
	}

	updatedRwLeader := rwLeaderRepository.Update(config.DB, rwLeader, id)

	if updatedRwLeader == nil {
		return entity.Error(ctx, fiber.StatusNotFound, constant.Errors["NotFound"].Message, constant.Errors["NotFound"].Clue)
	}

	return entity.Success(ctx, updatedRwLeader, "Data updated successfully")
}
