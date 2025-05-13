package rw_profile

import (
	"github.com/abiyyu03/siruta/config"
	"github.com/abiyyu03/siruta/entity/model"
	"gorm.io/gorm"
)

type RWLeaderRepository struct{}

func (r *RWLeaderRepository) Fetch() ([]*model.RWLeader, error) {
	var rwLeaders []*model.RWLeader

	if err := config.DB.Preload("RWProfile").Find(&rwLeaders).Error; err != nil {
		return nil, err
	}

	return rwLeaders, nil
}

func (r *RWLeaderRepository) FetchById(id string) (*model.RWLeader, error) {
	var rwLeader *model.RWLeader

	if err := config.DB.Where("id =?", id).First(&rwLeader).Error; err != nil {
		return nil, err
	}

	return rwLeader, nil
}

func (r *RWLeaderRepository) Store(tx *gorm.DB, rwLeader *model.RWLeader) error {
	if err := tx.Create(&rwLeader).Error; err != nil {
		return err
	}

	return nil
}

func (r *RWLeaderRepository) Update(tx *gorm.DB, rwLeader *model.RWLeader, id string) error {
	if err := tx.Where("id =?", id).Updates(&rwLeader).Error; err != nil {
		return err
	}
	return nil
}
