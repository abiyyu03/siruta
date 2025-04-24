package rt_profile

import (
	"github.com/abiyyu03/siruta/config"
	"github.com/abiyyu03/siruta/entity/model"
	"gorm.io/gorm"
)

type RTLeaderRepository struct{}

func (r *RTLeaderRepository) Fetch() ([]*model.RTLeader, error) {
	var rtLeaders []*model.RTLeader

	if err := config.DB.Preload("RTProfile").Find(&rtLeaders).Error; err != nil {
		return nil, err
	}

	return rtLeaders, nil
}

func (r *RTLeaderRepository) FetchByRTProfileId(rtProfileId string) ([]*model.RTLeader, error) {
	var rtLeader []*model.RTLeader

	if err := config.DB.Preload("RTProfile").Where("rt_profile_id =?", rtProfileId).Find(&rtLeader).Error; err != nil {
		return nil, err
	}

	return rtLeader, nil
}

func (r *RTLeaderRepository) FetchById(id string) (*model.RTLeader, error) {
	var rtLeader model.RTLeader

	if err := config.DB.Where("id = ?", id).First(&rtLeader).Error; err != nil {
		return nil, err
	}

	return &rtLeader, nil
}

func (r *RTLeaderRepository) Store(tx *gorm.DB, rtLeader *model.RTLeader) error {
	if err := tx.Create(&rtLeader).Error; err != nil {
		return err
	}

	return nil
}

func (r *RTLeaderRepository) Update(tx *gorm.DB, rtLeader *model.RTLeader, id string) error {
	if err := tx.Where("id =?", id).Updates(&rtLeader).Error; err != nil {
		return err
	}

	return nil
}
