package letter

import (
	"github.com/abiyyu03/siruta/config"
	"github.com/abiyyu03/siruta/entity/model"
)

type IncomingLetterRepository struct{}

func (i *IncomingLetterRepository) Fetch() ([]*model.IncomingLetter, error) {
	var incomingLetters []*model.IncomingLetter

	if err := config.DB.Find(&incomingLetters).Error; err != nil {
		return nil, err
	}

	return incomingLetters, nil
}

func (i *IncomingLetterRepository) FetchById(id int) (*model.IncomingLetter, error) {
	var incomingLetter *model.IncomingLetter

	if err := config.DB.Where("id = ?", id).First(&incomingLetter).Error; err != nil {
		return nil, err
	}

	return incomingLetter, nil
}

func (i *IncomingLetterRepository) FetchByRTProfileId(rwProfileId string) ([]*model.IncomingLetter, error) {
	var incomingLetters []*model.IncomingLetter

	if err := config.DB.Where("rt_profile_id =?", rwProfileId).Find(&incomingLetters).Error; err != nil {
		return nil, err
	}

	return incomingLetters, nil
}

func (i *IncomingLetterRepository) Store(incomingLetter *model.IncomingLetter) (*model.IncomingLetter, error) {
	if err := config.DB.Create(&incomingLetter).Error; err != nil {
		return nil, err
	}

	return incomingLetter, nil
}

func (i *IncomingLetterRepository) Update(incomingLetter *model.IncomingLetter, id int) (*model.IncomingLetter, error) {
	if err := config.DB.Model(&incomingLetter).Updates(&incomingLetter).Error; err != nil {
		return nil, err
	}

	return incomingLetter, nil
}

func (i *IncomingLetterRepository) Delete(incomingLetter *model.IncomingLetter, id int) (*model.IncomingLetter, error) {
	if err := config.DB.Where("id = ?", id).Delete(&incomingLetter).Error; err != nil {
		return nil, err
	}

	return incomingLetter, nil
}
