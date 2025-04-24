package letter_type

import (
	"github.com/abiyyu03/siruta/entity"
	"github.com/abiyyu03/siruta/entity/constant"
	"github.com/abiyyu03/siruta/entity/model"
	"github.com/abiyyu03/siruta/repository/letter_type"
	"github.com/gofiber/fiber/v2"
)

type LetterTypeUsecase struct{}

var letterTypeRepository *letter_type.LetterTypeRepository

func (l *LetterTypeUsecase) Fetch(ctx *fiber.Ctx) error {
	letterTypes, err := letterTypeRepository.Fetch()

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	return entity.Success(ctx, &letterTypes, "Data fetched successfully")
}

func (l *LetterTypeUsecase) FetchById(ctx *fiber.Ctx, id int) error {
	letterType, err := letterTypeRepository.FetchById(id)

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	if letterType == nil {
		return entity.Error(ctx, fiber.StatusNotFound, constant.Errors["NotFound"].Message, constant.Errors["NotFound"].Clue)
	}

	return entity.Success(ctx, &letterType, "Data fetched successfully")
}

func (l *LetterTypeUsecase) Store(letterType *model.LetterType, ctx *fiber.Ctx) error {
	createdLetterType := &model.LetterType{
		TypeName: letterType.TypeName,
		Code:     letterType.Code,
	}

	storedLetterType, err := letterTypeRepository.Store(createdLetterType)

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	return entity.Success(ctx, &storedLetterType, "Data updated successfully")
}

func (l *LetterTypeUsecase) Update(letterType *model.LetterType, ctx *fiber.Ctx, id int) error {

	updateLetterType := &model.LetterType{
		TypeName: letterType.TypeName,
		Code:     letterType.Code,
	}

	updatedLetterType, err := letterTypeRepository.Update(updateLetterType, id)

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	return entity.Success(ctx, &updatedLetterType, "Data updated successfully")
}

func (l *LetterTypeUsecase) Delete(ctx *fiber.Ctx, id int) error {
	village, err := letterTypeRepository.Delete(id)

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	return entity.Success(ctx, &village, "Data deleted successfully")
}
