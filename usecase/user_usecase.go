package usecase

import (
	"github.com/abiyyu03/siruta/entity"
	"github.com/abiyyu03/siruta/entity/model"
	"github.com/abiyyu03/siruta/repository"
	"github.com/abiyyu03/siruta/repository/register"
	"github.com/gofiber/fiber/v2"
)

type UserUsecase struct{}

var userRepository = new(repository.UserRepository)

func (u *UserUsecase) Fetch(ctx *fiber.Ctx) error {
	users, err := userRepository.Fetch()

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, "Server error")
	}

	var userResponses []entity.UserResponse
	for _, u := range users {
		userResponses = append(userResponses, entity.UserResponse{
			ID:       u.ID,
			RoleID:   u.RoleID,
			Email:    u.Email,
			Username: u.Username,
		})
	}

	return entity.Success(ctx, &userResponses, "Data fetched successfully")
}

func (u *UserUsecase) FetchById(ctx *fiber.Ctx, id string) error {
	user, err := userRepository.FetchById(id)

	if user == nil {
		return entity.Error(ctx, fiber.StatusNotFound, "user not found")
	}

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, fiber.ErrInternalServerError.Message)
	}
	userResponse := entity.UserResponse{
		ID:       user.ID,
		RoleID:   user.RoleID,
		Email:    user.Email,
		Username: user.Username,
	}

	return entity.Success(ctx, &userResponse, "Data fetched successfully")
}

func (u *UserUsecase) Register(ctx *fiber.Ctx) error {
	var user *model.User

	if err := ctx.BodyParser(&user); err != nil {
		return entity.Error(ctx, fiber.StatusUnprocessableEntity, fiber.ErrUnprocessableEntity.Message)
	}

	_, err := userRepository.RegisterUser(user, user.RoleID)

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, fiber.ErrInternalServerError.Message)
	}

	return entity.Success(ctx, nil, "Register user successfully")

}

func (u *UserUsecase) RegisterUserWithTokenVerification(ctx *fiber.Ctx, user *model.User, token string) error {
	newUser := &model.User{
		Email:    user.Email,
		Username: user.Username,
		Password: user.Password,
		RoleID:   4,
	}
	user, status, err := u.registerUserWithTokenVerification(newUser, 4, token)

	if status == "invalid" {
		return entity.Error(ctx, fiber.StatusForbidden, "Token verification failed")
	}

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, fiber.ErrInternalServerError.Message)
	}

	return entity.Success(ctx, user, "Registration successful")
}

func (u *UserUsecase) registerUserWithTokenVerification(user *model.User, roleId uint, token string) (*model.User, string, error) {
	regToken := &register.RegTokenRepository{}
	isTokenValid, err := regToken.Validate(token)

	if err != nil {
		return nil, "invalid", err
	}

	if !isTokenValid {
		return nil, "invalid", nil
	}

	registerUser, err := userRepository.RegisterUser(user, roleId)

	if err != nil {
		return nil, "invalid", err
	}

	return registerUser, "valid", nil

}
