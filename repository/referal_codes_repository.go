package repository

import (
	"errors"
	"time"

	"github.com/abiyyu03/siruta/config"
	"github.com/abiyyu03/siruta/entity/model"
)

type ReferalCodeRepository struct{}

var rwProfileRepository = new(RWProfileRepository)

func (r *ReferalCodeRepository) Fetch() ([]*model.ReferalCode, error) {
	var referalCodes []*model.ReferalCode

	if err := config.DB.First(&referalCodes).Error; err != nil {
		return nil, err
	}

	return referalCodes, nil
}

func (r *ReferalCodeRepository) FetchById(id string) (*model.ReferalCode, error) {
	var referalCode *model.ReferalCode

	if err := config.DB.Where("id = ?", id).First(&referalCode).Error; err != nil {
		return nil, err
	}

	return referalCode, nil
}

func (r *ReferalCodeRepository) GetAndVerifyRWReferalCode(inputedReferalCode string) (bool, string, error) {
	var referalCode *model.ReferalCode

	if err := config.DB.Where("code = ? ", inputedReferalCode).Where("expired_at > ?", time.Now()).First(&referalCode).Error; err != nil {
		return false, "", err
	}

	if referalCode == nil {
		return false, "", errors.New("Kode referal tidak valid")
	}

	return true, referalCode.RWProfileId, nil
}

func (r *ReferalCodeRepository) Validate(code string) (*model.RWProfile, bool, error) {
	var referalCode *model.ReferalCode

	if err := config.DB.Where("code = ? AND expired_at > ?", code, time.Now()).First(&referalCode).Error; err != nil {
		return nil, false, err
	}

	if referalCode == nil {
		return nil, false, nil
	}

	rwProfile, err := rwProfileRepository.FetchById(referalCode.RWProfileId)

	if err != nil {
		return nil, false, err
	}

	return rwProfile, true, nil
}
