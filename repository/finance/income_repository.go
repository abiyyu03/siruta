package finance

import (
	"github.com/abiyyu03/siruta/config"
	"github.com/abiyyu03/siruta/entity/model"
)

type IncomeRepository struct{}

func (i *IncomeRepository) Fetch() ([]*model.Income, error) {
	var incomes []*model.Income

	if err := config.DB.Find(&incomes).Error; err != nil {
		return nil, err
	}

	return incomes, nil
}

func (i *IncomeRepository) FetchById(id int) (*model.Income, error) {
	var income *model.Income

	if err := config.DB.Where("id =?", id).First(&income).Error; err != nil {
		return nil, err
	}

	return income, nil
}

func (i *IncomeRepository) FetchByPlanId(planId string) ([]*model.Income, error) {
	var incomes []*model.Income

	if err := config.DB.Where("plan_id =?", planId).Find(&incomes).Error; err != nil {
		return nil, err
	}

	return incomes, nil
}

func (i *IncomeRepository) Store(incomeData *model.Income) (*model.Income, error) {
	if err := config.DB.Create(&incomeData).Error; err != nil {
		return nil, err
	}

	return incomeData, nil
}

func (i *IncomeRepository) Update(incomeData *model.Income, id int) (*model.Income, error) {
	if err := config.DB.Where("id =?", id).Updates(&incomeData).Error; err != nil {
		return nil, err
	}

	return incomeData, nil
}

func (i *IncomeRepository) Delete(id int) (*model.Income, error) {
	var income *model.Income

	if err := config.DB.Where("id =?", id).Delete(&income).Error; err != nil {
		return nil, err
	}

	return income, nil
}
