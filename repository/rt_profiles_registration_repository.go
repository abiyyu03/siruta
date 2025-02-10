package repository

import (
	"log"
	"time"

	"github.com/abiyyu03/siruta/config"
	"github.com/abiyyu03/siruta/entity/model"
	"gorm.io/gorm"
)

type RTProfileRegistrationRepository struct{}

// var emailNotificationRepository = new(email.EmailUserRegistrationRepository)

func (r *RTProfileRegistrationRepository) Register(rtProfile *model.RTProfile, referalCode string) (*model.RTProfile, error) {
	tx := config.DB.Begin()

	defer func() {
		if rec := recover(); rec != nil {
			tx.Rollback()
		}
	}()

	isVerified, rwProfileId, err := r.getAndVerifyRWReferalCode(referalCode)

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if !isVerified {
		return nil, err
	}

	_, err = r.createRTProfile(tx, rtProfile, rwProfileId)

	if err != nil {
		log.Printf("failed : %v", err)
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return rtProfile, nil
}

func (r *RTProfileRegistrationRepository) createRTProfile(tx *gorm.DB, rtProfile *model.RTProfile, rwProfileId string) (*model.RTProfile, error) {
	rtProfile.RWProfileId = rwProfileId

	if err := tx.Debug().Create(&rtProfile).Error; err != nil {
		log.Printf("failed to create RTProfile: %v", err)
		return nil, err
	}

	return rtProfile, nil
}

func (r *RTProfileRegistrationRepository) getAndVerifyRWReferalCode(inputedReferalCode string) (bool, string, error) {
	var referalCode *model.ReferalCode

	if err := config.DB.Where("code = ? AND expired_at > ?", inputedReferalCode, time.Now()).First(&referalCode).Error; err != nil {
		log.Printf("error : %v", err)
		return false, "", err
	}

	if referalCode == nil {
		return false, "", nil
	}

	return true, referalCode.RWProfileId, nil
}

func (r *RTProfileRegistrationRepository) CheckRtNumberAvailability(rtProfile *model.RTProfile, RtNumber string) (bool, error) {
	if err := config.DB.Where("rt_number = ?", RtNumber).First(&rtProfile).Error; err != nil {
		return false, err
	}

	if rtProfile == nil {
		return false, nil
	}

	return true, nil
}

func (r *RTProfileRegistrationRepository) ApproveRegistrant(rtProfileId string) (*model.RTProfile, error) {
	var rtProfile *model.RTProfile

	if err := config.DB.Model(&rtProfile).Where("id = ?", rtProfileId).Update("is_authorized", true).Error; err != nil {
		return nil, err
	}

	return rtProfile, nil
}
