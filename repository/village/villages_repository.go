package village

import (
	"github.com/abiyyu03/siruta/config"
	"github.com/abiyyu03/siruta/entity/model"
)

type VillageRepository struct{}

func (v *VillageRepository) Fetch() ([]*model.Village, error) {
	var villages []*model.Village

	if err := config.DB.Find(&villages).Error; err != nil {
		return nil, err
	}

	return villages, nil
}

func (v *VillageRepository) FetchById(id int) (*model.Village, error) {
	var village *model.Village

	if err := config.DB.Where("id = ?", id).First(&village).Error; err != nil {
		return nil, err
	}

	return village, nil
}

func (v *VillageRepository) Store(village *model.Village) (*model.Village, error) {
	if err := config.DB.Create(&village).Error; err != nil {
		return nil, err
	}

	return village, nil
}

func (v *VillageRepository) Update(village *model.Village, id int) (*model.Village, error) {
	if err := config.DB.Where("id = ?", id).Updates(&village).Error; err != nil {
		return nil, err
	}

	return village, nil
}

func (v *VillageRepository) Delete(id int) (*model.Village, error) {
	var village model.Village

	if err := config.DB.Where("id = ?", id).Delete(&village).Error; err != nil {
		return nil, err
	}

	return nil, nil
}
