package letter

import (
	"github.com/abiyyu03/siruta/config"
	"github.com/abiyyu03/siruta/entity/model"
)

type OutcomingLetterRepository struct{}

func (o *OutcomingLetterRepository) Fetch() ([]*model.OutcomingLetter, error) {
	var outcomingLetters []*model.OutcomingLetter

	if err := config.DB.Preload("Member").Preload("LetterType").Find(&outcomingLetters).Error; err != nil {
		return nil, err
	}

	return outcomingLetters, nil
}

func (o *OutcomingLetterRepository) FetchById(id string) (*model.OutcomingLetter, error) {
	var outcomingLetter *model.OutcomingLetter

	if err := config.DB.Preload("Member").Preload("LetterType").Where("id =?", id).First(&outcomingLetter).Error; err != nil {
		return nil, err
	}

	return outcomingLetter, nil
}

func (o *OutcomingLetterRepository) FetchByRTProfileId(rtProfileId string) ([]*model.OutcomingLetter, error) {
	var outcomingLetters []*model.OutcomingLetter

	if err := config.DB.Preload("Member").Preload("LetterType").Where("rt_profile_id =?", rtProfileId).Find(&outcomingLetters).Error; err != nil {
		return nil, err
	}

	return outcomingLetters, nil
}

func (o *OutcomingLetterRepository) FetchPreview(id string) (*model.OutcomingLetter, error) {
	var letterReq *model.OutcomingLetter

	if err := config.DB.Preload("Member").Preload("LetterType").Preload("RTProfile").Where("is_rt_approved =?", true).Where("id =?", id).First(&letterReq).Error; err != nil {
		return nil, err
	}

	return letterReq, nil
}
