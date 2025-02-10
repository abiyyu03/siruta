package usecase

import (
	"github.com/abiyyu03/siruta/entity"
	"github.com/abiyyu03/siruta/entity/model"
	"github.com/gofiber/fiber/v2"
)

type UserUsecase struct{}

// var c = new(entity.Constant)

func (u *UserUsecase) RegisterUserWithTokenVerification(ctx *fiber.Ctx, user *model.User, token string) error {
	newUser := &model.User{
		Email:    user.Email,
		Username: user.Username,
		Password: user.Password,
		RoleID:   4,
	}
	user, status, err := authRepository.RegisterUserWithTokenVerification(newUser, 4, token)

	if status == "invalid" {
		return entity.Error(ctx, fiber.StatusForbidden, "Token verification failed")
	}

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, fiber.ErrInternalServerError.Message)
	}

	return entity.Success(ctx, user, "Registration successful")
}
