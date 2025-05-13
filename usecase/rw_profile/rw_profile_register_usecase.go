package rw_profile

import (
	"github.com/abiyyu03/siruta/entity"
	"github.com/abiyyu03/siruta/entity/constant"
	"github.com/abiyyu03/siruta/entity/model"
	"github.com/abiyyu03/siruta/repository/rw_profile"
	"github.com/abiyyu03/siruta/usecase/referal_code"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type RWProfileRegisterUsecase struct {
	rwProfileRegisterRepository rw_profile.RWProfileRegisterRepository
	regTokenUsecase             referal_code.RegistrationTokenUsecase
}

var rwProfileRegisterRepository *rw_profile.RWProfileRegisterRepository
var regTokenUsecase *referal_code.RegistrationTokenUsecase

type RWProfileRegisterUsecaseInterface interface {
	RegisterProfileRW(rwProfile *model.RWProfile, ctx *fiber.Ctx) error
	RegisterUserRw(register *entity.LeaderRegisterRequest, ctx *fiber.Ctx, token string) error
	Approve(emailDestination string, rwProfileId string, ctx *fiber.Ctx) error
}

func (r *RWProfileRegisterUsecase) RegisterProfileRW(rwProfile *model.RWProfile, ctx *fiber.Ctx) error {
	id, _ := uuid.NewV7()

	newRWProfile := &model.RWProfile{
		ID:          id.String(),
		RWNumber:    rwProfile.RWNumber,
		VillageID:   rwProfile.VillageID,
		RwEmail:     rwProfile.RwEmail,
		MobilePhone: rwProfile.MobilePhone,
		FullAddress: rwProfile.FullAddress,
		RWLogo:      rwProfile.RWLogo,
		RegencyLogo: rwProfile.RegencyLogo,
	}

	rwProfileRegistration, err := rwProfileRegisterRepository.RegisterRWProfile(newRWProfile)

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	return entity.Success(ctx, &rwProfileRegistration, "RW Profile Registered successfully")
}
func (r *RWProfileRegisterUsecase) RegisterUserRw(register *entity.LeaderRegisterRequest, ctx *fiber.Ctx, token string) error {
	//token verif
	userId, _ := uuid.NewV7()
	leaderId, _ := uuid.NewV7()

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(register.Password),
		14,
	)
	if err != nil {
		return err
	}

	rwProfileId, err := regTokenUsecase.DecodeToken(token)

	if err != nil {
		return err
	}

	user := &model.User{
		ID:       userId.String(),
		RoleID:   uint(4),
		Email:    register.Email,
		Password: string(hashedPassword),
	}

	newLeader := &model.RWLeader{
		ID:          leaderId.String(),
		Fullname:    register.Fullname,
		NikNumber:   register.NikNumber,
		KKNumber:    register.KKNumber,
		FullAddress: register.FullAddress,
		UserId:      user.ID,
		RWProfileId: rwProfileId,
	}
	err = rwProfileRegisterRepository.RegisterUserRW(newLeader, user, 3, token)

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	return entity.Success(ctx, nil, "User Registered successfully")
}

func (r *RWProfileRegisterUsecase) Approve(emailDestination string, rwProfileId string, ctx *fiber.Ctx) error {
	err := rwProfileRegisterRepository.ApproveRegistrant(rwProfileId)

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	return entity.Success(ctx, nil, "RW Profile approved successfully")
}
