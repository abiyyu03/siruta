package usecase

import (
	"github.com/abiyyu03/siruta/entity"
	"github.com/abiyyu03/siruta/entity/model"
	"github.com/abiyyu03/siruta/repository"
	"github.com/gofiber/fiber/v2"
)

type IncomingLetterUsecase struct{}

var IncomingLetterRepository = new(repository.IncomingLetterRepository)

func (i *IncomingLetterUsecase) Fetch(ctx *fiber.Ctx) error {
	incomingLetters, err := IncomingLetterRepository.Fetch()

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, "Server error")
	}

	if incomingLetters == nil {
		return entity.Error(ctx, fiber.StatusNotFound, fiber.ErrNotFound.Message)
	}

	return entity.Success(ctx, &incomingLetters, "Data fetched successfully")
}

func (i *IncomingLetterUsecase) FetchById(ctx *fiber.Ctx, id int) error {
	incomingLetter, err := IncomingLetterRepository.FetchById(id)

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, "Server error")
	}

	if incomingLetter == nil {
		return nil
	}

	return entity.Success(ctx, &incomingLetter, "Data fetched successfully")
}

func (i *IncomingLetterUsecase) Delete(ctx *fiber.Ctx, id int) error {
	var incomingLetter *model.IncomingLetter
	incomingLetter, err := IncomingLetterRepository.Delete(incomingLetter, id)

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, "Server error")
	}

	if incomingLetter == nil {
		return entity.Error(ctx, fiber.StatusNotFound, fiber.ErrNotFound.Message)
	}

	return entity.Success(ctx, &incomingLetter, "Data deleted successfully")
}

func (i *IncomingLetterUsecase) Store(incomingLetter *model.IncomingLetter, ctx *fiber.Ctx) error {
	createdIncomingLetter := &model.IncomingLetter{
		Title:        incomingLetter.Title,
		LetterDate:   incomingLetter.LetterDate,
		OriginLetter: incomingLetter.OriginLetter,
		RTProfileId:  incomingLetter.RTProfileId,
	}

	newIncomingLetter, err := IncomingLetterRepository.Store(createdIncomingLetter)

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, "Server error")
	}

	return entity.Success(ctx, &newIncomingLetter, "Data stored successfully")
}

func (i *IncomingLetterUsecase) Update(incomingLetter *model.IncomingLetter, ctx *fiber.Ctx, id int) error {
	updatedIncomingLetter := &model.IncomingLetter{
		Title:        incomingLetter.Title,
		LetterDate:   incomingLetter.LetterDate,
		OriginLetter: incomingLetter.OriginLetter,
		RTProfileId:  incomingLetter.RTProfileId,
	}

	incomingLetter, err := IncomingLetterRepository.Update(updatedIncomingLetter, id)

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, fiber.ErrInternalServerError.Message)
	}

	if incomingLetter == nil {
		return entity.Error(ctx, fiber.StatusNotFound, fiber.ErrNotFound.Message)
	}

	return entity.Success(ctx, fiber.StatusOK, "Data updated successfully")
}