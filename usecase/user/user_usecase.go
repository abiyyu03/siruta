package user

import (
	"github.com/abiyyu03/siruta/entity"
	"github.com/abiyyu03/siruta/entity/constant"
	"github.com/abiyyu03/siruta/entity/model"
	"github.com/abiyyu03/siruta/repository/register"
	"github.com/abiyyu03/siruta/repository/user"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type UserUsecase struct {
	db *gorm.DB
}

var userRepository = new(user.UserRepository)

func (u *UserUsecase) Fetch(ctx *fiber.Ctx) error {
	users, err := userRepository.Fetch()

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	var userResponses []entity.UserResponse
	for _, u := range users {
		userResponses = append(userResponses, entity.UserResponse{
			ID:     u.ID,
			RoleID: u.RoleID,
			Email:  u.Email,
		})
	}

	return entity.Success(ctx, &userResponses, "Data fetched successfully")
}

func (u *UserUsecase) FetchById(ctx *fiber.Ctx, id string) error {
	user, err := userRepository.FetchById(id)

	if user == nil {
		return entity.Error(ctx, fiber.StatusNotFound, constant.Errors["NotFound"].Message, constant.Errors["NotFound"].Clue)
	}

	if err != nil {
		return entity.Error(ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}
	userResponse := entity.UserResponse{
		ID:     user.ID,
		RoleID: user.RoleID,
		Email:  user.Email,
	}

	return entity.Success(ctx, &userResponse, "Data fetched successfully")
}

func (u *UserUsecase) Register(ctx *fiber.Ctx) error {
	var user *model.User

	if err := ctx.BodyParser(&user); err != nil {
		return entity.Error(
			ctx,
			fiber.ErrBadRequest.Code,
			constant.Errors["UnprocessableEntity"].Message,
			constant.Errors["UnprocessableEntity"].Clue,
		)
	}

	_, err := userRepository.RegisterUser(u.db, user, user.RoleID)

	if err != nil {
		return entity.Error(
			ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	return entity.Success(ctx, nil, "Registrasi Pengguna/User Berhasil")

}

func (u *UserUsecase) RegisterUserWithTokenVerification(ctx *fiber.Ctx, user *model.User, token string) error {
	newUser := &model.User{
		Email:    user.Email,
		Password: user.Password,
		RoleID:   4,
	}
	user, status, err := u.registerUserWithTokenVerification(newUser, 4, token)

	if status == "invalid" {
		return entity.Error(ctx, fiber.StatusForbidden, constant.Errors["InvalidToken"].Message, constant.Errors["InvalidToken"].Clue)
	}

	if err != nil {
		return entity.Error(
			ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
	}

	return entity.Success(ctx, user, "Registrasi Berhasil")
}

func (u *UserUsecase) registerUserWithTokenVerification(user *model.User, roleId uint, token string) (*model.User, string, error) {
	regToken := &register.RegTokenRepository{}
	tokenFetched, err := regToken.Validate(token)

	if err != nil {
		return nil, "invalid", err
	}

	if tokenFetched == "" {
		return nil, "invalid", nil
	}

	registerUser, err := userRepository.RegisterUser(u.db, user, roleId)

	if err != nil {
		return nil, "invalid", err
	}

	return registerUser, "valid", nil

}
