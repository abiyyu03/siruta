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
