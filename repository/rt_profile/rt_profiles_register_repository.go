package rt_profile

import (
	"errors"
	"log"
	"time"

	"github.com/abiyyu03/siruta/config"
	"github.com/abiyyu03/siruta/entity/model"
	"github.com/abiyyu03/siruta/repository/email"
	"github.com/abiyyu03/siruta/repository/member"
	"github.com/abiyyu03/siruta/repository/register"
	"github.com/abiyyu03/siruta/repository/user"
	"gorm.io/gorm"
)

type RTProfileRegisterRepository struct {
	rtProfileRepository     *RTProfileRepository
	rtNotification          *email.EmailRegistrationRepository
	tokenRegisterRepository *register.RegTokenRepository
	userRepository          *user.UserRepository
	memberRepository        *member.MemberRepository
}

func (r *RTProfileRegisterRepository) Register(rtProfile *model.RTProfile, referalCode string) (*model.RTProfile, error) {
	tx := config.DB.Begin()

	defer func() {
		if rec := recover(); rec != nil {
			tx.Rollback()
		}
	}()

	isCodeValid, rwProfileId, err := r.GetAndVerifyRWReferalCode(referalCode)

	rtProfile.RWProfileId = rwProfileId

	if !isCodeValid {
		tx.Rollback()
		return nil, err
	}

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	_, err = r.CreateRTProfile(tx, rtProfile, rwProfileId)

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return rtProfile, nil
}

func (r *RTProfileRegisterRepository) CreateRTProfile(tx *gorm.DB, rtProfile *model.RTProfile, RWProfileId string) (*model.RTProfile, error) {
	if err := tx.Create(&rtProfile).Error; err != nil {
		log.Printf("failed to create RTProfile: %v", err)
		return nil, err
	}

	return rtProfile, nil
}

func (r *RTProfileRegisterRepository) GetAndVerifyRWReferalCode(inputedReferalCode string) (bool, string, error) {
	var referalCode *model.ReferalCode

	if err := config.DB.Where("code = ? ", inputedReferalCode).Where("expired_at > ?", time.Now()).First(&referalCode).Error; err != nil {
		return false, "", err
	}

	if referalCode == nil {
		return false, "", errors.New("kode referal tidak valid")
	}

	return true, referalCode.ProfileId, nil
}

func (r *RTProfileRegisterRepository) CheckRTNumberAvailability(rtProfile *model.RTProfile, rtNumber string) (bool, error) {
	if err := config.DB.Where("rt_number = ?", rtNumber).First(&rtProfile).Error; err != nil {
		return false, err
	}

	return true, nil
}

func (r *RTProfileRegisterRepository) ApproveRegistrant(rtProfileId string) error {
	err := config.DB.Transaction(func(tx *gorm.DB) error {
		var rtProfile *model.RTProfile

		rt, err := r.rtProfileRepository.FetchById(rtProfileId)

		if err != nil {
			return err
		}

		err = r.rtNotification.RtNotification(rt.RTEmail)

		if err != nil {
			return err
		}

		err = r.UpdateRTAuthorization(tx, rtProfileId, rtProfile)

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		log.Printf("Transaction failed, rolled back due to error: %v", err)
		return err
	}

	log.Println("Transaction committed successfully.")

	return nil

}

func (r *RTProfileRegisterRepository) UpdateRTAuthorization(tx *gorm.DB, id string, rtProfile *model.RTProfile) error {
	if err := tx.Model(&rtProfile).Where("id = ?", id).Update("is_authorized", true).Error; err != nil {
		return err
	}

	return nil
}

func (r *RTProfileRegisterRepository) RegisterUserRt(memberData *model.Member, user *model.User, roleId uint, token string) error {
	err := config.DB.Transaction(func(tx *gorm.DB) error {

		_, err := r.userRepository.RegisterUser(tx, user, roleId)

		if err != nil {
			return err
		}

		_, err = r.tokenRegisterRepository.RemoveToken(tx, token)

		if err != nil {
			return err
		}

		_, err = r.memberRepository.Store(tx, memberData)

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		log.Printf("Transaction failed, rolled back due to error: %v", err)
		return err
	}

	log.Println("Transaction committed successfully.")

	return nil
}
