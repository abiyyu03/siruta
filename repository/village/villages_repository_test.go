package village

import (
	"errors"
	"testing"

	"github.com/abiyyu03/siruta/config"
	"github.com/abiyyu03/siruta/entity/model"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var repo VillageRepository

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)

	err = db.AutoMigrate(&model.Village{})
	assert.NoError(t, err)

	return db
}

func TestFetch(t *testing.T) {
	db := setupTestDB(t)

	// Seed
	db.Create(&model.Village{Name: "Desa A", AltName: "Desa A1", Latitude: -6.2, Longitude: 106.8, CodePostal: "12345"})
	db.Create(&model.Village{Name: "Desa B", AltName: "Desa B1", Latitude: -6.3, Longitude: 106.9, CodePostal: "67890"})

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

	data := model.Village{Name: "Desa C", AltName: "Alt C", Latitude: 1.1, Longitude: 2.2, CodePostal: "54321"}
	db.Create(&data)

	oldDB := config.DB
	config.DB = db
	defer func() { config.DB = oldDB }()

	result, err := repo.FetchById(1)
	assert.NoError(t, err)
	assert.Equal(t, "Desa C", result.Name)
}

func TestStore(t *testing.T) {
	db := setupTestDB(t)

	oldDB := config.DB
	config.DB = db
	defer func() { config.DB = oldDB }()
	data := &model.Village{Name: "Desa Store", AltName: "DStore", Latitude: 1.2, Longitude: 3.4, CodePostal: "77777"}
	result, err := repo.Store(data)
	assert.NoError(t, err)
	assert.NotZero(t, result.ID)
}

func TestUpdate(t *testing.T) {
	db := setupTestDB(t)

	data := &model.Village{ID: 1, Name: "Old Name", AltName: "OldAlt", Latitude: 0.0, Longitude: 0.0, CodePostal: "00000"}
	db.Create(&data)

	updated := &model.Village{Name: "New Name", AltName: "NewAlt", Latitude: 1.1, Longitude: 2.2, CodePostal: "99999"}

	result, err := repo.Update(updated, data.ID)
	assert.NoError(t, err)

	var check *model.Village
	db.First(&check, data.ID)
	assert.Equal(t, "New Name", check.Name)
	assert.Equal(t, result.Name, check.Name)
}

func TestDelete(t *testing.T) {
	db := setupTestDB(t)

	data := &model.Village{Name: "ToDelete", AltName: "Del", Latitude: 5.5, Longitude: 6.6, CodePostal: "40404"}
	db.Create(&data)

	oldDB := config.DB
	config.DB = db
	defer func() { config.DB = oldDB }()

	_, err := repo.Delete(int(data.ID))
	assert.NoError(t, err)

	var check model.Village
	tx := db.First(&check, data.ID)
	assert.True(t, errors.Is(tx.Error, gorm.ErrRecordNotFound))
}
