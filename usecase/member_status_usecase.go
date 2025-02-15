package usecase

import (
	"github.com/abiyyu03/siruta/entity"
	"github.com/abiyyu03/siruta/entity/model"
	"github.com/abiyyu03/siruta/repository"
	"github.com/gofiber/fiber/v2"
)

type MemberStatusUsecase struct{}

var memberStatusRepository = new(repository.MemberStatusRepository)

func (v *MemberStatusUsecase) Fetch(ctx *fiber.Ctx) error {
	memberStatus, err := memberStatusRepository.Fetch()

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, "Server error")
	}

	return entity.Success(ctx, &memberStatus, "Data fetched successfully")
}

func (v *MemberStatusUsecase) FetchById(ctx *fiber.Ctx, id int) error {
	memberStatus, err := memberStatusRepository.FetchById(id)

	if memberStatus == nil {
		return entity.Error(ctx, fiber.StatusNotFound, "Member Status not found")
	}

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, fiber.ErrInternalServerError.Message)
	}

	return entity.Success(ctx, &memberStatus, "Data fetched successfully")
}

func (v *MemberStatusUsecase) Store(memberStatus *model.MemberStatus, ctx *fiber.Ctx) error {
	createdStatus := &model.MemberStatus{
		Status: memberStatus.Status,
	}

	storedMemberStatus, err := memberStatusRepository.Store(createdStatus)

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, err.Error())
	}

	return entity.Success(ctx, &storedMemberStatus, "Data updated successfully")
}

func (v *MemberStatusUsecase) Update(memberStatus *model.MemberStatus, ctx *fiber.Ctx, id int) error {

	updateMemberStatus := &model.MemberStatus{
		Status: memberStatus.Status,
	}

	updatedMemberStatus, err := memberStatusRepository.Update(updateMemberStatus, id)

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, fiber.ErrInternalServerError.Message)
	}

	return entity.Success(ctx, &updatedMemberStatus, "Data updated successfully")
}

func (v *MemberStatusUsecase) Delete(ctx *fiber.Ctx, id int) error {
	village, err := memberStatusRepository.Delete(id)

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, "Server error")
	}

	return entity.Success(ctx, &village, "Data deleted successfully")
}
