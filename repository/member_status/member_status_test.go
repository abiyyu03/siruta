package member_status

import (
	"errors"
	"testing"

	"github.com/abiyyu03/siruta/config"
	"github.com/abiyyu03/siruta/entity/model"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var repo *MemberStatusRepository

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)

	err = db.AutoMigrate(&model.MemberStatus{})
	assert.NoError(t, err)

	return db
}

func TestFetch(t *testing.T) {
	db := setupTestDB(t)

	// Seed data
	db.Create(&model.MemberStatus{ID: 1, Status: "Active"})
	db.Create(&model.MemberStatus{ID: 2, Status: "Inactive"})

	// Inject test DB
	oldDB := config.DB
	config.DB = db
	defer func() { config.DB = oldDB }()

	result, err := repo.Fetch()
	assert.NoError(t, err)
	assert.Len(t, result, 2)
}

func TestFetchById(t *testing.T) {
	db := setupTestDB(t)

	db.Create(&model.MemberStatus{ID: 1, Status: "Active"})

	oldDB := config.DB
	config.DB = db
	defer func() { config.DB = oldDB }()

	result, err := repo.FetchById(1)
	assert.NoError(t, err)
	assert.Equal(t, "Active", result.Status)
}

func TestStore(t *testing.T) {
	db := setupTestDB(t)

	oldDB := config.DB
	config.DB = db
	defer func() { config.DB = oldDB }()

	newStatus := &model.MemberStatus{Status: "Pending"}

	result, err := repo.Store(newStatus)
	assert.NoError(t, err)
	assert.NotZero(t, result.ID)
}

func TestUpdate(t *testing.T) {
	db := setupTestDB(t)

	status := &model.MemberStatus{Status: "Old"}
	db.Create(status)

	oldDB := config.DB
	config.DB = db
	defer func() { config.DB = oldDB }()

	status.Status = "Updated"
	updatedStatus, err := repo.Update(status, int(status.ID))
	assert.NoError(t, err)

	var check model.MemberStatus
	db.First(&check, status.ID)
	assert.Equal(t, "Updated", check.Status)
	assert.Equal(t, updatedStatus.Status, "Updated")
}

func TestDelete(t *testing.T) {
	db := setupTestDB(t)

	status := &model.MemberStatus{Status: "ToDelete"}
	db.Create(status)

	oldDB := config.DB
	config.DB = db
	defer func() { config.DB = oldDB }()

	_, err := repo.Delete(int(status.ID))
	assert.NoError(t, err)

	var check model.MemberStatus
	tx := db.First(&check, status.ID)
	assert.True(t, errors.Is(tx.Error, gorm.ErrRecordNotFound))
}
