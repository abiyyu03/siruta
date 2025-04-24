package letter

import (
	"testing"

	"github.com/abiyyu03/siruta/config"
	"github.com/abiyyu03/siruta/entity/model"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupReqRepoDB() {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	config.DB = db

	config.DB.AutoMigrate(&model.OutcomingLetter{}, &model.Member{})
}

func TestStoreOutcomingLetterWithGuest(t *testing.T) {
	setupReqRepoDB()
	repo := LetterReqRepository{}

	// Data Guest
	guest := &model.Member{
		ID:             "m002",
		Fullname:       "Tamu Undangan",
		BornPlace:      "Bogor",
		BirthDate:      "1999-12-12", // Format string
		Gender:         "P",
		ReligionId:     2,
		MemberStatusId: 3,
		UserId:         "user002",
		Status:         "Aktif",
		RTProfileId:    "rt01",
	}

	// Data Letter
	letter := &model.OutcomingLetter{
		ID:           "l003",
		LetterNumber: 3,            // Number sekarang tipe int
		Date:         "2025-04-10", // Format string
		MemberId:     "m002",
		LetterTypeId: 4,
		RTProfileId:  "rt01",
		IsRTApproved: false,
	}

	// Store Outcoming Letter and Guest
	result, err := repo.StoreOutcomingLetterWithGuest(letter, guest)

	// Validasi
	assert.NoError(t, err)
	assert.NotNil(t, result)

	// Validasi jika guest ditambahkan ke database
	var member model.Member
	config.DB.First(&member, "id = ?", "m002")
	assert.Equal(t, "Tamu Undangan", member.Fullname)
}

func TestUpdateApprovalStatusByRT(t *testing.T) {
	setupReqRepoDB()
	repo := LetterReqRepository{}

	// Seed data untuk letter yang belum di-approve
	letter := &model.OutcomingLetter{
		ID:           "l004",
		LetterNumber: 4,            // Nomor surat dalam format int
		Date:         "2025-05-15", // Format string
		MemberId:     "m001",
		LetterTypeId: 4,
		RTProfileId:  "rt01",
		IsRTApproved: false,
	}

	config.DB.Create(&letter)

	// Test update approval status by RT
	success, err := repo.UpdateApprovalStatusByRT("l004")
	assert.NoError(t, err)
	assert.True(t, success)

	// Verify if 'is_rt_approved' is set to true
	var updatedLetter model.OutcomingLetter
	config.DB.First(&updatedLetter, "id = ?", "l004")
	assert.True(t, success)
}

func TestStoreOutcomingLetter(t *testing.T) {
	setupReqRepoDB()
	repo := LetterReqRepository{}

	// Data Letter
	letter := &model.OutcomingLetter{
		ID:           "l005",
		LetterNumber: 5,            // Nomor surat dalam format int
		Date:         "2025-06-01", // Format string
		MemberId:     "m001",
		LetterTypeId: 4,
		RTProfileId:  "rt01",
		IsRTApproved: false,
	}

	// Store Outcoming Letter
	result, err := repo.StoreOutcomingLetter(letter)

	// Validasi
	assert.NoError(t, err)
	assert.NotNil(t, result)

	// Validasi surat tersimpan di database
	var storedLetter model.OutcomingLetter
	config.DB.First(&storedLetter, "id = ?", "l005")
	assert.Equal(t, 5, letter.LetterNumber)
}

func TestCheckMemberResidentExist(t *testing.T) {
	setupReqRepoDB()
	repo := LetterReqRepository{}

	nik := "1234567890123456"

	// Seed data untuk member
	member := &model.Member{
		ID:             "m001",
		NikNumber:      &nik,
		Fullname:       "Agus Setiawan",
		BornPlace:      "Depok",
		BirthDate:      "2000-01-01", // Format string
		Gender:         "L",
		ReligionId:     1,
		MemberStatusId: 1,
		UserId:         "user001",
		Status:         "Aktif",
		RTProfileId:    "rt01",
	}

	// Seed member
	config.DB.Create(&member)

	// Check Member Exists
	exists, err := repo.CheckMemberResidentExist("2000-01-01", "1234567890123456")
	assert.NoError(t, err)
	assert.True(t, exists)

	// Check Member Not Exists
	exists, err = repo.CheckMemberResidentExist("2001-01-01", "0000000000000000")
	assert.NoError(t, err)
	assert.False(t, exists)
}
