package usecase

import (
	"github.com/abiyyu03/siruta/entity"
	"github.com/abiyyu03/siruta/entity/model"
	"github.com/abiyyu03/siruta/repository"
	"github.com/gofiber/fiber/v2"
)

type LetterTypeUsecase struct{}

var letterTypeRepository = new(repository.LetterTypeRepository)

func (l *LetterTypeUsecase) Fetch(ctx *fiber.Ctx) error {
	letterTypes, err := letterTypeRepository.Fetch()

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, "Server error")
	}

	return entity.Success(ctx, &letterTypes, "Data fetched successfully")
}

func (m *LetterTypeUsecase) FetchById(ctx *fiber.Ctx, id string) error {
	letterType, err := letterTypeRepository.FetchById(id)

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, "Error fetching data")
	}

	if letterType == nil {
		return entity.Error(ctx, fiber.StatusNotFound, fiber.ErrNotFound.Message)
	}

	return entity.Success(ctx, &letterType, "Data fetched successfully")
}

func (v *LetterTypeUsecase) Store(letterType *model.LetterType, ctx *fiber.Ctx) error {
	createdLetterType := &model.LetterType{
		TypeName: letterType.TypeName,
		Code:     letterType.Code,
	}

	storedLetterType, err := letterTypeRepository.Store(createdLetterType)

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, fiber.ErrInternalServerError.Message)
	}

	return entity.Success(ctx, &storedLetterType, "Data updated successfully")
}

func (v *LetterTypeUsecase) Update(letterType *model.LetterType, ctx *fiber.Ctx, id int) error {

	updateLetterType := &model.LetterType{
		TypeName: letterType.TypeName,
		Code:     letterType.Code,
	}

	updatedLetterType, err := letterTypeRepository.Update(updateLetterType, id)

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, fiber.ErrInternalServerError.Message)
	}

	return entity.Success(ctx, &updatedLetterType, "Data updated successfully")
}

func (v *LetterTypeUsecase) Delete(ctx *fiber.Ctx, id int) error {
	village, err := letterTypeRepository.Delete(id)

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, "Server error")
	}

	return entity.Success(ctx, &village, "Data deleted successfully")
}
