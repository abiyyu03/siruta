package member

import (
	"github.com/abiyyu03/siruta/config"
	"github.com/abiyyu03/siruta/entity/model"
	"gorm.io/gorm"
)

type MemberRepository struct{}

func (m *MemberRepository) Fetch() ([]*model.Member, error) {
	var members []*model.Member

	if err := config.DB.Find(&members).Error; err != nil {
		return nil, err
	}

	return members, nil
}

func (m *MemberRepository) FetchById(id string) (*model.Member, error) {
	var member *model.Member

	if err := config.DB.Where("id = ?", id).First(&member).Error; err != nil {
		return nil, err
	}

	return member, nil
}

func (m *MemberRepository) FetchByRTProfileId(rtProfileId string) ([]*model.Member, error) {
	var members []*model.Member

	if err := config.DB.Where("rt_profile_id =?", rtProfileId).Find(&members).Error; err != nil {
		return nil, err
	}

	return members, nil
}

func (m *MemberRepository) Store(tx *gorm.DB, memberData *model.Member) (*model.Member, error) {
	if err := tx.Create(&memberData).Error; err != nil {
		return nil, err
	}

	return memberData, nil
}

func (m *MemberRepository) Update(memberData *model.Member, id string) (*model.Member, error) {
	if err := config.DB.Where("id = ?", id).Updates(&memberData).Error; err != nil {
		return nil, err
	}

	return memberData, nil
}
func (m *MemberRepository) Delete(memberData *model.Member, id string) (*model.Member, error) {
	if err := config.DB.Where("id = ?", id).Delete(&memberData).Error; err != nil {
		return nil, err
	}

	return memberData, nil
}
