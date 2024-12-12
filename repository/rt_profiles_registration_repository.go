package repository

import (
	"log"
	"time"

	"github.com/abiyyu03/siruta/config"
	"github.com/abiyyu03/siruta/entity/model"
	"github.com/abiyyu03/siruta/entity/request"
	"gorm.io/gorm"
)

type RTProfileRegistrationRepository struct{}

func (r *RTProfileRegistrationRepository) Register(rtProfile *request.RTProfileRegisterRequest) (*request.RTProfileRegisterRequest, error) {

	// err := config.DB.Transaction(func(tx *gorm.DB) error {
	tx := config.DB.Begin()

	defer func() {
		if rec := recover(); rec != nil {
			tx.Rollback()
		}
	}()

	isVerified, rwProfileId, err := r.GetAndVerifyRWReferalCode(rtProfile.ReferalCode)

	if isVerified {
		return nil, err
	}

	r.CreateRTProfile(tx, rtProfile, rwProfileId)

	if err != nil {
		tx.Rollback()
		return nil, err
	}
	// if err != nil {
	// 	tx.Rollback()
	// 	return nil, err
	// }
	tx.Commit()

	return rtProfile, nil
}

func (r *RTProfileRegistrationRepository) CreateRTProfile(tx *gorm.DB, rtProfile *request.RTProfileRegisterRequest, RWProfileId string) (*model.RTProfile, error) {
	newRTProfile := model.RTProfile{
		RTNumber:    rtProfile.RTNumber,
		Latitude:    rtProfile.Latitude,
		Longitude:   rtProfile.Longitude,
		RTEmail:     rtProfile.RTEmail,
		MobilePhone: rtProfile.MobilePhone,
		RWProfileId: RWProfileId,
		// RTLogo:      rtProfile.RTLogo,
		// Description: rtProfile.Description,
	}

	if err := tx.Create(&newRTProfile).Error; err != nil {
		log.Printf("failed to create RTProfile: %v", err)
		return nil, err
	}

	return &newRTProfile, nil
}

func (r *RTProfileRegistrationRepository) GetAndVerifyRWReferalCode(inputedReferalCode string) (bool, string, error) {
	var referalCode *model.ReferalCode
	// var rtProfile *model.RTProfile

	if err := config.DB.Where("code = ? OR expired_at > ?", inputedReferalCode, time.Now()).First(&referalCode).Error; err != nil {
		log.Printf("error : %v", err)
		return false, "", err
	}

	if &referalCode == nil {
		return false, "", nil
	}

	return true, referalCode.RWProfileId, nil
}

func (r *RTProfileRegistrationRepository) CheckRTNumberAvailability(rtProfile *model.RTProfile, rtNumber string) (bool, error) {
	if err := config.DB.Where("rt_number = ?", rtNumber).First(&rtProfile).Error; err != nil {
		return false, err
	}

	return true, nil
}

func (r *RTProfileRegistrationRepository) ApproveRegistrant(rtProfileId string) (*model.RTProfile, error) {
	var rtProfile *model.RTProfile

	if err := config.DB.First(&rtProfile, rtProfileId).Error; err != nil {
		return nil, err
	}

	return rtProfile, nil

}
