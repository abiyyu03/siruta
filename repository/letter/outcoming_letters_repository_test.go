package letter_test

import (
	"testing"

	"github.com/abiyyu03/siruta/config"
	"github.com/abiyyu03/siruta/entity/model"
	"github.com/abiyyu03/siruta/repository/letter"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB() {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	config.DB = db

	// Migrate required models
	_ = db.AutoMigrate(
		&model.OutcomingLetter{},
		&model.Member{},
		&model.LetterType{},
		&model.RTProfile{},
	)
	// Seed letters
	db.Create(&model.OutcomingLetter{
		ID:           "s001",
		LetterNumber: 22,
		Date:         "20-20-2022",
		MemberId:     "m001",
		LetterTypeId: 2,
		RTProfileId:  "rt01",
		IsRTApproved: true,
	})
}

func TestOutcomingLetterRepository_Fetch(t *testing.T) {
	setupTestDB()
	repo := &letter.OutcomingLetterRepository{}

	letters, err := repo.Fetch()
	assert.NoError(t, err)
	assert.Len(t, letters, 1)
	assert.Equal(t, 22, letters[0].LetterNumber)
}

func TestOutcomingLetterRepository_FetchById(t *testing.T) {
	setupTestDB()
	repo := &letter.OutcomingLetterRepository{}

	letter, err := repo.FetchById("s001")
	assert.NoError(t, err)
	assert.NotNil(t, letter)
	assert.Equal(t, 22, letter.LetterNumber)
}

func TestOutcomingLetterRepository_FetchByRTProfileId(t *testing.T) {
	setupTestDB()
	repo := &letter.OutcomingLetterRepository{}

	letters, err := repo.FetchByRTProfileId("rt01")
	assert.NoError(t, err)
	assert.Len(t, letters, 1)
}

func TestOutcomingLetterRepository_FetchPreview(t *testing.T) {
	setupTestDB()
	repo := &letter.OutcomingLetterRepository{}

	letter, err := repo.FetchPreview("s001")
	assert.NoError(t, err)
	assert.NotNil(t, letter)
	assert.Equal(t, true, letter.IsRTApproved)
}

func TestOutcomingLetterRepository_Delete(t *testing.T) {
	setupTestDB()
	repo := &letter.OutcomingLetterRepository{}

	err := repo.Delete("s001")
	assert.NoError(t, err)

	// Make sure it's deleted
	var count int64
	config.DB.Model(&model.OutcomingLetter{}).Where("id = ?", "s001").Count(&count)
	assert.Equal(t, int64(0), count)
}
