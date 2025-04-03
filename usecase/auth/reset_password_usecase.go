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

type ResetPasswordUsecase struct{}

var resetPasswordRepository *auth.ResetPasswordRepository
var sendEmail *email.ResetPasswordEmailUsecase
var userRepository *user.UserRepository

func (r *ResetPasswordUsecase) SendForgotPasswordLink(ctx *fiber.Ctx, email string) error {
	randomToken, _ := helper.GenerateSecureToken(48)

	user, err := userRepository.FetchByEmail(email)

	if err != nil {
		return entity.Error(ctx, fiber.ErrInternalServerError.Code, constant.Errors["internalError"].Message, constant.Errors["internalError"].Clue)
	}

	createdToken := &model.ResetPassword{
		Token:     randomToken,
		UserID:    user.ID,
		ExpiredAt: time.Now(),
	}

	_, err = resetPasswordRepository.StoreToken(createdToken)

	if err != nil {
		return entity.Error(ctx, fiber.ErrInternalServerError.Code, constant.Errors["internalError"].Message, constant.Errors["internalError"].Clue)
	}

	err = sendEmail.ResetPasswordEmail(email, randomToken)

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

	err = resetPasswordRepository.ResetPassword(string(hashedPassword), token)

	if err != nil {
		return entity.Error(ctx, fiber.ErrInternalServerError.Code, constant.Errors["internalError"].Message, constant.Errors["internalError"].Clue)
	}

	return entity.Success(ctx, nil, "Reset Password Successfully")
}
