package letter

import (
	"log"

	"github.com/abiyyu03/siruta/config"
	"github.com/abiyyu03/siruta/entity/model"
	"gorm.io/gorm"
)

type LetterReqRepository struct{}

func (o *LetterReqRepository) StoreOutcomingLetter(outcomingLetter *model.OutcomingLetter) (*model.OutcomingLetter, error) {
	if err := config.DB.Create(&outcomingLetter).Error; err != nil {
		return nil, err
	}

	return outcomingLetter, nil
}

func (o *LetterReqRepository) CheckMemberResidentExist(birthDate string, nik string) (bool, error) {
	var member *model.Member
	var countData int64

	isMemberExist := config.DB.Model(&member).Where("nik_number = ?", nik).Where("birth_date = ?", birthDate).Count(&countData)

	if isMemberExist.Error != nil {
		return false, isMemberExist.Error
	}

	return countData > 0, nil
}

func (o *LetterReqRepository) StoreOutcomingLetterWithGuest(outcomingLetter *model.OutcomingLetter, member *model.Member) (*model.OutcomingLetter, error) {
	tx := config.DB.Begin()

	defer func() {
		if rec := recover(); rec != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Create(&outcomingLetter).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Create(&member).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return outcomingLetter, nil
}

func (o *LetterReqRepository) UpdateApprovalStatusByRT(id string) (bool, error) {
	var outcomingLetter model.OutcomingLetter

	tx := config.DB.Begin()

	defer func() {
		if rec := recover(); rec != nil {
			tx.Rollback()
		}
	}()

	outcoming := config.DB.Model(&outcomingLetter).Where("id = ?", id)

	if err := outcoming.Update("is_rt_approved", true).Error; err != nil {
		tx.Rollback()
		return false, err
	}

	outcoming.First(&outcomingLetter)

	number, err := o.checkNumberAvailability(tx, &outcomingLetter)

	if err != nil {
		tx.Rollback()
		return false, err
	}

	err = o.incrementNumberOfLetter(tx, &outcomingLetter, number, outcomingLetter.ID)

	if err != nil {
		tx.Rollback()
		return false, err
	}

	tx.Commit()

	return true, nil
}

func (o *LetterReqRepository) checkNumberAvailability(tx *gorm.DB, outcomingLetter *model.OutcomingLetter) (int64, error) {
	var latestNumber int64

	if err := tx.Model(&outcomingLetter).Where(
		"letter_type_id = ?", outcomingLetter.LetterTypeId,
	).Where(
		"rt_profile_id = ?", outcomingLetter.RTProfileId,
	).Select(
		"MAX(letter_number)",
	).Find(&latestNumber).Error; err != nil {
		return 0, err
	}

	return latestNumber, nil
}

func (o *LetterReqRepository) incrementNumberOfLetter(tx *gorm.DB, outcomingLetter *model.OutcomingLetter, latestNumber int64, id string) error {
	latestNumber = latestNumber + 1
	log.Print("updated number : ", latestNumber)

	if err := tx.Model(&outcomingLetter).Where("id = ?", id).Update("letter_number", latestNumber).Error; err != nil {
		return err
	}

	return nil
}
