package referal_code

import (
	"github.com/abiyyu03/siruta/entity"
	"github.com/abiyyu03/siruta/entity/constant"
	"github.com/abiyyu03/siruta/repository/referal_code"
	"github.com/gofiber/fiber/v2"
)

type ReferalCodeUsecase struct {
	referalCodeRepository referal_code.ReferalCodeRepository
}

type ReferalCodeUsecaseInterface interface {
	Fetch(ctx *fiber.Ctx) error
	FetchById(ctx *fiber.Ctx, id string) error
	FetchByRTProfileId(ctx *fiber.Ctx, rtProfileId string) error
	RegenerateReferalCode(ctx *fiber.Ctx, profileId string, code string) error
	Validate(ctx *fiber.Ctx, code string) (error, string)
	Delete(ctx *fiber.Ctx, id string) error
}

type IdProfileResponse struct {
	ID string `json:"id"`
}

func (r *ReferalCodeUsecase) Fetch(ctx *fiber.Ctx) error {
	referals, err := r.referalCodeRepository.Fetch()

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	return entity.Success(ctx, &referals, "Data fetched successfully")
}

func (r *ReferalCodeUsecase) FetchById(ctx *fiber.Ctx, id string) error {
	referal, err := r.referalCodeRepository.FetchById(id)

	if referal == nil {
		return entity.Error(ctx, fiber.StatusNotFound, constant.Errors["NotFound"].Message, constant.Errors["NotFound"].Clue)
	}

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	return entity.Success(ctx, &referal, "Data fetched successfully")
}

func (r *ReferalCodeUsecase) FetchByRTProfileId(ctx *fiber.Ctx, rtProfileId string) error {
	referals, err := r.referalCodeRepository.FetchByRTProfileId(rtProfileId)

	if referals == nil {
		return entity.Error(ctx, fiber.StatusNotFound, constant.Errors["NotFound"].Message, constant.Errors["NotFound"].Clue)
	}

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	return entity.Success(ctx, &referals, "Data fetched successfully")
}

func (r *ReferalCodeUsecase) RegenerateReferalCode(ctx *fiber.Ctx, profileId string, code string) error {
	regeneratedCode, err := r.referalCodeRepository.RegenerateReferalCode(profileId, code)

	if regeneratedCode == "" {
		return entity.Error(ctx, fiber.StatusBadRequest, constant.Errors["InvalidReferalCode"].Message, constant.Errors["InvalidReferalCode"].Clue)
	}

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	return entity.Success(ctx, map[string]string{"code": regeneratedCode}, "Referal code has been regenerated successfully")
}

func (r *ReferalCodeUsecase) Validate(ctx *fiber.Ctx, code string) (error, string) {
	profileId, isValid, err := r.referalCodeRepository.Validate(code)

	if !isValid {
		return entity.Error(ctx, fiber.StatusBadRequest, constant.Errors["InvalidReferalCode"].Message, constant.Errors["InvalidReferalCode"].Clue), ""
	}

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue), ""
	}

	restructureRwProfile := &IdProfileResponse{
		ID: profileId,
	}

	return entity.Success(ctx, &restructureRwProfile, "Referal code is valid"), profileId
}

func (r *ReferalCodeUsecase) Delete(ctx *fiber.Ctx, id int) error {
	err := r.referalCodeRepository.Delete(id)

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	return entity.Success(ctx, nil, "Referal data deleted successfully")
}
