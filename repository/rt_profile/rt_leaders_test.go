package rt_profile_test

import (
	"testing"

	"github.com/abiyyu03/siruta/config"
	"github.com/abiyyu03/siruta/entity/model"
	"github.com/abiyyu03/siruta/repository/rt_profile"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func setupTestDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&model.RTLeader{}, &model.RTLeader{})
	if err != nil {
		panic(err)
	}

	return db
}

func TestRTLeaderRepository_Fetch(t *testing.T) {
	db := setupTestDB()
	config.DB = db

	repo := rt_profile.RTLeaderRepository{}

	gambar := "budi.jpg"

	leader := model.RTLeader{
		ID:          "lead-000",
		Fullname:    "Budi",
		NikNumber:   "1234567890",
		KKNumber:    "0987654321",
		RTProfileId: "rt-20002",
		Photo:       &gambar,
		UserId:      "user-001",
		FullAddress: "Jl. Mawar No.1",
	}

	db.Create(&leader)

	leader = model.RTLeader{
		ID:          "lead-001",
		Fullname:    "Budi",
		NikNumber:   "1234567890",
		KKNumber:    "0987654321",
		RTProfileId: "rt-20001",
		Photo:       &gambar,
		UserId:      "user-001",
		FullAddress: "Jl. Mawar No.1",
	}
	db.Create(&leader)

	results, err := repo.Fetch()
	assert.NoError(t, err)
	assert.Len(t, results, 1)
	assert.Equal(t, "Budi", results[0].Fullname)
	assert.Equal(t, "rt-20002", results[0].RTProfileId)
}

func TestRTLeaderRepository_FetchByRTProfileId(t *testing.T) {
	db := setupTestDB()
	config.DB = db

	repo := rt_profile.RTLeaderRepository{}

	rt := model.RTLeader{ID: "rt-002", Fullname: "03", RTProfileId: "vill-002"}
	db.Create(&rt)

	gambar := "siti.jpg"

	leader := model.RTLeader{
		ID:          "lead-002",
		Fullname:    "Siti",
		NikNumber:   "111222333444",
		KKNumber:    "555666777888",
		RTProfileId: rt.ID,
		Photo:       &gambar,
		UserId:      "user-002",
		FullAddress: "Jl. Melati No.5",
	}
	db.Create(&leader)

	results, err := repo.FetchByRTProfileId("rt-002")
	assert.NoError(t, err)
	assert.Len(t, results, 1)
	assert.Equal(t, "Siti", results[0].Fullname)
	assert.Equal(t, "rt-002", results[0].RTProfileId)
}

func TestRTLeaderRepository_FetchById(t *testing.T) {
	db := setupTestDB()
	config.DB = db

	repo := rt_profile.RTLeaderRepository{}

	gambar := "aminah.jpg"

	leader := model.RTLeader{
		ID:          "lead-003",
		Fullname:    "Aminah",
		NikNumber:   "999000111222",
		KKNumber:    "333444555666",
		RTProfileId: "rt-003",
		Photo:       &gambar,
		UserId:      "user-003",
		FullAddress: "Jl. Kenanga No.3",
	}
	db.Create(&leader)

	result, err := repo.FetchById("lead-003")
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "Aminah", result.Fullname)
}

func TestRTLeaderRepository_Store(t *testing.T) {
	db := setupTestDB()
	config.DB = db

	repo := rt_profile.RTLeaderRepository{}

	tx := db.Begin()
	defer tx.Rollback()

	rt := model.RTLeader{ID: "rt-004", Fullname: "Lama", NikNumber: "000000000000"}
	tx.Create(&rt)

	gambar := "lama.jpg"

	newLeader := &model.RTLeader{
		ID:          "lead-004",
		Fullname:    "Hadi",
		NikNumber:   "123123123123",
		KKNumber:    "321321321321",
		RTProfileId: rt.ID,
		Photo:       &gambar,
		UserId:      "user-004",
		FullAddress: "Jl. Anggrek No.7",
	}

	err := repo.Store(tx, newLeader)
	assert.NoError(t, err)

	var result model.RTLeader
	tx.First(&result, "id = ?", "lead-004")
	tx.Commit()

	assert.Equal(t, "Hadi", result.Fullname)
}

func TestRTLeaderRepository_Update(t *testing.T) {
	db := setupTestDB()
	config.DB = db

	repo := rt_profile.RTLeaderRepository{}

	tx := db.Begin()
	defer tx.Rollback()

	gambar := "lama.jpg"

	leader := model.RTLeader{
		ID:          "lead-005",
		Fullname:    "Lama",
		NikNumber:   "000000000000",
		KKNumber:    "999999999999",
		RTProfileId: "rt-005",
		Photo:       &gambar,
		UserId:      "user-005",
		FullAddress: "Jl. Lama No.9",
	}
	tx.Create(&leader)

	leader.Fullname = "Baru"
	leader.FullAddress = "Jl. Baru No.1"

	err := repo.Update(tx, &leader, "lead-005")
	assert.NoError(t, err)

	var updated model.RTLeader
	tx.First(&updated, "id = ?", "lead-005")
	tx.Commit()
	assert.Equal(t, "Baru", updated.Fullname)
	assert.Equal(t, "Jl. Baru No.1", updated.FullAddress)
}

func TestRTLeaderRepository_FetchById_NotFound(t *testing.T) {
	db := setupTestDB()
	config.DB = db

	repo := rt_profile.RTLeaderRepository{}

	result, err := repo.FetchById("non-existent-id")
	assert.Error(t, err)
	assert.Nil(t, result)
}
