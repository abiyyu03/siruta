package letter_test

import (
	"testing"
	"time"

	"github.com/abiyyu03/siruta/config"
	"github.com/abiyyu03/siruta/entity/model"
	"github.com/abiyyu03/siruta/repository/letter"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupIncomingLetterRepoDB() {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	config.DB = db
	// Seed database, sebaiknya menggunakan factory atau gorm untuk menyiapkan data awal
	config.DB.AutoMigrate(&model.IncomingLetter{})
}

func TestStoreIncomingLetter(t *testing.T) {
	setupIncomingLetterRepoDB()
	repo := &letter.IncomingLetterRepository{}

	now := time.Now()

	// Data Incoming Letter
	incomingLetter := &model.IncomingLetter{
		Title:        "Surat Permohonan",
		LetterDate:   now,                // Format string
		OriginLetter: "Surat dari RT 01", // Surat Asal
		RTProfileId:  "rt01",
	}

	// Store Incoming Letter
	result, err := repo.Store(incomingLetter)

	// Validasi
	assert.NoError(t, err)
	assert.NotNil(t, result)

	// Validasi surat tersimpan di database
	var storedLetter model.IncomingLetter
	config.DB.First(&storedLetter, "title = ?", "Surat Permohonan")
	assert.Equal(t, "Surat Permohonan", storedLetter.Title)
	assert.Equal(t, "rt02", storedLetter.LetterDate)
}

func TestFetchIncomingLetters(t *testing.T) {
	setupIncomingLetterRepoDB()
	repo := &letter.IncomingLetterRepository{}

	now := time.Now()

	// Seed data untuk Incoming Letter
	letter := &model.IncomingLetter{
		Title:        "Surat Pemberitahuan",
		LetterDate:   now,                // Format string
		OriginLetter: "Surat dari RT 02", // Surat Asal
		RTProfileId:  "rt02",
	}

	config.DB.Create(&letter)

	// Fetch all Incoming Letters
	result, err := repo.Fetch()

	// Validasi
	assert.NoError(t, err)
	assert.Len(t, result, 1) // Pastikan ada 1 data

	// Validasi jika data surat yang sesuai ada
	assert.Equal(t, "Surat Pemberitahuan", result[0].Title)
	assert.Equal(t, "rt02", result[0].RTProfileId)
}

func TestFetchIncomingLetterById(t *testing.T) {
	setupIncomingLetterRepoDB()
	repo := &letter.IncomingLetterRepository{}

	now := time.Now()
	// Seed data untuk Incoming Letter
	letter := &model.IncomingLetter{
		Title:        "Surat Undangan",
		LetterDate:   now,                // Format string
		OriginLetter: "Surat dari RT 03", // Surat Asal
		RTProfileId:  "rt03",
	}

	config.DB.Create(&letter)

	// Fetch Incoming Letter by ID
	result, err := repo.FetchById(1)

	// Validasi
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "Surat Undangan", result.Title)
	assert.Equal(t, "rt02", result.LetterDate)
}

func TestFetchIncomingLetterByRTProfileId(t *testing.T) {
	setupIncomingLetterRepoDB()
	repo := &letter.IncomingLetterRepository{}

	now := time.Now()
	// Seed data untuk Incoming Letter
	letter := &model.IncomingLetter{
		Title:        "Surat Keterangan",
		LetterDate:   now,                // Format string
		OriginLetter: "Surat dari RT 04", // Surat Asal
		RTProfileId:  "rt04",             // RTProfileId yang sesuai
	}

	config.DB.Create(&letter)

	// Fetch Incoming Letter by RTProfileId
	result, err := repo.FetchByRTProfileId("rt04")

	// Validasi
	assert.NoError(t, err)
	assert.Len(t, result, 1) // Pastikan ada 1 data
	assert.Equal(t, "Surat Keterangan", result[0].Title)
	assert.Equal(t, "rt04", result[0].RTProfileId)
}

func TestUpdateIncomingLetter(t *testing.T) {
	setupIncomingLetterRepoDB()
	repo := &letter.IncomingLetterRepository{}

	now := time.Now()
	// Seed data untuk Incoming Letter
	letter := &model.IncomingLetter{
		Title:        "Surat Tugas",
		LetterDate:   now,                // Format string
		OriginLetter: "Surat dari RT 05", // Surat Asal
		RTProfileId:  "rt05",
	}

	config.DB.Create(&letter)

	// Update Incoming Letter
	letter.Title = "Surat Persetujuan"
	updatedLetter, err := repo.Update(letter, 1)

	// Validasi
	assert.NoError(t, err)
	assert.NotNil(t, updatedLetter)

	// Validasi jika surat yang diupdate ada dan judulnya berubah
	var storedLetter model.IncomingLetter
	config.DB.First(&storedLetter, "id = ?", 1)
	assert.Equal(t, "Surat Persetujuan", storedLetter.Title)
}

func TestDeleteIncomingLetter(t *testing.T) {
	setupIncomingLetterRepoDB()
	repo := &letter.IncomingLetterRepository{}
	now := time.Now()

	// Seed data untuk Incoming Letter
	letter := &model.IncomingLetter{
		ID:           1,
		Title:        "Surat Pembatalan",
		LetterDate:   now,                // Format string
		OriginLetter: "Surat dari RT 06", // Surat Asal
		RTProfileId:  "rt06",
	}

	config.DB.Create(&letter)

	// Delete Incoming Letter
	deletedLetter, err := repo.Delete(letter, 1)

	// Validasi
	assert.NoError(t, err)
	assert.NotNil(t, deletedLetter)

	// Pastikan data sudah dihapus dari database
	var deletedLetterCheck model.IncomingLetter
	err = config.DB.First(&deletedLetterCheck, "id = ?", 1).Error
	assert.Error(t, err)
}
