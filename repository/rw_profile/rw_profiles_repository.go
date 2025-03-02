package rw_profile

import (
	"github.com/abiyyu03/siruta/config"
	"github.com/abiyyu03/siruta/entity/model"
)

type RWProfileRepository struct{}

func (r *RWProfileRepository) Fetch() ([]*model.RWProfile, error) {
	var rwProfiles []*model.RWProfile

	if err := config.DB.Find(&rwProfiles).Error; err != nil {
		return nil, err
	}

	return rwProfiles, nil
}

func (r *RWProfileRepository) FetchById(id string) (*model.RWProfile, error) {
	var rwProfile *model.RWProfile

	if err := config.DB.Where("id = ?", id).First(&rwProfile).Error; err != nil {
		return nil, err
	}

	return rwProfile, nil
}

// func Update() error {

// }

// func Delete() error {

// }

// func Store() error {

// }
