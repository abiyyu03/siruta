package religion

import (
	"github.com/abiyyu03/siruta/entity"
	"github.com/abiyyu03/siruta/entity/constant"
	"github.com/abiyyu03/siruta/entity/model"
	"github.com/abiyyu03/siruta/repository/religion"
	"github.com/gofiber/fiber/v2"
)

type ReligionUsecase struct {
	religionRepository *religion.ReligionRepository
}

func (r *ReligionUsecase) Fetch(ctx *fiber.Ctx) error {
	religions, err := r.religionRepository.Fetch()

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	return entity.Success(ctx, &religions, "Data fetched successfully")
}

func (r *ReligionUsecase) FetchById(ctx *fiber.Ctx, id int) error {
	religion, err := r.religionRepository.FetchById(id)

	if religion == nil {
		return entity.Error(ctx, fiber.StatusNotFound, constant.Errors["NotFound"].Message, constant.Errors["NotFound"].Clue)
	}

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	return entity.Success(ctx, &religion, "Data fetched successfully")
}

func (r *ReligionUsecase) Store(religion *model.Religion, ctx *fiber.Ctx) error {
	createdReligion := &model.Religion{
		ReligionName: religion.ReligionName,
	}

	storedReligion, err := r.religionRepository.Store(createdReligion)

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	return entity.Success(ctx, &storedReligion, "Data updated successfully")
}

func (r *ReligionUsecase) Update(religion *model.Religion, ctx *fiber.Ctx, id int) error {
	updateReligion := &model.Religion{
		ReligionName: religion.ReligionName,
	}

	updatedReligion, err := r.religionRepository.Update(updateReligion, id)

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	return entity.Success(ctx, &updatedReligion, "Data updated successfully")
}

func (r *ReligionUsecase) Delete(ctx *fiber.Ctx, id int) error {
	_, err := r.religionRepository.Delete(id)

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	return entity.Success(ctx, nil, "Data deleted successfully")
}
