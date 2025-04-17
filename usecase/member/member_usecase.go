package member

import (
	"github.com/abiyyu03/siruta/entity"
	"github.com/abiyyu03/siruta/entity/constant"
	"github.com/abiyyu03/siruta/entity/model"
	"github.com/abiyyu03/siruta/repository/member"
	"github.com/gofiber/fiber/v2"
)

type MemberUsecase struct{}

var memberRepository *member.MemberRepository

func (m *MemberUsecase) Fetch(ctx *fiber.Ctx) error {
	members, err := memberRepository.Fetch()

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	if members == nil {
		return nil
	}

	return entity.Success(ctx, &members, "Data fetched successfully")
}

// func (m *MemberUsecase) BatchStore(ctx fiber.Ctx) error {

// }

func (m *MemberUsecase) FetchById(ctx *fiber.Ctx, id string) error {
	members, err := memberRepository.FetchById(id)

	if members == nil {
		return entity.Error(ctx, fiber.StatusNotFound, constant.Errors["NotFound"].Message, constant.Errors["NotFound"].Clue)
	}

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	return entity.Success(ctx, &members, "Data fetched successfully")
}

func (m *MemberUsecase) FetchByRTProfileId(ctx *fiber.Ctx, rtProfileId string) error {
	members, err := memberRepository.FetchByRTProfileId(rtProfileId)
	if members == nil {
		return entity.Error(ctx, fiber.StatusNotFound, constant.Errors["NotFound"].Message, constant.Errors["NotFound"].Clue)
	}

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	return entity.Success(ctx, &members, "Data fetched successfully")
}

func (m *MemberUsecase) Update(ctx *fiber.Ctx, id string, memberData *model.Member) error {
	updatedMember := &model.Member{
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
	}

	update, err := memberRepository.Update(updatedMember, id)

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	if update == nil {
		return entity.Error(ctx, fiber.StatusNotFound, constant.Errors["NotFound"].Message, constant.Errors["NotFound"].Clue)
	}

	return entity.Success(ctx, fiber.StatusOK, "Data updated successfully")
}
