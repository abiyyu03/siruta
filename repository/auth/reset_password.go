package auth

import (
	"time"

	"github.com/abiyyu03/siruta/config"
	"github.com/abiyyu03/siruta/entity/model"
	"github.com/abiyyu03/siruta/repository/user"
	"gorm.io/gorm"
)

type ResetPasswordRepository struct {
	userRepository *user.UserRepository
}

func (r *ResetPasswordRepository) ForgotPassword(token string) (bool, error) {
	var user *model.User

	if err := config.DB.Where("email = ?", token).First(&user).Error; err != nil {
		return false, err
	}

	return true, nil
}

func (r *ResetPasswordRepository) StoreToken(reset *model.ResetPassword) (*model.ResetPassword, error) {
	if err := config.DB.Create(&reset).Error; err != nil {
		return nil, err
	}

	return reset, nil
}

func (r *ResetPasswordRepository) VerifyAndGetResetToken(token string) (*model.ResetPassword, error) {
	var reset *model.ResetPassword

	if err := config.DB.Where("token = ?", token).Where("expired_at > ?", time.Now()).First(&reset).Error; err != nil {
		return nil, err
	}

	return reset, nil
}
func (r *ResetPasswordRepository) DeleteResetToken(tx *gorm.DB, token string) error {
	var reset model.ResetPassword

	err := tx.Where("token = ?", token).First(&reset).Error
	if err != nil {
		return err
	}

	if err := tx.Delete(&reset).Error; err != nil {
		return err
	}

	return nil
}

func (r *ResetPasswordRepository) ResetPassword(hashedPassword string, token string) error {
	tx := config.DB.Begin()

	defer func() {
		if rec := recover(); rec != nil {
			tx.Rollback()
		}
	}()

	tokenData, err := r.VerifyAndGetResetToken(token)

	if err != nil || tokenData == nil {
		return err
	}

	if err := r.userRepository.UpdatePassword(tx, tokenData.UserID, hashedPassword); err != nil {
		return err
	}
	if err := r.DeleteResetToken(tx, token); err != nil {
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}
