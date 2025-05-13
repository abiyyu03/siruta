package finance

import (
	"github.com/abiyyu03/siruta/config"
	"github.com/abiyyu03/siruta/entity/model"
)

type CashflowRepository struct{}

func (c *CashflowRepository) Fetch(logType string) ([]*model.Cashflow, error) {
	var incomes []*model.Cashflow

	if err := config.DB.Where("log_type =?", logType).Find(&incomes).Error; err != nil {
		return nil, err
	}

	return incomes, nil
}

func (c *CashflowRepository) FetchById(id int, logType string) (*model.Cashflow, error) {
	var income *model.Cashflow

	if err := config.DB.Where("log_type =?", logType).Where("id =?", id).First(&income).Error; err != nil {
		return nil, err
	}

	return income, nil
}

func (c *CashflowRepository) FetchByRTProfileId(rtProfileId string, logType string) ([]*model.Cashflow, error) {
	var incomes []*model.Cashflow

	if err := config.DB.Where("log_type =?", logType).Where("rt_profile_id = ?", rtProfileId).Find(&incomes).Error; err != nil {
		return nil, err
	}

	return incomes, nil
}

func (c *CashflowRepository) Store(incomeData *model.Cashflow) (*model.Cashflow, error) {
	if err := config.DB.Create(&incomeData).Error; err != nil {
		return nil, err
	}

	return incomeData, nil
}

func (c *CashflowRepository) Update(incomeData *model.Cashflow, id int) (*model.Cashflow, error) {
	if err := config.DB.Where("id =?", id).Updates(&incomeData).Error; err != nil {
		return nil, err
	}

	return incomeData, nil
}

func (c *CashflowRepository) Delete(id int) error {
	var income *model.Cashflow

	if err := config.DB.Where("id =?", id).Delete(&income).Error; err != nil {
		return err
	}

	return nil
}
