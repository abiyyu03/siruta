package repository

import (
	"github.com/abiyyu03/siruta/config"
	"github.com/abiyyu03/siruta/entity/model"
	"github.com/abiyyu03/siruta/repository/register"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct{}

func (u *UserRepository) Fetch() ([]*model.User, error) {
	var users []*model.User

	if err := config.DB.Preload("Role").Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (u *UserRepository) FetchById(id string) (*model.User, error) {
	var user *model.User

	if err := config.DB.Preload("Role").Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserRepository) FetchLogin(username, password string) (*model.User, error) {
	var user model.User

	var query = config.DB.Where("username = ?", username).First(&user)
	if err := query.Error; err != nil {
		// log.Panic(err)
		return nil, err
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, err
	}

	return &user, nil

}

func (u *UserRepository) RegisterUser(user *model.User, roleId uint) (*model.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(user.Password),
		14,
	)
	if err != nil {
		return nil, err
	}

	id := uuid.New()

	createdUser := &model.User{
		ID:       id.String(),
		RoleID:   roleId,
		Email:    user.Email,
		Username: user.Username,
		Password: string(hashedPassword),
	}

	config.DB.Create(&createdUser)

	return createdUser, nil
}

func (u *UserRepository) RegisterUserWithTokenVerification(user *model.User, roleId uint, token string) (*model.User, string, error) {
	regToken := &register.RegTokenRepository{}
	isTokenValid, err := regToken.Validate(token)

	if err != nil {
		return nil, "invalid", err
	}

	if !isTokenValid {
		return nil, "invalid", nil
	}

	registerUser, err := u.RegisterUser(user, roleId)

	if err != nil {
		return nil, "invalid", err
	}

	return registerUser, "valid", nil

}
