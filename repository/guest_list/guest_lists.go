package guest_list

import (
	"github.com/abiyyu03/siruta/config"
	"github.com/abiyyu03/siruta/entity/model"
)

type GuestListRepository struct{}

func (g *GuestListRepository) Fetch() ([]*model.GuestList, error) {
	var guestLists []*model.GuestList

	if err := config.DB.Find(&guestLists).Error; err != nil {
		return nil, err
	}

	return guestLists, nil
}

func (g *GuestListRepository) FetchById(id int) (*model.GuestList, error) {
	var guestList *model.GuestList

	if err := config.DB.Where("id = ?", id).First(&guestList).Error; err != nil {
		return nil, err
	}

	return guestList, nil
}

func (g *GuestListRepository) FetchByRTProfileId(rtProfileId string) ([]*model.GuestList, error) {
	var guestLists []*model.GuestList

	if err := config.DB.Where("rt_profile_id = ?", rtProfileId).Find(&guestLists).Error; err != nil {
		return nil, err
	}

	return guestLists, nil
}

func (g *GuestListRepository) Store(guestListData *model.GuestList) (*model.GuestList, error) {
	if err := config.DB.Create(&guestListData).Error; err != nil {
		return nil, err
	}

	return guestListData, nil
}

func (g *GuestListRepository) Update(guestListData *model.GuestList, id int) (*model.GuestList, error) {
	if err := config.DB.Where("id =?", id).Updates(&guestListData).Error; err != nil {
		return nil, err
	}

	return guestListData, nil
}

func (g *GuestListRepository) Delete(id int) (*model.GuestList, error) {
	var guestList *model.GuestList

	if err := config.DB.Where("id =?", id).Delete(&guestList).Error; err != nil {
		return nil, err
	}

	return guestList, nil
}
