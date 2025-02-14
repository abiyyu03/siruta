package usecase

import (
	"github.com/abiyyu03/siruta/entity"
	"github.com/abiyyu03/siruta/entity/model"
	"github.com/abiyyu03/siruta/repository"
	"github.com/gofiber/fiber/v2"
)

type ReligionUsecase struct{}

var ReligionRepository = new(repository.ReligionRepository)

func (v *ReligionUsecase) Fetch(ctx *fiber.Ctx) error {
	religions, err := ReligionRepository.Fetch()

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, "Server error")
	}

	return entity.Success(ctx, &religions, "Data fetched successfully")
}

func (v *ReligionUsecase) FetchById(ctx *fiber.Ctx, id int) error {
	religion, err := ReligionRepository.FetchById(id)

	if religion == nil {
		return entity.Error(ctx, fiber.StatusNotFound, "Religion not found")
	}

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, fiber.ErrInternalServerError.Message)
	}

	return entity.Success(ctx, &religion, "Data fetched successfully")
}

func (v *ReligionUsecase) Store(religion *model.Religion, ctx *fiber.Ctx) error {
	createdReligion := &model.Religion{
		ReligionName: religion.ReligionName,
	}

	storedReligion, err := ReligionRepository.Store(createdReligion)

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, fiber.ErrInternalServerError.Message)
	}

	return entity.Success(ctx, &storedReligion, "Data updated successfully")
}

func (v *ReligionUsecase) Update(religion *model.Religion, ctx *fiber.Ctx, id int) error {
	updateReligion := &model.Religion{
		ReligionName: religion.ReligionName,
	}

	updatedReligion, err := ReligionRepository.Update(updateReligion, id)

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, fiber.ErrInternalServerError.Message)
	}

	return entity.Success(ctx, &updatedReligion, "Data updated successfully")
}

func (v *ReligionUsecase) Delete(ctx *fiber.Ctx, id int) error {
	_, err := ReligionRepository.Delete(id)

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, "Server error")
	}

	return entity.Success(ctx, nil, "Data deleted successfully")
}
