package letter_type

import (
	"errors"
	"testing"

	"github.com/abiyyu03/siruta/config"
	"github.com/abiyyu03/siruta/entity/model"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var repo *LetterTypeRepository

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)

	err = db.AutoMigrate(&model.LetterType{})
	assert.NoError(t, err)

	return db
}

func TestFetch(t *testing.T) {
	db := setupTestDB(t)

	// Seed data
	db.Create(&model.LetterType{TypeName: "Domisili", Code: "DOM", IsForLocalResident: true})
	db.Create(&model.LetterType{TypeName: "Keterangan Usaha", Code: "USAH", IsForLocalResident: false})

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

	// _ = &model.LetterType{TypeName: "Kelahiran", Code: "BRTH", IsForLocalResident: true}
	db.Create(&model.LetterType{TypeName: "Domisili", Code: "DOM", IsForLocalResident: true})

	oldDB := config.DB
	config.DB = db
	defer func() { config.DB = oldDB }()

	result, err := repo.FetchById(1)
	assert.NoError(t, err)
	assert.Equal(t, "Kelahiran", result.TypeName)
}

func TestStore(t *testing.T) {
	db := setupTestDB(t)

	oldDB := config.DB
	config.DB = db
	defer func() { config.DB = oldDB }()

	newType := &model.LetterType{TypeName: "Kelahiran", Code: "BRTH", IsForLocalResident: true}

	result, err := repo.Store(newType)
	assert.NoError(t, err)
	assert.NotZero(t, result.ID)
}

func TestUpdate(t *testing.T) {
	db := setupTestDB(t)

	updated := &model.LetterType{TypeName: "Kelahiran", Code: "BRTH", IsForLocalResident: true}
	db.Create(updated)

	oldDB := config.DB
	config.DB = db
	defer func() { config.DB = oldDB }()

	updated.Code = "KEL"
	updatedType, err := repo.Update(updated, int(updated.ID))
	assert.NoError(t, err)

	var check model.LetterType
	db.First(&check, updated.ID)
	assert.Equal(t, "Kelahiran", check.TypeName)
	assert.Equal(t, updatedType.TypeName, "Kelahiran")
}

func TestDelete(t *testing.T) {
	db := setupTestDB(t)

	deleted := &model.LetterType{TypeName: "Kelahiran", Code: "BRTH", IsForLocalResident: true}
	db.Create(deleted)

	oldDB := config.DB
	config.DB = db
	defer func() { config.DB = oldDB }()

	_, err := repo.Delete(int(deleted.ID))
	assert.NoError(t, err)

	var check model.LetterType
	tx := db.First(&check, deleted.ID)
	assert.True(t, errors.Is(tx.Error, gorm.ErrRecordNotFound))
}
