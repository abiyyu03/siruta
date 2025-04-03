package finance

import (
	"github.com/abiyyu03/siruta/config"
	"github.com/abiyyu03/siruta/entity/model"
)

type IncomePlanRepository struct{}

func (i *IncomePlanRepository) Fetch() ([]*model.IncomePlan, error) {
	var incomePlans []*model.IncomePlan

	if err := config.DB.Find(&incomePlans).Error; err != nil {
		return nil, err
	}

	return incomePlans, nil
}

func (i *IncomePlanRepository) FetchById(id int) (*model.IncomePlan, error) {
	var incomePlan *model.IncomePlan

	if err := config.DB.Where("id = ?", id).First(&incomePlan).Error; err != nil {
		return nil, err
	}

	return incomePlan, nil
}

func (i *IncomePlanRepository) FetchByRTProfileId(rtProfileId string) ([]*model.IncomePlan, error) {
	var incomePlans []*model.IncomePlan

	if err := config.DB.Where("rt_profile_id = ?", rtProfileId).Find(&incomePlans).Error; err != nil {
		return nil, err
	}

	return incomePlans, nil
}

func (i *IncomePlanRepository) Store(incomePlanData *model.IncomePlan) (*model.IncomePlan, error) {
	if err := config.DB.Create(&incomePlanData).Error; err != nil {
		return nil, err
	}

	return incomePlanData, nil
}

func (i *IncomePlanRepository) Update(incomePlanData *model.IncomePlan, id int) (*model.IncomePlan, error) {
	if err := config.DB.Where("id =?", id).Updates(&incomePlanData).Error; err != nil {
		return nil, err
	}

	return incomePlanData, nil
}

func (i *IncomePlanRepository) Delete(id int) (*model.IncomePlan, error) {
	var incomePlan *model.IncomePlan

	if err := config.DB.Where("id =?", id).Delete(&incomePlan).Error; err != nil {
		return nil, err
	}

	return incomePlan, nil
}
