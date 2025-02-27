package usecase

import (
	"github.com/abiyyu03/siruta/entity"
	"github.com/abiyyu03/siruta/entity/constant"
	"github.com/abiyyu03/siruta/entity/model"
	"github.com/abiyyu03/siruta/entity/request"
	"github.com/abiyyu03/siruta/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type RWProfileRegisterUsecase struct{}

var rwProfileRegisterRepository = new(repository.RWProfileRegisterRepository)

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
func (r *RWProfileRegisterUsecase) RegisterUserRw(register *request.RegisterRWRequest, ctx *fiber.Ctx, token string) error {
	//token verif
	userId, _ := uuid.NewV7()
	memberId, _ := uuid.NewV7()

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(register.Password),
		14,
	)
	if err != nil {
		return err
	}

	user := &model.User{
		ID:       userId.String(),
		RoleID:   uint(4),
		Email:    register.Email,
		Password: string(hashedPassword),
	}

	member := &model.Member{
		ID:             memberId.String(),
		Fullname:       register.Fullname,
		NikNumber:      &register.NikNumber,
		KKNumber:       &register.KKNumber,
		BornPlace:      register.BornPlace,
		BirthDate:      register.BirthDate,
		Gender:         register.Gender,
		HomeAddress:    &register.HomeAddress,
		MaritalStatus:  &register.MaritalStatus,
		ReligionId:     register.ReligionId,
		MemberStatusId: register.MemberStatusId,
		UserId:         &user.ID,
		Occupation:     &register.Occupation,
		Status:         register.Status,
	}
	err = rwProfileRegisterRepository.RegisterUserRW(member, user, 3, token)

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
