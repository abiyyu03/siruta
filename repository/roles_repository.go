package repository

import (
	"github.com/abiyyu03/siruta/config"
	"github.com/abiyyu03/siruta/entity/model"
)

type RolesRepository struct{}

func (r *RolesRepository) Fetch() ([]*model.Role, error) {
	var roles []*model.Role

	if err := config.DB.Find(&roles).Error; err != nil {
		return roles, nil
	}

	return roles, nil
}

func (r *RolesRepository) FetchById(id float64) (*model.Role, error) {
	var role *model.Role

	if err := config.DB.Select("id").First(&role, id).Error; err != nil {
		return nil, err
	}

	return role, nil
}
