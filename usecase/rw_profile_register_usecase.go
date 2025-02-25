package usecase

import (
	"log"

	"github.com/abiyyu03/siruta/entity"
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
		return entity.Error(ctx, fiber.StatusInternalServerError, err.Error())
	}

	return entity.Success(ctx, &rwProfileRegistration, "RW Profile Registered successfully")
}
func (r *RWProfileRegisterUsecase) RegisterUserRW(register *request.RegisterRWRequest, ctx *fiber.Ctx, token string) error {
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
		Username: register.Username,
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

	log.Print("member status: ", register.MemberStatusId)

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, err.Error())
	}

	return entity.Success(ctx, nil, "User Registered successfully")
}

func (r *RWProfileRegisterUsecase) Approve(emailDestination string, rwProfileId string, ctx *fiber.Ctx) error {
	if err := rwProfileRegisterRepository.ApproveRegistrant(rwProfileId); err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, err.Error())
	}

	return entity.Success(ctx, nil, "RW Profile approved successfully")
}
