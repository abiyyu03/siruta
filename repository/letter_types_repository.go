package repository

import (
	"github.com/abiyyu03/siruta/config"
	"github.com/abiyyu03/siruta/entity/model"
)

type LetterTypeRepository struct{}

func (l *LetterTypeRepository) Fetch() ([]*model.LetterType, error) {
	var letterTypes []*model.LetterType

	if err := config.DB.Find(&letterTypes).Error; err != nil {
		return nil, err
	}

	return letterTypes, nil
}

func (l *LetterTypeRepository) FetchById(id string) (*model.LetterType, error) {
	var letterType *model.LetterType

	if err := config.DB.Where("id = ?", id).First(&letterType).Error; err != nil {
		return nil, err
	}

	return letterType, nil
}
