package usecase

import (
	"github.com/abiyyu03/siruta/entity"
	"github.com/abiyyu03/siruta/entity/model"
	"github.com/abiyyu03/siruta/repository"
	"github.com/abiyyu03/siruta/repository/email"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type RTProfileRegistrationUsecase struct{}

var rtProfileRegistrationRepository = new(repository.RTProfileRegistrationRepository)

var emailRegistrationRepository = new(email.EmailUserRegistrationRepository)

func (r *RTProfileRegistrationUsecase) Register(ctx *fiber.Ctx, rtProfile *model.RTProfile, rwProfileId string) error {
	id, _ := uuid.NewV7()

	newRTProfile := &model.RTProfile{
		ID:          id.String(),
		RtNumber:    rtProfile.RtNumber,
		Latitude:    rtProfile.Latitude,
		Longitude:   rtProfile.Longitude,
		RtEmail:     rtProfile.RtEmail,
		MobilePhone: rtProfile.MobilePhone,
		RWProfileId: rwProfileId,
	}

	newRTProfile, err := rtProfileRegistrationRepository.Register(newRTProfile, rwProfileId)

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, err.Error())
	}

	return entity.Success(ctx, newRTProfile, "RT Profile registered successfully")
}

func (r *RTProfileRegistrationUsecase) ApproveRegistrant(ctx *fiber.Ctx, rtProfileId string) error {
	rtProfileRegistration, err := rtProfileRegistrationRepository.ApproveRegistrant(rtProfileId)

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, fiber.ErrInternalServerError.Message)
	}

	err = emailRegistrationRepository.Send(rtProfileRegistration.RtEmail)

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, fiber.ErrInternalServerError.Message)
	}

	return entity.Success(ctx, &rtProfileRegistration, "Your Registration has been approved, please check your email")
}
