package repository

import (
	"github.com/abiyyu03/siruta/config"
	"github.com/abiyyu03/siruta/entity/model"
)

type MemberStatusRepository struct{}

func (m *MemberStatusRepository) Fetch() ([]*model.MemberStatus, error) {
	var memberStatus []*model.MemberStatus

	if err := config.DB.Find(&memberStatus).Error; err != nil {
		return nil, err
	}

	return memberStatus, nil
}

func (m *MemberStatusRepository) FetchById(id int) (*model.MemberStatus, error) {
	var memberStatus *model.MemberStatus

	if err := config.DB.Where("id = ?", id).First(&memberStatus).Error; err != nil {
		return nil, err
	}

	return memberStatus, nil
}

func (m *MemberStatusRepository) Store(memberStatus *model.MemberStatus) (*model.MemberStatus, error) {
	if err := config.DB.Create(&memberStatus).Error; err != nil {
		return nil, err
	}

	return memberStatus, nil
}

func (m *MemberStatusRepository) Update(memberStatus *model.MemberStatus, id int) (*model.MemberStatus, error) {
	if err := config.DB.Where("id = ?", id).Updates(&memberStatus).Error; err != nil {
		return nil, err
	}

	return memberStatus, nil
}

func (m *MemberStatusRepository) Delete(id int) (*model.MemberStatus, error) {
	var memberStatus model.MemberStatus

	if err := config.DB.Where("id = ?", id).Delete(&memberStatus).Error; err != nil {
		return nil, err
	}

	return nil, nil
}
