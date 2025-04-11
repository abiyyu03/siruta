package letter

import (
	"github.com/abiyyu03/siruta/entity"
	"github.com/abiyyu03/siruta/entity/constant"
	"github.com/abiyyu03/siruta/entity/model"
	"github.com/abiyyu03/siruta/repository/letter"
	"github.com/abiyyu03/siruta/repository/member"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type LetterReqUsecase struct{}

var letterReqRepository *letter.LetterReqRepository
var memberRepository *member.MemberRepository

func (l *LetterReqUsecase) StoreOutcomingLetter(ctx *fiber.Ctx, memberData *model.Member, outcommingLetter *model.OutcomingLetter, memberStatus string, birthDate string, nik string) error {
	letterId, _ := uuid.NewV7()
	var err error

	memberResult, _ := memberRepository.FetchByNikAndBirtDate(nik, birthDate)

	newOutcomingLetter := &model.OutcomingLetter{
		ID: letterId.String(),
		// LetterNumber: outcommingLetter.LetterNumber,
		Date:         outcommingLetter.Date,
		LetterTypeId: outcommingLetter.LetterTypeId,
		RTProfileId:  memberResult.RTProfileId,
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
			return entity.Error(ctx, fiber.StatusNotFound, constant.Errors["UserNotFound"].Message, constant.Errors["UserNotFound"].Clue)
		}

		_, err = letterReqRepository.StoreOutcomingLetter(newOutcomingLetter)
	}

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	return entity.Success(ctx, nil, "Request letter created successfully")
}

func (l *LetterReqUsecase) UpdateApprovalStatusByRT(ctx *fiber.Ctx, id string) error {
	approvalStatus, err := letterReqRepository.UpdateApprovalStatusByRT(id)

	if !approvalStatus {
		return entity.Error(ctx, fiber.StatusBadRequest, constant.Errors["LetterRejected"].Message, constant.Errors["LetterRejected"].Clue)
	}

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	return entity.Success(ctx, nil, "Approval status updated successfully")
}
