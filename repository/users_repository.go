package repository

import (
	"github.com/abiyyu03/siruta/config"
	"github.com/abiyyu03/siruta/entity/model"
	"gorm.io/gorm"
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

func (u *UserRepository) RegisterUser(tx *gorm.DB, user *model.User, roleId uint) (*model.User, error) {
	if err := tx.Create(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
