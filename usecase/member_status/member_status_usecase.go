package member_status

import (
	"github.com/abiyyu03/siruta/entity"
	"github.com/abiyyu03/siruta/entity/constant"
	"github.com/abiyyu03/siruta/entity/model"
	"github.com/abiyyu03/siruta/repository/member_status"
	"github.com/gofiber/fiber/v2"
)

type MemberStatusUsecase struct{}

var memberStatusRepository *member_status.MemberStatusRepository

type MemberStatusUsecaseInterface interface {
	Fetch(ctx *fiber.Ctx) error
	FetchById(ctx *fiber.Ctx, id int) error
	Store(memberStatus *model.MemberStatus, ctx *fiber.Ctx) error
	Update(memberStatus *model.MemberStatus, ctx *fiber.Ctx, id int) error
	Delete(ctx *fiber.Ctx, id int) error
}

func (m *MemberStatusUsecase) Fetch(ctx *fiber.Ctx) error {
	memberStatus, err := memberStatusRepository.Fetch()

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	return entity.Success(ctx, &memberStatus, "Data fetched successfully")
}

func (m *MemberStatusUsecase) FetchById(ctx *fiber.Ctx, id int) error {
	memberStatus, err := memberStatusRepository.FetchById(id)

	if memberStatus == nil {
		return entity.Error(ctx, fiber.StatusNotFound, constant.Errors["NotFound"].Message, constant.Errors["NotFound"].Clue)
	}

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	return entity.Success(ctx, &memberStatus, "Data fetched successfully")
}

func (m *MemberStatusUsecase) Store(memberStatus *model.MemberStatus, ctx *fiber.Ctx) error {
	createdStatus := &model.MemberStatus{
		Status: memberStatus.Status,
	}

	storedMemberStatus, err := memberStatusRepository.Store(createdStatus)

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	return entity.Success(ctx, &storedMemberStatus, "Data updated successfully")
}

func (m *MemberStatusUsecase) Update(memberStatus *model.MemberStatus, ctx *fiber.Ctx, id int) error {

	updateMemberStatus := &model.MemberStatus{
		Status: memberStatus.Status,
	}

	updatedMemberStatus, err := memberStatusRepository.Update(updateMemberStatus, id)

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	return entity.Success(ctx, &updatedMemberStatus, "Data updated successfully")
}

func (m *MemberStatusUsecase) Delete(ctx *fiber.Ctx, id int) error {
	village, err := memberStatusRepository.Delete(id)

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	return entity.Success(ctx, &village, "Data deleted successfully")
}
