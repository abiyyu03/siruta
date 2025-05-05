package auth

import (
	"time"

	"github.com/abiyyu03/siruta/entity"
	"github.com/abiyyu03/siruta/entity/constant"
	"github.com/abiyyu03/siruta/entity/model"
	"github.com/abiyyu03/siruta/entity/request"
	"github.com/abiyyu03/siruta/helper"
	"github.com/abiyyu03/siruta/repository/auth"
	"github.com/abiyyu03/siruta/repository/user"
	"github.com/abiyyu03/siruta/usecase/email"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type ResetPasswordUsecase struct {
	resetPasswordRepository auth.ResetPasswordRepository
	sendEmail               email.ResetPasswordEmailUsecase
	userRepository          user.UserRepository
}

type ResetPasswordUsecaseInterface interface {
	SendForgotPasswordLink(ctx *fiber.Ctx, email string) error
	ResetPassword(ctx *fiber.Ctx, reset *request.ResetPassword, token string) error
}

func (r *ResetPasswordUsecase) SendForgotPasswordLink(ctx *fiber.Ctx, email string) error {
	randomToken, _ := helper.GenerateSecureToken(48)

	user, err := r.userRepository.FetchByEmail(email)

	if err != nil {
		return entity.Error(ctx, fiber.ErrInternalServerError.Code, constant.Errors["internalError"].Message, constant.Errors["internalError"].Clue)
	}

	createdToken := &model.ResetPassword{
		Token:     randomToken,
		UserID:    user.ID,
		ExpiredAt: time.Now(),
	}

	_, err = r.resetPasswordRepository.StoreToken(createdToken)

	if err != nil {
		return entity.Error(ctx, fiber.ErrInternalServerError.Code, constant.Errors["internalError"].Message, constant.Errors["internalError"].Clue)
	}

	err = r.sendEmail.ResetPasswordEmail(email, randomToken)

	if err != nil {
		return entity.Error(ctx, fiber.ErrInternalServerError.Code, constant.Errors["internalError"].Message, constant.Errors["internalError"].Clue)
	}

	return entity.Success(ctx, nil, "Link reset password sent successfully")
}

func (r *ResetPasswordUsecase) ResetPassword(ctx *fiber.Ctx, reset *request.ResetPassword, token string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(reset.Password),
		14,
	)

	if err != nil {
		return entity.Error(ctx, fiber.ErrInternalServerError.Code, constant.Errors["internalError"].Message, constant.Errors["internalError"].Clue)
	}

	err = r.resetPasswordRepository.ResetPassword(string(hashedPassword), token)

	if err != nil {
		return entity.Error(ctx, fiber.ErrInternalServerError.Code, constant.Errors["internalError"].Message, constant.Errors["internalError"].Clue)
	}

	return entity.Success(ctx, nil, "Reset Password Successfully")
}
