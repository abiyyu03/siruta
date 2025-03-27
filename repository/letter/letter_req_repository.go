package letter

import (
	"github.com/abiyyu03/siruta/config"
	"github.com/abiyyu03/siruta/entity/model"
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

func (o *LetterReqRepository) UpdateApprovalStatusByRT(outcomingLetter *model.OutcomingLetter, id string) (bool, error) {
	outcoming := config.DB.Model(&outcomingLetter).Where("id = ?", id)

	if err := outcoming.Update("is_rt_approved", true).Error; err != nil {
		return false, err
	}

	number, err := o.checkNumberAvailability(outcomingLetter)

	if err != nil {
		return false, err
	}

	err = o.incrementingNumberLetter(outcomingLetter, number, outcomingLetter.ID)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (o *LetterReqRepository) checkNumberAvailability(outcomingLetter *model.OutcomingLetter) (int64, error) {
	var latestNumber int64

	if err := config.DB.Where(
		"letter_type_id = ?", outcomingLetter.LetterTypeId,
	).Where(
		"rt_profile_id = ?", outcomingLetter.RTProfileId,
	).Select(
		"MAX(number_letter)",
	).Scan(&latestNumber).Error; err != nil {
		return 0, err
	}

	return latestNumber, nil
}

func (o *LetterReqRepository) incrementingNumberLetter(outcomingLetter *model.OutcomingLetter, latestNumber int64, id string) error {
	latestNumber = latestNumber + 1

	if err := config.DB.Model(&outcomingLetter).Where("id = ?", id).Update("number_letter", latestNumber).Error; err != nil {
		return err
	}

	return nil
}

// func (o *LetterReqRepository) CheckYearLetter(id string) {}
