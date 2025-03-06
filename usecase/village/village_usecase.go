package village

import (
	"github.com/abiyyu03/siruta/entity"
	"github.com/abiyyu03/siruta/entity/constant"
	"github.com/abiyyu03/siruta/entity/model"
	"github.com/abiyyu03/siruta/repository/village"
	"github.com/gofiber/fiber/v2"
)

type VillageUsecase struct{}

var villageRepository *village.VillageRepository

func (v *VillageUsecase) Fetch(ctx *fiber.Ctx) error {
	villages, err := villageRepository.Fetch()

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	return entity.Success(ctx, &villages, "Data fetched successfully")
}

func (v *VillageUsecase) FetchById(ctx *fiber.Ctx, id int) error {
	village, err := villageRepository.FetchById(id)

	if village == nil {
		return entity.Error(ctx, fiber.StatusNotFound, constant.Errors["NotFound"].Message, constant.Errors["NotFound"].Clue)
	}

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	return entity.Success(ctx, &village, "Data fetched successfully")
}

func (v *VillageUsecase) Store(village *model.Village, ctx *fiber.Ctx) error {
	createdVillage := &model.Village{
		Name:       village.Name,
		AltName:    village.AltName,
		Latitude:   village.Latitude,
		Longitude:  village.Longitude,
		CodePostal: village.CodePostal,
	}

	storedVillage, err := villageRepository.Store(createdVillage)

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	return entity.Success(ctx, &storedVillage, "Data updated successfully")
}

func (v *VillageUsecase) Update(village *model.Village, ctx *fiber.Ctx, id int) error {

	updateVillage := &model.Village{
		Name:       village.Name,
		AltName:    village.AltName,
		Latitude:   village.Latitude,
		Longitude:  village.Longitude,
		CodePostal: village.CodePostal,
	}

	updatedVillage, err := villageRepository.Update(updateVillage, id)

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	return entity.Success(ctx, &updatedVillage, "Data updated successfully")
}

func (v *VillageUsecase) Delete(ctx *fiber.Ctx, id int) error {
	_, err := villageRepository.Delete(id)

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	return entity.Success(ctx, nil, "Data deleted successfully")
}
