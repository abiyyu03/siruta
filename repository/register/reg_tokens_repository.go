package register

import (
	"time"

	"github.com/abiyyu03/siruta/config"
	"github.com/abiyyu03/siruta/entity/model"
	"gorm.io/gorm"
)

type RegTokenRepository struct{}

func (r *RegTokenRepository) CreateToken(token string) (string, error) {
	now := time.Now()

	if err := config.DB.Create(&model.RegistrationToken{
		Token:     token,
		ExpiredAt: now.Add(24 * time.Hour),
	}).Error; err != nil {
		return "", err
	}

	return token, nil
}

func (r *RegTokenRepository) Validate(token string) (string, error) {
	var regisToken *model.RegistrationToken

	if err := config.DB.Where("token = ?", token).Where(" expired_at > ?", time.Now()).First(&regisToken).Error; err != nil {
		return "", nil
	}

	if regisToken == nil {
		return "", nil
	}

	return token, nil
}

func (r *RegTokenRepository) RemoveToken(tx *gorm.DB, token string) (bool, error) {
	var tokenData *model.RegistrationToken

	tokens := tx.Where("token = ?", token).First(&tokenData)

	if tokens == nil {
		return false, nil
	}

	if err := tokens.Error; err != nil {
		return false, err
	}

	deletedToken := config.DB.Where("token = ?", token).Delete(&tokenData)

	if err := deletedToken.Error; err != nil {
		return false, err
	}

	return true, nil
}
