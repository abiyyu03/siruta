package usecase

import (
	"github.com/abiyyu03/siruta/entity"
	"github.com/abiyyu03/siruta/entity/constant"
	"github.com/abiyyu03/siruta/entity/model"
	"github.com/abiyyu03/siruta/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MemberUsecase struct {
	db *gorm.DB
}

var memberRepository = new(repository.MemberRepository)

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

func (m *MemberUsecase) FetchById(ctx *fiber.Ctx, id string) error {
	members, err := memberRepository.FetchById(id)

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	if members == nil {
		return entity.Error(ctx, fiber.StatusNotFound, constant.Errors["NotFound"].Message, constant.Errors["NotFound"].Clue)
	}

	return entity.Success(ctx, &members, "Data fetched successfully")
}

func (m *MemberUsecase) Store(ctx *fiber.Ctx, memberData *model.Member) error {
	id := uuid.New()

	createdMember := &model.Member{
		ID:             id.String(),
		ReligionId:     memberData.ReligionId,
		UserId:         memberData.UserId,
		Occupation:     memberData.Occupation,
		Fullname:       memberData.Fullname,
		NikNumber:      memberData.NikNumber,
		KKNumber:       memberData.KKNumber,
		BornPlace:      memberData.BornPlace,
		BirthDate:      memberData.BirthDate,
		Gender:         memberData.Gender,
		HomeAddress:    memberData.HomeAddress,
		MaritalStatus:  memberData.MaritalStatus,
		MemberStatusId: memberData.MemberStatusId,
	}

	create, err := memberRepository.Store(m.db, createdMember)

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	return entity.Success(ctx, create, "Data created successfully")
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
