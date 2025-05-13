package inventory

import (
	"errors"
	"testing"

	"github.com/abiyyu03/siruta/config"
	"github.com/abiyyu03/siruta/entity/model"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var repo *InventoryRepository

func setup(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)
	assert.NoError(t, db.AutoMigrate(&model.Inventory{}, &model.RTProfile{}))

	config.DB = db
	repo = &InventoryRepository{}
	return db
}

func TestInventory_Store(t *testing.T) {
	db := setup(t)

	// Seed RTProfile
	rt := model.RTProfile{RTNumber: "RT 02/18"}
	db.Create(&rt)

	image := "meja.jpg"

	input := &model.Inventory{Name: "Kursi", Quantity: 20, Image: &image, RTProfileId: rt.ID}
	result, err := repo.Store(input)

	assert.NoError(t, err)
	assert.NotZero(t, result.ID)
	assert.Equal(t, "Kursi", result.Name)
}

func TestInventory_Fetch(t *testing.T) {
	db := setup(t)
	rt := model.RTProfile{RTNumber: "RT 01"}
	db.Create(&rt)

	image := "meja.jpg"

	db.Create(&model.Inventory{Name: "Meja", Quantity: 10, Image: &image, RTProfileId: rt.ID})
	db.Create(&model.Inventory{Name: "Tenda", Quantity: 5, Image: &image, RTProfileId: rt.ID})

	result, err := repo.Fetch()
	assert.NoError(t, err)
	assert.Len(t, result, 2)
}

func TestInventory_Update(t *testing.T) {
	db := setup(t)
	rt := model.RTProfile{RTNumber: "RT 03"}
	db.Create(&rt)

	data := &model.Inventory{Name: "Lampu", Quantity: 15, RTProfileId: rt.ID}
	db.Create(data)

	update := &model.Inventory{Name: "Lampu LED", Quantity: 30}
	_, err := repo.Update(update, int(data.ID))
	assert.NoError(t, err)

	var check model.Inventory
	db.First(&check, data.ID)
	assert.Equal(t, "Lampu LED", check.Name)
	assert.Equal(t, 30, check.Quantity)
}

func TestInventory_Delete(t *testing.T) {
	db := setup(t)
	data := &model.Inventory{Name: "Sound System", Quantity: 2}
	db.Create(data)

	err := repo.Delete(int(data.ID))
	assert.NoError(t, err)

	var check model.Inventory
	tx := db.First(&check, data.ID)
	assert.True(t, errors.Is(tx.Error, gorm.ErrRecordNotFound))
}
