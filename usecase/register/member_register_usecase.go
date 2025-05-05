package register

import (
	"github.com/abiyyu03/siruta/entity"
	"github.com/abiyyu03/siruta/entity/constant"
	"github.com/abiyyu03/siruta/entity/model"
	"github.com/abiyyu03/siruta/entity/request"
	"github.com/abiyyu03/siruta/repository/register"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type MemberRegisterUsecase struct{}

var memberRegisterRepository *register.MemberRegisterRepository

type MemberRegisterUsecaseInterface interface {
	RegisterMember(ctx *fiber.Ctx, userMember *request.MemberRegisterRequest, profileId string) error
}

func (m *MemberRegisterUsecase) RegisterMember(ctx *fiber.Ctx, userMember *request.MemberRegisterRequest, profileId string) error {
	userId, _ := uuid.NewV7()
	memberId, _ := uuid.NewV7()

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(userMember.Password),
		14,
	)

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	newUser := &model.User{
		ID:       userId.String(),
		RoleID:   constant.ROLE_MEMBER,
		Email:    userMember.Email,
		Password: string(hashedPassword),
	}

	newMember := &model.Member{
		ID:             memberId.String(),
		Fullname:       userMember.Fullname,
		NikNumber:      &userMember.NikNumber,
		KKNumber:       &userMember.KKNumber,
		BornPlace:      userMember.BornPlace,
		BirthDate:      userMember.BirthDate,
		Gender:         userMember.Gender,
		HomeAddress:    &userMember.HomeAddress,
		MaritalStatus:  &userMember.MaritalStatus,
		ReligionId:     userMember.ReligionId,
		MemberStatusId: userMember.MemberStatusId,
		UserId:         newUser.ID,
		Occupation:     &userMember.Occupation,
		Status:         userMember.Status,
		RTProfileId:    profileId,
	}

	err = memberRegisterRepository.RegisterMember(newMember, newUser)

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	return entity.Success(ctx, nil, "User Registered successfully")
}
