package religion

import (
	"errors"
	"testing"

	"github.com/abiyyu03/siruta/config"
	"github.com/abiyyu03/siruta/entity/model"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var repo *ReligionRepository

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)

	err = db.AutoMigrate(&model.Religion{})
	assert.NoError(t, err)

	return db
}

func TestFetch(t *testing.T) {
	db := setupTestDB(t)

	// Seed data
	db.Create(&model.Religion{ID: 1, ReligionName: "Islam"})
	db.Create(&model.Religion{ID: 2, ReligionName: "Non-Islam"})

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

	db.Create(&model.Religion{ID: 1, ReligionName: "Islam"})

	oldDB := config.DB
	config.DB = db
	defer func() { config.DB = oldDB }()

	result, err := repo.FetchById(1)
	assert.NoError(t, err)
	assert.Equal(t, "Islam", result.ReligionName)
}

func TestStore(t *testing.T) {
	db := setupTestDB(t)

	oldDB := config.DB
	config.DB = db
	defer func() { config.DB = oldDB }()

	inputReligion := &model.Religion{ReligionName: "Konghucu"}

	result, err := repo.Store(inputReligion)
	assert.NoError(t, err)
	assert.NotZero(t, result.ID)
}

func TestUpdate(t *testing.T) {
	db := setupTestDB(t)

	status := &model.Religion{ReligionName: "Old"}
	db.Create(status)

	oldDB := config.DB
	config.DB = db
	defer func() { config.DB = oldDB }()

	status.ReligionName = "Updated"
	updatedReligion, err := repo.Update(status, int(status.ID))
	assert.NoError(t, err)

	var check model.Religion
	db.First(&check, status.ID)
	assert.Equal(t, "Updated", check.ReligionName)
	assert.Equal(t, updatedReligion.ReligionName, "Updated")
}

func TestDelete(t *testing.T) {
	db := setupTestDB(t)

	religion := &model.Religion{ReligionName: "ToDelete"}
	db.Create(religion)

	oldDB := config.DB
	config.DB = db
	defer func() { config.DB = oldDB }()

	_, err := repo.Delete(int(religion.ID))
	assert.NoError(t, err)

	var check model.Religion
	tx := db.First(&check, religion.ID)
	assert.True(t, errors.Is(tx.Error, gorm.ErrRecordNotFound))
}
