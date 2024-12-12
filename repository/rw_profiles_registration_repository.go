package repository

import (
	"log"
	"time"

	"github.com/abiyyu03/siruta/config"
	"github.com/abiyyu03/siruta/entity/model"
	"github.com/abiyyu03/siruta/usecase/helper"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RWProfileRegistrationRepository struct {
	RWProfileData *model.RWProfile
}

// func (r *RWProfileRegistrationRepository) Register(rwProfile *model.RWProfile) (*model.RWProfile, error) {

// 	err := config.DB.Transaction(func(tx *gorm.DB) error {

// 		newRWProfile, err := r.CreateRWProfile(tx, rwProfile)

// 		if err != nil {
// 			log.Printf("failed to create RWProfile: %v", err)
// 			return err
// 		}

// 		err = r.GenerateReferalCode(tx, newRWProfile.ID)
// 		if err != nil {
// 			log.Printf("Failed to create ReferralCode: %v", err)
// 			return err // Trigger rollback
// 		}

// 		r.RWProfileData = newRWProfile

// 		return nil
// 	})

// 	if err != nil {
// 		log.Printf("Transaction failed, rolled back due to error: %v", err)
// 		return nil, err
// 	}

// 	log.Println("Transaction committed successfully.")

// 	return r.RWProfileData, nil
// }

func (r *RWProfileRegistrationRepository) RegisterRWProfile(rwProfileRequest *model.RWProfile) (*model.RWProfile, error) {
	// var rwModel *model.RWProfile

	id := uuid.New()

	newRWProfile := &model.RWProfile{
		ID:          id.String(),
		RWNumber:    rwProfileRequest.RWNumber,
		VillageID:   rwProfileRequest.VillageID,
		RwEmail:     rwProfileRequest.RwEmail,
		MobilePhone: rwProfileRequest.MobilePhone,
		Description: rwProfileRequest.Description,
		RWLogo:      rwProfileRequest.RWLogo,
		RegencyLogo: rwProfileRequest.RegencyLogo,
	}

	if err := config.DB.Create(&newRWProfile).Error; err != nil {
		log.Panic("wadidaw")
		log.Printf("failed to create RWProfile: %v", err)
		return nil, err
	}

	return newRWProfile, nil
}

func (r *RWProfileRegistrationRepository) GenerateReferalCode(tx *gorm.DB, rwProfileId string) error {
	code := helper.RandomString(6)

	referal := model.ReferalCode{
		Code:        code,
		ExpiredAt:   time.Now().AddDate(1, 0, 0),
		RWProfileId: rwProfileId,
	}

	if err := tx.Create(&referal).Error; err != nil {
		log.Printf("failed to create referalcode: %v", err)
		return err
	}

	return nil
}

func (r *RWProfileRegistrationRepository) FetchRWProfile(id string, rwProfile *model.RWProfile) (*model.RWProfile, error) {
	if err := config.DB.First(&rwProfile, "id = ?", id).Error; err != nil {
		log.Printf("hahaa %v", err)
		return nil, err
	}

	return rwProfile, nil
}

func (r *RWProfileRegistrationRepository) CheckRWAuthorizationIsTrue(id string, rwProfile *model.RWProfile) bool {
	fetchedRWProfile, err := r.FetchRWProfile(id, rwProfile)

	if err != nil {
		return false
	}

	if fetchedRWProfile.IsAuthorized {
		return true
	}

	return false
}

func (r *RWProfileRegistrationRepository) UpdateRWAuthorization(tx *gorm.DB, id string, rwProfile *model.RWProfile) error {
	if err := tx.Model(&rwProfile).Where("id = ?", id).Update("is_authorized", true).Error; err != nil {
		return err
	}
	return nil
}

func (r *RWProfileRegistrationRepository) ApproveRegistrant(rwProfileId string) error {
	err := config.DB.Transaction(func(tx *gorm.DB) error {
		var rwProfile *model.RWProfile

		// Fetch the RW profile data
		if _, err := r.FetchRWProfile(rwProfileId, rwProfile); err != nil {
			log.Printf("Failed to fetch RWProfile: %v", err)
			return err // Trigger rollback on error
		}

		// Check if authorization is already true
		if r.CheckRWAuthorizationIsTrue(rwProfileId, rwProfile) {
			log.Printf("RWProfile %v is already authorized", rwProfileId)
			return nil // No update needed
		}
		// log.Printf("User is not authorized yet: %v", rwProfile)

		// Update authorization status
		if err := r.UpdateRWAuthorization(tx, rwProfileId, rwProfile); err != nil {
			log.Printf("Failed to update authorization: %v", err)
			return err // Trigger rollback
		}
		// log.Printf("RWProfile %v authorized successfully", rwProfileId)

		// Generate referral code
		if err := r.GenerateReferalCode(tx, rwProfile.ID); err != nil {
			log.Printf("Failed to generate referral code: %v", err)
			return err // Trigger rollback
		}
		// log.Printf("Referral code generated successfully for RWProfile %v", rwProfileId)

		return nil // Commit transaction if no errors
	})
	if err != nil {
		log.Printf("Transaction failed, rolled back due to error: %v", err)
		return err
	}

	log.Println("Transaction committed successfully.")

	return nil
}

// func (r *RWProfileRegistrationRepository) CheckRWNumberAvailability() (bool, error) {
// 	var rwProfile *model.RWProfile

// 	if err := config.DB.Where("rw_number = ?").First(&rwProfile).Error; err != nil {
// 		return false, err
// 	}

// 	return true, nil
// }
