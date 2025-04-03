package finance

import (
	"github.com/abiyyu03/siruta/config"
	"github.com/abiyyu03/siruta/entity/model"
)

type ExpenseRepository struct{}

func (e *ExpenseRepository) Fetch() ([]*model.Expense, error) {
	var expenses []*model.Expense

	if err := config.DB.Find(&expenses).Error; err != nil {
		return nil, err
	}

	return expenses, nil
}

func (e *ExpenseRepository) FetchById(id int) (*model.Expense, error) {
	var expense *model.Expense

	if err := config.DB.Where("id = ?", id).First(&expense).Error; err != nil {
		return nil, err
	}

	return expense, nil
}

func (e *ExpenseRepository) FetchByRTProfileId(rtProfileId string) ([]*model.Expense, error) {
	var expenses []*model.Expense

	if err := config.DB.Where("rt_profile_id = ?", rtProfileId).Find(&expenses).Error; err != nil {
		return nil, err
	}

	return expenses, nil
}

func (e *ExpenseRepository) Store(expenseData *model.Expense) (*model.Expense, error) {
	if err := config.DB.Create(&expenseData).Error; err != nil {
		return nil, err
	}

	return expenseData, nil
}

func (e *ExpenseRepository) Update(expenseData *model.Expense, id int) (*model.Expense, error) {
	if err := config.DB.Where("id =?", id).Updates(&expenseData).Error; err != nil {
		return nil, err
	}

	return expenseData, nil
}

func (e *ExpenseRepository) Delete(id int) (*model.Expense, error) {
	var expense *model.Expense

	if err := config.DB.Where("id =?", id).Delete(&expense).Error; err != nil {
		return nil, err
	}

	return expense, nil
}
