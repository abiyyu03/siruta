package auth

import (
	"log"

	"github.com/abiyyu03/siruta/config"
	"github.com/abiyyu03/siruta/entity"
	"github.com/abiyyu03/siruta/entity/model"
	"golang.org/x/crypto/bcrypt"
)

type AuthRepository struct{}

func (u *AuthRepository) FetchLogin(email, password string) (resp *entity.LoginRepositoryResponse, err error) {
	var user *model.User

	if err = config.DB.Preload("Role").Where("is_authorized", true).Where("email = ? ", email).First(&user).Error; err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, err
	}

	log.Print(user.ID)
	if user.RoleID == 1 {
		resp = &entity.LoginRepositoryResponse{
			FullName: "Super Administrator",
			Email:    user.Email,
			RoleName: user.Role.Name,
			RoleID:   user.Role.ID,
		}
	} else if user.RoleID == 2 {
		var profile *model.RWLeader
		if err = config.DB.Where("user_id = ?", user.ID).First(&profile).Error; err != nil {
			return nil, nil
		}
		resp = &entity.LoginRepositoryResponse{
			FullName: profile.Fullname,
			Email:    user.Email,
			RoleName: user.Role.Name,
			RoleID:   user.Role.ID,
		}
	} else if user.RoleID == 3 {
		var profile *model.RTLeader
		if err = config.DB.Where("user_id = ?", user.ID).First(&profile).Error; err != nil {
			return nil, nil
		}
		resp = &entity.LoginRepositoryResponse{
			FullName: profile.Fullname,
			Email:    user.Email,
			RoleName: user.Role.Name,
			RoleID:   user.Role.ID,
		}
	} else if user.RoleID == 4 {
		var profile *model.Member
		if err = config.DB.Where("user_id = ?", user.ID).First(&profile).Error; err != nil {
			return nil, nil
		}
		resp = &entity.LoginRepositoryResponse{
			FullName: profile.Fullname,
			Email:    user.Email,
			RoleName: user.Role.Name,
			RoleID:   user.Role.ID,
		}
	}

	return resp, nil
}
