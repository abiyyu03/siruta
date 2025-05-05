package letter

import (
	"github.com/abiyyu03/siruta/entity"
	"github.com/abiyyu03/siruta/entity/constant"
	"github.com/abiyyu03/siruta/repository/letter"
	"github.com/gofiber/fiber/v2"
)

type OutcomingLetterUsecase struct {
	outcomingLetterRepository letter.OutcomingLetterRepository
}

type OutcomingLetterUsecaseInterface interface {
	Fetch(ctx *fiber.Ctx) error
	FetchById(ctx *fiber.Ctx, id string) error
	FetchByRTProfileId(ctx *fiber.Ctx, rtProfileId string) error
	FetchPreview(ctx *fiber.Ctx, id string) error
	Delete(ctx *fiber.Ctx, id string) error
}

func (o *OutcomingLetterUsecase) Fetch(ctx *fiber.Ctx) error {
	letters, err := o.outcomingLetterRepository.Fetch()

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	return entity.Success(ctx, &letters, "Data fetched successfully")
}

func (o *OutcomingLetterUsecase) FetchById(ctx *fiber.Ctx, id string) error {
	letter, err := o.outcomingLetterRepository.FetchById(id)

	if letter == nil {
		return entity.Error(ctx, fiber.StatusNotFound, constant.Errors["NotFound"].Message, constant.Errors["NotFound"].Clue)
	}

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	return entity.Success(ctx, &letter, "Data fetched successfully")
}

func (o *OutcomingLetterUsecase) FetchByRTProfileId(ctx *fiber.Ctx, rtProfileId string) error {
	letters, err := o.outcomingLetterRepository.FetchByRTProfileId(rtProfileId)

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	return entity.Success(ctx, &letters, "Data fetched successfully")
}

func (o *OutcomingLetterUsecase) FetchPreview(ctx *fiber.Ctx, id string) error {
	letter, err := o.outcomingLetterRepository.FetchPreview(id)

	if letter == nil {
		return entity.Error(ctx, fiber.StatusNotFound, constant.Errors["NotFound"].Message, constant.Errors["NotFound"].Clue)
	}

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	previewLetter := &entity.OutcomeLetterResponse{
		ID:           letter.ID,
		LetterNumber: letter.LetterNumber,
		Date:         letter.Date,
		Member: entity.MemberPreview{
			ID:             letter.Member.ID,
			Fullname:       letter.Member.Fullname,
			NIKNumber:      letter.Member.NikNumber,
			KKNumber:       letter.Member.KKNumber,
			BornPlace:      letter.Member.BornPlace,
			BirthDate:      letter.Member.BirthDate,
			Gender:         letter.Member.Gender,
			HomeAddress:    letter.Member.HomeAddress,
			MaritalStatus:  letter.Member.MaritalStatus,
			ReligionID:     letter.Member.ReligionId,
			MemberStatusID: letter.Member.MemberStatusId,
			Occupation:     letter.Member.Occupation,
			Status:         letter.Member.Status,
		},
		LetterType: entity.LetterTypePreview{
			ID:       letter.LetterType.ID,
			TypeName: letter.LetterType.TypeName,
			Code:     letter.LetterType.Code,
		},
		RTProfile: entity.RTProfilePreview{
			ID:       letter.RTProfile.ID,
			RTNumber: letter.RTProfile.RTNumber,
			RTLogo:   letter.RTProfile.RTLogo,
		},
		IsRTApproved: letter.IsRTApproved,
		Description:  letter.Description,
	}

	return entity.Success(ctx, &previewLetter, "Data fetched successfully")
}

func (o *OutcomingLetterUsecase) Delete(ctx *fiber.Ctx, id string) error {
	err := o.outcomingLetterRepository.Delete(id)

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	return entity.Success(ctx, nil, "Data deleted successfully")
}
