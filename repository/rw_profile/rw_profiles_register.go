package rw_profile

import (
	"log"

	"github.com/abiyyu03/siruta/config"
	"github.com/abiyyu03/siruta/entity/model"
	"github.com/abiyyu03/siruta/repository/member"
	"github.com/abiyyu03/siruta/repository/referal_code"
	"github.com/abiyyu03/siruta/repository/register"
	"github.com/abiyyu03/siruta/repository/user"
	"gorm.io/gorm"
)

type RWProfileRegisterRepository struct{}

var RWProfileData *model.RWProfile
var userRepository *user.UserRepository
var memberRepository *member.MemberRepository
var tokenRegisterRepository *register.RegTokenRepository
var referalCode *referal_code.ReferalCodeRepository
var leaderRepository *RWLeaderRepository

func (r *RWProfileRegisterRepository) RegisterRWProfile(rwProfileRequest *model.RWProfile) (*model.RWProfile, error) {
	if err := config.DB.Create(&rwProfileRequest).Error; err != nil {
		return nil, err
	}

	return rwProfileRequest, nil
}

func (r *RWProfileRegisterRepository) FetchRWProfile(id string, rwProfile *model.RWProfile) (*model.RWProfile, error) {
	if err := config.DB.First(&rwProfile, "id = ?", id).Error; err != nil {
		log.Printf("hahaa %v", err)
		return nil, err
	}

	return rwProfile, nil
}

func (r *RWProfileRegisterRepository) CheckRWAuthorizationIsTrue(id string, rwProfile *model.RWProfile) bool {
	fetchedRWProfile, err := r.FetchRWProfile(id, rwProfile)

	if err != nil {
		return false
	}

	if fetchedRWProfile.IsAuthorized {
		return true
	}

	return false
}

func (r *RWProfileRegisterRepository) UpdateRWAuthorization(tx *gorm.DB, id string, rwProfile *model.RWProfile) error {
	if err := tx.Model(&rwProfile).Where("id = ?", id).Update("is_authorized", true).Error; err != nil {
		return err
	}
	return nil
}

func (r *RWProfileRegisterRepository) ApproveRegistrant(rwProfileId string) error {
	err := config.DB.Transaction(func(tx *gorm.DB) error {
		var rwProfile model.RWProfile

		// Fetch the RW profile data
		_, err := r.FetchRWProfile(rwProfileId, &rwProfile)

		if err != nil {
			log.Printf("Failed to fetch RWProfile: %v", err.Error())
			return err // Trigger rollback on error
		}

		// Check if authorization is already true
		if r.CheckRWAuthorizationIsTrue(rwProfileId, &rwProfile) {
			log.Printf("RWProfile %v is already authorized", rwProfileId)
			return nil // No update needed
		}

		// Update authorization status
		if err := r.UpdateRWAuthorization(tx, rwProfileId, &rwProfile); err != nil {
			log.Printf("Failed to update authorization: %v", err)
			return err // Trigger rollback
		}

		// Generate referral code
		if err := referalCode.GenerateReferalCode(tx, rwProfile.ID); err != nil {
			log.Printf("Failed to generate referral code: %v", err)
			return err // Trigger rollback
		}

		return nil // Commit transaction if no errors
	})

	if err != nil {
		log.Printf("Transaction failed, rolled back due to error: %v", err)
		return err
	}

	log.Println("Transaction committed successfully.")

	return nil
}

func (r *RWProfileRegisterRepository) RegisterUserRW(leader *model.RWLeader, user *model.User, roleId uint, token string) error {
	err := config.DB.Transaction(func(tx *gorm.DB) error {
		_, err := userRepository.RegisterUser(tx, user, roleId)

		if err != nil {
			return err
		}

		err = leaderRepository.Store(tx, leader)

		if err != nil {
			return err
		}

		_, err = tokenRegisterRepository.RemoveToken(tx, token)

		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		log.Printf("Transaction failed, rolled back due to error: %v", err.Error())
		return err
	}

	log.Println("Transaction committed successfully.")

	return nil

}
