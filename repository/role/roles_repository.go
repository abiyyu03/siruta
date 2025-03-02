package role

import (
	"github.com/abiyyu03/siruta/config"
	"github.com/abiyyu03/siruta/entity/model"
)

type RoleRepository struct{}

func (r *RoleRepository) Fetch() ([]*model.Role, error) {
	var roles []*model.Role

	if err := config.DB.Find(&roles).Error; err != nil {
		return roles, nil
	}

	return roles, nil
}

func (r *RoleRepository) FetchById(id int) (*model.Role, error) {
	var role *model.Role

	if err := config.DB.Where("id = ?", id).First(&role).Error; err != nil {
		return nil, err
	}

	return role, nil
}
func (v *RoleRepository) Store(role *model.Role) (*model.Role, error) {
	if err := config.DB.Create(&role).Error; err != nil {
		return nil, err
	}

	return role, nil
}

func (v *RoleRepository) Update(role *model.Role, id int) (*model.Role, error) {
	if err := config.DB.Where("id = ?", id).Updates(&role).Error; err != nil {
		return nil, err
	}

	return role, nil
}

func (v *RoleRepository) Delete(id int) (*model.Role, error) {
	var role model.Role

	if err := config.DB.Where("id = ?", id).Delete(&role).Error; err != nil {
		return nil, err
	}

	return nil, nil
}
