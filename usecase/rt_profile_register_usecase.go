package usecase

import (
	"log"

	"github.com/abiyyu03/siruta/entity"
	"github.com/abiyyu03/siruta/entity/constant"
	"github.com/abiyyu03/siruta/entity/model"
	"github.com/abiyyu03/siruta/entity/request"
	"github.com/abiyyu03/siruta/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type RTProfileRegisterUsecase struct{}

var rtProfileRegisterRepository = new(repository.RTProfileRegisterRepository)

func (r *RTProfileRegisterUsecase) RegisterRTProfile(rtProfile *request.RTProfileRegisterRequest, ctx *fiber.Ctx) error {
	id, _ := uuid.NewV7()

	newRTProfile := &model.RTProfile{
		ID:          id.String(),
		RTNumber:    rtProfile.RTNumber,
		Latitude:    rtProfile.Latitude,
		Longitude:   rtProfile.Longitude,
		RTEmail:     rtProfile.RTEmail,
		MobilePhone: rtProfile.MobilePhone,
		FullAddress: rtProfile.FullAddress,
	}

	registeredUser, err := rtProfileRegisterRepository.Register(newRTProfile, rtProfile.ReferalCode)

	log.Print(registeredUser)

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	return entity.Success(ctx, newRTProfile, "RT Profile Registered successfully")
}

func (r *RTProfileRegisterUsecase) Approve(emailDestination string, rtProfileId string, ctx *fiber.Ctx) error {
	err := rtProfileRegisterRepository.ApproveRegistrant(rtProfileId)

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	return entity.Success(ctx, nil, "RT Profile approved successfully")
}

func (r *RTProfileRegisterUsecase) RegisterUserRt(userRt *request.RegisterRTRequest, ctx *fiber.Ctx, token string) error {
	//token verif
	userId, _ := uuid.NewV7()
	memberId, _ := uuid.NewV7()

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(userRt.Password),
		14,
	)

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	newUser := &model.User{
		ID:       userId.String(),
		Username: userRt.Username,
		RoleID:   constant.ROLE_RT,
		Email:    userRt.Email,
		Password: string(hashedPassword),
	}

	newMember := &model.Member{
		ID:             memberId.String(),
		Fullname:       userRt.Fullname,
		NikNumber:      &userRt.NikNumber,
		KKNumber:       &userRt.KKNumber,
		BornPlace:      userRt.BornPlace,
		BirthDate:      userRt.BirthDate,
		Gender:         userRt.Gender,
		HomeAddress:    &userRt.HomeAddress,
		MaritalStatus:  &userRt.MaritalStatus,
		ReligionId:     userRt.ReligionId,
		MemberStatusId: userRt.MemberStatusId,
		UserId:         &newUser.ID,
		Occupation:     &userRt.Occupation,
		Status:         userRt.Status,
	}

	err = rtProfileRegisterRepository.RegisterUserRt(newMember, newUser, constant.ROLE_RT, token)

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	return entity.Success(ctx, nil, "User Registered successfully")
}
