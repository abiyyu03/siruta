package referal_code

import (
	"github.com/abiyyu03/siruta/entity"
	"github.com/abiyyu03/siruta/entity/constant"
	"github.com/abiyyu03/siruta/repository/referal_code"
	"github.com/gofiber/fiber/v2"
)

type ReferalCodeUsecase struct{}

var referalCodeRepository *referal_code.ReferalCodeRepository

type IdProfileResponse struct {
	ID string `json:"id"`
}

func (r *ReferalCodeUsecase) Fetch(ctx *fiber.Ctx) error {
	referals, err := referalCodeRepository.Fetch()

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	return entity.Success(ctx, &referals, "Data fetched successfully")
}

func (r *ReferalCodeUsecase) FetchById(ctx *fiber.Ctx, id string) error {
	referal, err := referalCodeRepository.FetchById(id)

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	if referal == nil {
		return entity.Error(ctx, fiber.StatusNotFound, constant.Errors["NotFound"].Message, constant.Errors["NotFound"].Clue)
	}

	return entity.Success(ctx, &referal, "Data fetched successfully")
}

func (r *ReferalCodeUsecase) Validate(ctx *fiber.Ctx, code string) error {
	rwProfileId, isValid, err := referalCodeRepository.Validate(code)

	if !isValid {
		return entity.Error(ctx, fiber.StatusBadRequest, constant.Errors["InvalidReferalCode"].Message, constant.Errors["InvalidReferalCode"].Clue)
	}

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	restructureRwProfile := &IdProfileResponse{
		ID: rwProfileId,
	}

	return entity.Success(ctx, &restructureRwProfile, "Referal code is valid")
}
