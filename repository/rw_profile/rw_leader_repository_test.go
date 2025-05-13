package rw_profile_test

import (
	"testing"

	"github.com/abiyyu03/siruta/config"
	"github.com/abiyyu03/siruta/entity/model"
	"github.com/abiyyu03/siruta/repository/rw_profile"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Setup in-memory test DB
func setupTestDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic("failed to connect database")
	}
	err = db.AutoMigrate(&model.RWLeader{}, &model.RWProfile{})
	if err != nil {
		panic("failed to migrate tables")
	}
	return db
}

func TestFetch(t *testing.T) {
	db := setupTestDB()
	config.DB = db

	db.Create(&model.RWProfile{ID: "101", FullAddress: "Jalan Damai"})
	db.Create(&model.RWLeader{
		ID:          "1",
		Fullname:    "John Doe",
		KKNumber:    "0987654321",
		RWProfileId: "101",
		UserId:      "u123",
		FullAddress: "Jl. Sejahtera",
	})

	repo := &rw_profile.RWLeaderRepository{}
	data, err := repo.Fetch()

	assert.NoError(t, err)
	assert.Len(t, data, 1)
	assert.Equal(t, "John Doe", data[0].Fullname)
}

func TestFetchById(t *testing.T) {
	db := setupTestDB()
	config.DB = db

	db.Create(&model.RWLeader{
		ID:          "123",
		Fullname:    "Jane Doe",
		KKNumber:    "2222222222",
		RWProfileId: "102",
		UserId:      "u124",
		FullAddress: "Jl. Mawar",
	})

	repo := &rw_profile.RWLeaderRepository{}
	result, err := repo.FetchById("123")

	assert.NoError(t, err)
	assert.Equal(t, "Jane Doe", result.Fullname)
}

func TestStore(t *testing.T) {
	db := setupTestDB()
	tx := db.Begin()

	repo := &rw_profile.RWLeaderRepository{}
	rw := &model.RWLeader{
		ID:          "200",
		Fullname:    "RW Baru",
		KKNumber:    "4444444444",
		RWProfileId: "103",
		UserId:      "u125",
		FullAddress: "Jl. Baru",
	}

	err := repo.Store(tx, rw)
	assert.NoError(t, err)

	var found model.RWLeader
	err = tx.First(&found, "id = ?", "200").Error
	assert.NoError(t, err)
	assert.Equal(t, "RW Baru", found.Fullname)

	tx.Rollback()
}

func TestUpdate(t *testing.T) {
	db := setupTestDB()
	tx := db.Begin()

	rw := model.RWLeader{
		ID:          "300",
		Fullname:    "RW Lama",
		KKNumber:    "6666666666",
		RWProfileId: "104",
		UserId:      "u126",
		FullAddress: "Jl. Lama",
	}
	tx.Create(&rw)

	repo := &rw_profile.RWLeaderRepository{}

	rw.Fullname = "RW Update"
	err := repo.Update(tx, &rw, "300")
	assert.NoError(t, err)

	var updated model.RWLeader
	tx.First(&updated, "id = ?", "300")
	assert.Equal(t, "RW Update", updated.Fullname)

	tx.Rollback()
}
