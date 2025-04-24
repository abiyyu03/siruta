package referal_code

import (
	"testing"
	"time"

	"github.com/abiyyu03/siruta/config"
	"github.com/abiyyu03/siruta/entity/model"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db = config.DB
var repo *ReferalCodeRepository

func SetDB(testDB *gorm.DB) *gorm.DB {
	old := db
	db = testDB
	return old
}

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)

	err = db.AutoMigrate(&model.ReferalCode{})
	assert.NoError(t, err)

	return db
}

func TestFetch(t *testing.T) {
	db := setupTestDB(t)
	repo := ReferalCodeRepository{}

	// Seed data
	db.Create(&model.ReferalCode{Code: "ABC123", ExpiredAt: time.Now().Add(24 * time.Hour), ProfileId: "uuid-profile-1"})

	// Override global config.DB with test DB temporarily
	oldDB := config.DB
	config.DB = db
	defer func() { config.DB = oldDB }()

	result, err := repo.Fetch()

	assert.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, "ABC123", result[0].Code)
}

func TestFetchById_NotFound(t *testing.T) {
	db := setupTestDB(t)
	repo := ReferalCodeRepository{}

	// Inject test DB
	oldDB := config.DB
	config.DB = db
	defer func() { config.DB = oldDB }()

	res, err := repo.FetchById("non-existent-id")
	assert.Error(t, err)
	assert.Nil(t, res)
}

func TestGenerateReferalCode(t *testing.T) {
	db := setupTestDB(t)
	repo := ReferalCodeRepository{}

	tx := db.Begin()
	defer tx.Rollback()

	err := repo.GenerateReferalCode(tx, "uuid-profile-123")
	assert.NoError(t, err)

	tx.Commit()

	var count int64
	db.Model(&model.ReferalCode{}).Count(&count)
	assert.Equal(t, int64(1), count)
}

func TestValidate(t *testing.T) {
	db := setupTestDB(t)
	repo := ReferalCodeRepository{}

	now := time.Now()
	db.Create(&model.ReferalCode{
		Code:      "VALID123",
		ExpiredAt: now.Add(1 * time.Hour),
		ProfileId: "valid-profile-id",
	})

	// Inject test DB
	oldDB := config.DB
	config.DB = db
	defer func() { config.DB = oldDB }()

	profileId, isValid, err := repo.Validate("VALID123")

	assert.NoError(t, err)
	assert.True(t, isValid)
	assert.Equal(t, "valid-profile-id", profileId)
}

func TestGetAndVerifyRWReferalCode_Invalid(t *testing.T) {
	db := setupTestDB(t)
	repo := ReferalCodeRepository{}

	// Inject test DB
	oldDB := config.DB
	config.DB = db
	defer func() { config.DB = oldDB }()

	isValid, _, err := repo.GetAndVerifyRWReferalCode("INVALID")
	assert.Error(t, err)
	assert.False(t, isValid)
}

func TestDelete(t *testing.T) {
	db := setupTestDB(t)
	repo := ReferalCodeRepository{}

	// Inject test DB
	oldDB := config.DB
	config.DB = db
	defer func() { config.DB = oldDB }()

	// Seed referal code
	referal := &model.ReferalCode{
		Code:      "DELETE1",
		ExpiredAt: time.Now().Add(24 * time.Hour),
		ProfileId: "profile-delete-test",
	}
	db.Create(&referal)

	// Get the ID
	var stored model.ReferalCode
	db.Where("code = ?", "DELETE1").First(&stored)

	// Perform delete
	err := repo.Delete(int(stored.ID))
	assert.NoError(t, err)

	// Confirm it no longer exists
	var count int64
	db.Model(&model.ReferalCode{}).Where("id = ?", stored.ID).Count(&count)
	assert.Equal(t, int64(0), count)
}
