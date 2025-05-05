package register

import (
	"testing"
	"time"

	"github.com/abiyyu03/siruta/config"
	"github.com/abiyyu03/siruta/entity/model"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SetupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect to test database: %v", err)
	}

	// Migrate tabel yang akan dipakai di test
	err = db.AutoMigrate(&model.RegistrationToken{})
	if err != nil {
		t.Fatalf("failed to migrate database: %v", err)
	}

	return db
}

func TestCreateToken_Success(t *testing.T) {
	db := SetupTestDB(t)
	config.DB = db
	repo := RegTokenRepository{}

	token := "abc123"
	tokenResult, err := repo.CreateToken(token)

	assert.Nil(t, err)
	assert.Equal(t, token, tokenResult)
}

func TestValidateToken_Success(t *testing.T) {
	db := SetupTestDB(t)
	config.DB = db
	repo := RegTokenRepository{}
	token := "valid-token"

	// Insert token aktif
	err := db.Create(&model.RegistrationToken{
		Token:     token,
		ExpiredAt: time.Now().Add(1 * time.Hour),
	}).Error
	assert.Nil(t, err)

	result, err := repo.Validate(token)
	assert.Nil(t, err)
	assert.Equal(t, token, result)
}

func TestValidateToken_Expired(t *testing.T) {
	db := SetupTestDB(t)
	config.DB = db
	repo := RegTokenRepository{}
	token := "expired-token"

	// Insert token kadaluarsa
	_ = db.Create(&model.RegistrationToken{
		Token:     token,
		ExpiredAt: time.Now().Add(-1 * time.Hour),
	}).Error

	result, err := repo.Validate(token)
	assert.Nil(t, err)
	assert.Empty(t, result)
}

func TestRemoveToken_Success(t *testing.T) {
	SetupTestDB(t) // <- langsung panggil ini
	repo := RegTokenRepository{}
	token := "remove-me"

	// Insert token aktif
	_ = config.DB.Create(&model.RegistrationToken{
		Token:     token,
		ExpiredAt: time.Now().Add(1 * time.Hour),
	}).Error

	tx := config.DB.Begin()
	defer tx.Rollback()

	deleted, err := repo.RemoveToken(tx, token)
	tx.Commit()
	assert.Nil(t, err)
	assert.True(t, deleted)
}

func TestRemoveToken_NotFound(t *testing.T) {
	db := SetupTestDB(t)
	config.DB = db
	repo := RegTokenRepository{}
	token := "not-exist"

	// Insert token aktif
	err := db.Create(&model.RegistrationToken{
		Token:     token,
		ExpiredAt: time.Now().Add(1 * time.Hour),
	}).Error

	tx := db.Begin()
	defer tx.Rollback()

	deleted, err := repo.RemoveToken(tx, token)

	tx.Commit()
	assert.Nil(t, err)
	assert.False(t, deleted)
}
