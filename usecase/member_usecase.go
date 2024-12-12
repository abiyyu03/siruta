package usecase

import (
	"github.com/abiyyu03/siruta/entity"
	"github.com/abiyyu03/siruta/entity/model"
	"github.com/abiyyu03/siruta/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type MemberUsecase struct{}

var memberRepository = new(repository.MemberRepository)

func (m *MemberUsecase) Fetch(ctx *fiber.Ctx) error {
	members, err := memberRepository.Fetch(ctx)

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, "Error fetching data")
	}

	if members == nil {
		return nil
	}

	return entity.Success(ctx, &members, "Data fetched successfully")
}

func (m *MemberUsecase) FetchById(ctx *fiber.Ctx, id string) error {
	members, err := memberRepository.Fetch(ctx)

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, "Error fetching data")
	}

	if members == nil {
		return entity.Error(ctx, fiber.StatusNotFound, fiber.ErrNotFound.Message)
	}

	return entity.Success(ctx, &members, "Data fetched successfully")
}

func (m *MemberUsecase) Store(ctx *fiber.Ctx, memberData *model.Member) error {
	id := uuid.New()

	createdMember := &model.Member{
		ID:            id.String(),
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

	create, err := memberRepository.Store(createdMember)

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, fiber.ErrInternalServerError.Message)
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
		return entity.Error(ctx, fiber.StatusInternalServerError, fiber.ErrInternalServerError.Message)
	}

	if update == nil {
		return entity.Error(ctx, fiber.StatusNotFound, fiber.ErrNotFound.Message)
	}

	return entity.Success(ctx, fiber.StatusOK, "Data updated successfully")
}
