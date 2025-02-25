package repository

import (
	"github.com/abiyyu03/siruta/config"
	"github.com/abiyyu03/siruta/entity/model"
	"golang.org/x/crypto/bcrypt"
)

type AuthRepository struct{}

func (u *AuthRepository) FetchLogin(username, password string) (*model.User, *model.Member, error) {
	var user model.User
	var member model.Member

	if err := config.DB.Preload("Role").Where("username = ?", username).First(&user).Error; err != nil {
		return nil, nil, err
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, nil, err
	}

	fetchedMember := config.DB.Where("user_id = ?", user.ID).First(&member)

	if fetchedMember == nil {
		return &user, nil, nil
	}

	// if fetchedMember.Error != nil {
	// 	return nil, nil, err
	// }

	return &user, &member, nil

}
