package letter

import (
	"github.com/abiyyu03/siruta/entity"
	"github.com/abiyyu03/siruta/entity/constant"
	"github.com/abiyyu03/siruta/entity/model"
	"github.com/abiyyu03/siruta/repository/letter"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type LetterReqUsecase struct{}

var letterReqRepository *letter.LetterReqRepository

func (l *LetterReqUsecase) StoreOutcomingLetter(ctx *fiber.Ctx, memberData *model.Member, outcommingLetter *model.OutcomingLetter, memberStatus string, birthDate string, nik string) error {
	letterId, _ := uuid.NewV7()
	var err error

	newOutcomingLetter := &model.OutcomingLetter{
		ID: letterId.String(),
		// LetterNumber: outcommingLetter.LetterNumber,
		Date:         outcommingLetter.Date,
		LetterTypeId: outcommingLetter.LetterTypeId,
		RTProfileId:  outcommingLetter.RTProfileId,
		MemberId:     outcommingLetter.MemberId,
		IsRTApproved: outcommingLetter.IsRTApproved,
		Description:  outcommingLetter.Description,
	}

	if memberStatus == "guest" {
		memberId, _ := uuid.NewV7()

		createdGuestMember := &model.Member{
			ID:            memberId.String(),
			ReligionId:    memberData.ReligionId,
			UserId:        memberData.UserId,
			Occupation:    memberData.Occupation,
			Fullname:      memberData.Fullname,
			NikNumber:     memberData.NikNumber,
			KKNumber:      memberData.KKNumber,
			BornPlace:     memberData.BornPlace,
			BirthDate:     memberData.BirthDate,
			Gender:        memberData.Gender,
			HomeAddress:   memberData.HomeAddress,
			MaritalStatus: memberData.MaritalStatus,
			RTProfileId:   memberData.RTProfileId,
		}

		_, err = letterReqRepository.StoreOutcomingLetterWithGuest(newOutcomingLetter, createdGuestMember)

	} else {
		isMemberExist, _ := letterReqRepository.CheckMemberResidentExist(birthDate, nik)

		if !isMemberExist {
			return entity.Error(ctx, fiber.StatusNotFound, constant.Errors["NotFound"].Message, constant.Errors["NotFound"].Clue)
		}

		_, err = letterReqRepository.StoreOutcomingLetter(newOutcomingLetter)
	}

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	return entity.Success(ctx, newOutcomingLetter, "Data created successfully")
}

func (l *LetterReqUsecase) UpdateApprovalStatusByRT(outcomingLetter *model.OutcomingLetter, id string, ctx *fiber.Ctx) error {
	approvalStatus, err := letterReqRepository.UpdateApprovalStatusByRT(outcomingLetter, id)

	if !approvalStatus {
		return entity.Error(ctx, fiber.StatusBadRequest, constant.Errors["LetterRejected"].Message, constant.Errors["LetterRejected"].Clue)
	}

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	return entity.Success(ctx, nil, "Approval status updated successfully")
}
