package referal_code

import (
	"errors"
	"log"
	"time"

	"github.com/abiyyu03/siruta/config"
	"github.com/abiyyu03/siruta/entity/model"
	"github.com/abiyyu03/siruta/helper"
	"gorm.io/gorm"
)

type ReferalCodeRepository struct{}

func (r *ReferalCodeRepository) Fetch() ([]*model.ReferalCode, error) {
	var referalCodes []*model.ReferalCode

	if err := config.DB.First(&referalCodes).Error; err != nil {
		return nil, err
	}

	return referalCodes, nil
}

func (r *ReferalCodeRepository) FetchById(id int) (*model.ReferalCode, error) {
	var referalCode *model.ReferalCode

	if err := config.DB.Where("id = ?", id).First(&referalCode).Error; err != nil {
		return nil, err
	}

	return referalCode, nil
}

func (r *ReferalCodeRepository) FetchByRTProfileId(rtProfileId string) ([]*model.ReferalCode, error) {
	var referalCodes []*model.ReferalCode

	if err := config.DB.Where("profile_id =?", rtProfileId).Find(&referalCodes).Error; err != nil {
		return nil, err
	}

	return referalCodes, nil
}

func (r *ReferalCodeRepository) GenerateReferalCode(tx *gorm.DB, profileId string) error {
	code := helper.RandomString(6)

	referal := &model.ReferalCode{
		Code:      code,
		ExpiredAt: time.Now().AddDate(1, 0, 0),
		ProfileId: profileId,
	}

	if err := tx.Create(&referal).Error; err != nil {
		log.Printf("failed to create referalcode: %v", err)
		return err
	}

	return nil
}

func (r *ReferalCodeRepository) RegenerateReferalCode(code string, profileId string) (string, error) {
	var referalCode *model.ReferalCode
	regeneratedCode := helper.RandomString(6)

	if err := config.DB.Where("code = ? AND profile_id = ?", code, profileId).First(&referalCode).Error; err != nil {
		log.Printf("Gagal mengambil referal code: %v", err)
		return "", err
	}

	if err := config.DB.Model(&referalCode).Update("code", regeneratedCode).Error; err != nil {
		log.Printf("failed to create referalcode: %v", err)
		return "", err
	}

	return regeneratedCode, nil
}

func (r *ReferalCodeRepository) GetAndVerifyRWReferalCode(inputedReferalCode string) (bool, string, error) {
	var referalCode *model.ReferalCode

	if err := config.DB.Where("code = ? ", inputedReferalCode).Where("expired_at > ?", time.Now()).First(&referalCode).Error; err != nil {
		return false, "", err
	}

	if referalCode == nil {
		return false, "", errors.New("kode referal tidak valid")
	}

	return true, referalCode.ProfileId, nil
}

func (r *ReferalCodeRepository) Validate(code string) (string, bool, error) { //profileType like rw, rt
	var referalCode *model.ReferalCode

	if err := config.DB.Where("code = ? AND expired_at > ?", code, time.Now()).First(&referalCode).Error; err != nil {
		log.Print("Error referal code : ", err.Error())
		return "", false, err
	}

	if referalCode == nil {
		return "", false, nil
	}

	return referalCode.ProfileId, true, nil
}

func (r *ReferalCodeRepository) Delete(id int) error {
	var referalCode *model.ReferalCode

	if err := config.DB.Where("id =?", id).Delete(&referalCode).Error; err != nil {
		return err
	}

	return nil
}
