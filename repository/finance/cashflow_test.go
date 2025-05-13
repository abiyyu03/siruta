package finance

import (
	"errors"
	"testing"

	"github.com/abiyyu03/siruta/config"
	"github.com/abiyyu03/siruta/entity/model"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var repo *CashflowRepository

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)

	err = db.AutoMigrate(&model.Cashflow{})
	assert.NoError(t, err)

	return db
}

func TestMain(m *testing.M) {
	repo = &CashflowRepository{}
	m.Run()
}

func TestFetch(t *testing.T) {
	db := setupTestDB(t)

	db.Create(&model.Cashflow{Description: "Desc1", LogType: "Pemasukan", Amount: 1000, RTProfileId: "uuid-1", PaymentPeriodYear: "2024", PaymentPeriodMonth: "Mei"})
	db.Create(&model.Cashflow{Description: "Desc2", LogType: "Pengeluaran", Amount: 2000, RTProfileId: "uuid-2", PaymentPeriodYear: "2024", PaymentPeriodMonth: "Maret"})

	old := config.DB
	config.DB = db
	defer func() { config.DB = old }()

	result, err := repo.Fetch("Pemasukan")
	assert.NoError(t, err)
	assert.Len(t, result, 1)
}

func TestFetchById(t *testing.T) {
	db := setupTestDB(t)

	data := model.Cashflow{Description: "Desc A", LogType: "Pemasukan", Amount: 123000, RTProfileId: "uuid-3", PaymentPeriodYear: "2023", PaymentPeriodMonth: "desember"}
	db.Create(&data)

	old := config.DB
	config.DB = db
	defer func() { config.DB = old }()

	result, err := repo.FetchById(data.ID, "Pemasukan")
	assert.NoError(t, err)
	assert.Equal(t, data.Description, result.Description)
}

func TestFetchByRTProfileId(t *testing.T) {
	db := setupTestDB(t)

	rtID := "rt-uuid-123"
	db.Create(&model.Cashflow{Description: "RT A", LogType: "Pemasukan", Amount: 100, RTProfileId: rtID, PaymentPeriodYear: "2024", PaymentPeriodMonth: "MEI"})
	db.Create(&model.Cashflow{Description: "RT B", LogType: "Pengeluaran", Amount: 150, RTProfileId: rtID, PaymentPeriodYear: "2024", PaymentPeriodMonth: "Mei"})
	db.Create(&model.Cashflow{Description: "RT B", LogType: "Pengeluaran", Amount: 150, RTProfileId: rtID, PaymentPeriodYear: "2024", PaymentPeriodMonth: "Mei"})

	old := config.DB
	config.DB = db
	defer func() { config.DB = old }()

	result, err := repo.FetchByRTProfileId(rtID, "Pengeluaran")
	assert.NoError(t, err)
	assert.Len(t, result, 2)
}

func TestStore(t *testing.T) {
	db := setupTestDB(t)

	old := config.DB
	config.DB = db
	defer func() { config.DB = old }()

	data := &model.Cashflow{Description: "New Entry", LogType: "Pemasukan", Amount: 5000, RTProfileId: "uuid-new", PaymentPeriodYear: "2025", PaymentPeriodMonth: "Mei"}
	result, err := repo.Store(data)
	assert.NoError(t, err)
	assert.NotZero(t, result.ID)
}

func TestUpdate(t *testing.T) {
	db := setupTestDB(t)

	data := model.Cashflow{Description: "Old", LogType: "Pemasukan", Amount: 1000, RTProfileId: "uuid-upd", PaymentPeriodYear: "2024", PaymentPeriodMonth: "Mei"}
	db.Create(&data)

	old := config.DB
	config.DB = db
	defer func() { config.DB = old }()

	updateData := &model.Cashflow{Description: "Updated", LogType: "Pengeluaran", Amount: 9999}
	result, err := repo.Update(updateData, int(data.ID))
	assert.NoError(t, err)

	var updated model.Cashflow
	db.First(&updated, data.ID)
	assert.Equal(t, "Updated", updated.Description)
	assert.Equal(t, "Pengeluaran", updated.LogType)
	assert.Equal(t, result.Description, updated.Description)
}

func TestDelete(t *testing.T) {
	db := setupTestDB(t)

	data := model.Cashflow{Description: "Delete Me", LogType: "Pemasukan", Amount: 1111, RTProfileId: "uuid-del", PaymentPeriodYear: "2023", PaymentPeriodMonth: "Bogor"}
	db.Create(&data)

	old := config.DB
	config.DB = db
	defer func() { config.DB = old }()

	err := repo.Delete(int(data.ID))
	assert.NoError(t, err)

	var result model.Cashflow
	tx := db.First(&result, data.ID)
	assert.True(t, errors.Is(tx.Error, gorm.ErrRecordNotFound))
}
