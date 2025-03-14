package rt_profile

import (
	"github.com/abiyyu03/siruta/entity"
	"github.com/abiyyu03/siruta/entity/constant"
	"github.com/abiyyu03/siruta/entity/model"
	"github.com/abiyyu03/siruta/entity/request"
	"github.com/abiyyu03/siruta/repository/rt_profile"
	"github.com/abiyyu03/siruta/usecase/email"
	"github.com/abiyyu03/siruta/usecase/referal_code"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type RTProfileRegisterUsecase struct{}

var rtProfileRegisterRepository *rt_profile.RTProfileRegisterRepository
var rtNotification *email.EmailRegistrationUsecase
var regTokenUsecase *referal_code.RegistrationTokenUsecase

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

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	return entity.Success(ctx, newRTProfile, "RT Profile Registered successfully")
}

func (r *RTProfileRegisterUsecase) Approve(emailDestination string, rtProfileId string, ctx *fiber.Ctx) error {
	token, err := regTokenUsecase.CreateToken(rtProfileId)

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	err = rtProfileRegisterRepository.ApproveRegistrant(rtProfileId)

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	err = rtNotification.RtNotification(emailDestination, token)

	if err != nil {
		return err
	}

	return entity.Success(ctx, nil, "RT Profile approved successfully")
}

func (r *RTProfileRegisterUsecase) RegisterUserRt(userRt *request.LeaderRegisterRequest, ctx *fiber.Ctx, token string) error {
	//token verif
	userId, _ := uuid.NewV7()
	leaderId, _ := uuid.NewV7()

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(userRt.Password),
		14,
	)

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	rtProfileId, err := regTokenUsecase.DecodeToken(token)

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	newUser := &model.User{
		ID:       userId.String(),
		RoleID:   constant.ROLE_RT,
		Email:    userRt.Email,
		Password: string(hashedPassword),
	}

	newLead := &model.RTLeader{
		ID:          leaderId.String(),
		Fullname:    userRt.Fullname,
		NikNumber:   userRt.NikNumber,
		KKNumber:    userRt.KKNumber,
		FullAddress: userRt.FullAddress,
		UserId:      newUser.ID,
		RTProfileId: rtProfileId,
	}

	err = rtProfileRegisterRepository.RegisterUserRt(newLead, newUser, constant.ROLE_RT, token)

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	return entity.Success(ctx, nil, "User Registered successfully")
}
