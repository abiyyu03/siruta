package repository

import (
	"github.com/abiyyu03/siruta/config"
	"github.com/abiyyu03/siruta/entity/model"
)

type ReligionRepository struct{}

func (r *ReligionRepository) Fetch() ([]*model.Religion, error) {
	var religions []*model.Religion

	if err := config.DB.Find(&religions).Error; err != nil {
		return nil, err
	}

	return religions, nil
}

func (r *ReligionRepository) FetchById(id int) (*model.Religion, error) {
	var religion *model.Religion

	if err := config.DB.Where("id = ?", id).First(&religion).Error; err != nil {
		return nil, err
	}

	return religion, nil
}

func (r *ReligionRepository) Store(religion *model.Religion) (*model.Religion, error) {
	if err := config.DB.Create(&religion).Error; err != nil {
		return nil, err
	}

	return religion, nil
}

func (r *ReligionRepository) Update(religion *model.Religion, id int) (*model.Religion, error) {
	if err := config.DB.Where("id = ?", id).Updates(&religion).Error; err != nil {
		return nil, err
	}

	return religion, nil
}

func (v *ReligionRepository) Delete(id int) (*model.Religion, error) {
	var religion model.Religion

	if err := config.DB.Where("id = ?", id).Delete(&religion).Error; err != nil {
		return nil, err
	}

	return nil, nil
}
