package auth

import (
	"testing"
	"time"

	"github.com/abiyyu03/siruta/config"
	"github.com/abiyyu03/siruta/entity/model"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupInMemoryDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)

	// Migrate schema
	err = db.AutoMigrate(&model.User{}, &model.ResetPassword{})
	assert.NoError(t, err)

	// Set ke config.DB biar sama dengan yang di repo
	config.DB = db

	return db
}

func TestForgotPassword(t *testing.T) {
	db := setupInMemoryDB(t)

	// Insert user
	user := &model.User{
		ID:    "user-001",
		Email: "user@example.com",
	}
	db.Create(user)

	repo := &ResetPasswordRepository{}

	ok, err := repo.ForgotPassword("user@example.com")
	assert.NoError(t, err)
	assert.True(t, ok)
}

func TestStoreToken(t *testing.T) {
	setupInMemoryDB(t)

	repo := &ResetPasswordRepository{}

	reset := &model.ResetPassword{
		UserID:    "user-001",
		Token:     "token-123",
		ExpiredAt: time.Now().Add(1 * time.Hour),
	}

	result, err := repo.StoreToken(reset)
	assert.NoError(t, err)
	assert.Equal(t, "token-123", result.Token)
}

func TestVerifyAndGetResetToken(t *testing.T) {
	setupInMemoryDB(t)

	reset := &model.ResetPassword{
		UserID:    "user-001",
		Token:     "valid-token",
		ExpiredAt: time.Now().Add(2 * time.Hour),
	}
	config.DB.Create(&reset)

	repo := &ResetPasswordRepository{}
	found, err := repo.VerifyAndGetResetToken("valid-token")

	assert.NoError(t, err)
	assert.Equal(t, reset.Token, found.Token)
}

func TestDeleteResetToken(t *testing.T) {
	db := setupInMemoryDB(t)

	reset := &model.ResetPassword{
		UserID:    "user-001",
		Token:     "del-token",
		ExpiredAt: time.Now().Add(2 * time.Hour),
	}
	db.Create(&reset)

	repo := &ResetPasswordRepository{}
	err := repo.DeleteResetToken(db, "del-token")

	assert.NoError(t, err)

	var count int64
	db.Model(&model.ResetPassword{}).Where("token = ?", "del-token").Count(&count)
	assert.Equal(t, int64(0), count)
}
