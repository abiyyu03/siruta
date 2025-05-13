package guest_list

import (
	"errors"
	"testing"
	"time"

	"github.com/abiyyu03/siruta/config"
	"github.com/abiyyu03/siruta/entity/model"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var repo *GuestListRepository

func setup(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)

	err = db.AutoMigrate(&model.GuestList{})
	assert.NoError(t, err)

	config.DB = db
	repo = &GuestListRepository{}

	return db
}

func TestStoreGuest(t *testing.T) {
	setup(t)

	rtId, _ := uuid.NewV7()

	data := &model.GuestList{
		FullName:    "John Doe",
		PhoneNumber: "081234567890",
		VisitAt:     time.Now(),
		RTProfileId: rtId.String(),
	}

	result, err := repo.Store(data)
	assert.NoError(t, err)
	assert.NotZero(t, result.ID)
	assert.Equal(t, "John Doe", result.FullName)
}

func TestFetchGuests(t *testing.T) {
	db := setup(t)

	rtId, _ := uuid.NewV7()

	db.Create(&model.GuestList{FullName: "Alice", PhoneNumber: "0801", VisitAt: time.Now(), RTProfileId: rtId.String()})
	db.Create(&model.GuestList{FullName: "Bob", PhoneNumber: "0802", VisitAt: time.Now(), RTProfileId: rtId.String()})

	result, err := repo.Fetch()
	assert.NoError(t, err)
	assert.Len(t, result, 2)
}

func TestFetchById(t *testing.T) {
	db := setup(t)

	rtId, _ := uuid.NewV7()

	guest := model.GuestList{FullName: "Charlie", PhoneNumber: "0803", VisitAt: time.Now(), RTProfileId: rtId.String()}
	db.Create(&guest)

	result, err := repo.FetchById(int(guest.ID))
	assert.NoError(t, err)
	assert.Equal(t, guest.FullName, result.FullName)
}

func TestFetchByRTProfileId(t *testing.T) {
	db := setup(t)

	rtId, _ := uuid.NewV7()

	db.Create(&model.GuestList{FullName: "RT1 Guest", PhoneNumber: "0804", VisitAt: time.Now(), RTProfileId: rtId.String()})
	db.Create(&model.GuestList{FullName: "RT1 Guest 2", PhoneNumber: "0805", VisitAt: time.Now(), RTProfileId: rtId.String()})
	db.Create(&model.GuestList{FullName: "RT2 Guest", PhoneNumber: "0806", VisitAt: time.Now(), RTProfileId: rtId.String()})

	result, err := repo.FetchByRTProfileId(rtId.String())
	assert.NoError(t, err)
	assert.Len(t, result, 3)
}

func TestUpdateGuest(t *testing.T) {
	db := setup(t)

	rtId, _ := uuid.NewV7()

	guest := &model.GuestList{FullName: "Old Name", PhoneNumber: "0000", VisitAt: time.Now(), RTProfileId: rtId.String()}
	db.Create(guest)

	updated := &model.GuestList{FullName: "New Name", PhoneNumber: "9999", VisitAt: time.Now()}
	result, err := repo.Update(updated, int(guest.ID))

	assert.NoError(t, err)

	var check model.GuestList
	db.First(&check, guest.ID)
	assert.Equal(t, "New Name", check.FullName)
	assert.Equal(t, result.FullName, check.FullName)
}

func TestDeleteGuest(t *testing.T) {
	db := setup(t)

	rtId, _ := uuid.NewV7()
	guest := &model.GuestList{FullName: "To Be Deleted", PhoneNumber: "1234", VisitAt: time.Now(), RTProfileId: rtId.String()}
	db.Create(guest)

	_, err := repo.Delete(int(guest.ID))
	assert.NoError(t, err)

	var check model.GuestList
	tx := db.First(&check, guest.ID)
	assert.True(t, errors.Is(tx.Error, gorm.ErrRecordNotFound))
}
