package rt_profile

import (
	"github.com/abiyyu03/siruta/config"
	"github.com/abiyyu03/siruta/entity/model"
)

type RTProfileRepository struct{}

func (r *RTProfileRepository) Fetch() ([]*model.RTProfile, error) {
	var rtProfiles []*model.RTProfile

	if err := config.DB.Find(&rtProfiles).Error; err != nil {
		return rtProfiles, nil
	}

	return rtProfiles, nil
}

func (r *RTProfileRepository) FetchById(id string) (*model.RTProfile, error) {
	var rtProfile *model.RTProfile

	if err := config.DB.Where("id = ?", id).First(&rtProfile).Error; err != nil {
		return nil, err
	}

	return rtProfile, nil
}
