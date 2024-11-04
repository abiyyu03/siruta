package repository

import (
	"errors"

	"github.com/abiyyu03/siruta/config"
	"github.com/abiyyu03/siruta/entity/model"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct{}

func (u *UserRepository) FetchLogin(username, password string) (*model.User, error) {
	// if config.DB == nil {
	// 	return nil, errors.New("database connection is not initialized")
	// }
	var user model.User

	var query = config.DB.Where("username = ?", username).First(&user)
	if err := query.Error; err != nil {
		return nil, errors.New("Username or password not found")
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	// var encryptedPassword = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("Username or password not found")
	}

	return &user, nil

}

func (u *UserRepository) RegisterUser(user *model.User, roleId uint) (*model.User, error) {

	id := uuid.New()
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(user.Password),
		10,
	)
	if err != nil {
		return nil, err
	}

	createdUser := model.User{
		ID:        id.String(),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		RoleID:    roleId,
		Email:     user.Email,
		Username:  user.Username,
		Password:  string(hashedPassword),
	}

	config.DB.Create(&createdUser)

	return &createdUser, nil
}
