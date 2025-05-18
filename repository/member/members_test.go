package member

import (
	"errors"
	"testing"

	"github.com/abiyyu03/siruta/config"
	"github.com/abiyyu03/siruta/entity/model"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var memberRepo *MemberRepository

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)

	err = db.AutoMigrate(&model.Member{})
	assert.NoError(t, err)

	return db
}

func seedMemberData(db *gorm.DB) *model.Member {
	member := &model.Member{
		ID:             "1",
		Fullname:       "John Doe",
		BornPlace:      "Jakarta",
		BirthDate:      "2000-01-01",
		Gender:         "Male",
		ReligionId:     1,
		MemberStatusId: 1,
		UserId:         "user-1",
		Status:         "contract",
		RTProfileId:    "rt-1",
	}
	db.Create(member)
	return member
}

func TestFetch(t *testing.T) {
	db := setupTestDB(t)
	memberRepo = &MemberRepository{}
	seedMemberData(db)

	oldDB := config.DB
	config.DB = db
	defer func() { config.DB = oldDB }()

	result, err := memberRepo.Fetch()
	assert.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, "John Doe", result[0].Fullname)
}

func TestFetchByNikAndBirthDate(t *testing.T) {
	db := setupTestDB(t)
	memberRepo = &MemberRepository{}

	nik := "1234567890"
	member := &model.Member{
		ID:             "2",
		Fullname:       "Jane Smith",
		BornPlace:      "Bandung",
		BirthDate:      "1995-05-20",
		Gender:         "Female",
		ReligionId:     1,
		MemberStatusId: 1,
		UserId:         "user-2",
		Status:         "resident",
		RTProfileId:    "rt-2",
		NikNumber:      &nik,
	}
	db.Create(member)

	oldDB := config.DB
	config.DB = db
	defer func() { config.DB = oldDB }()

	result, err := memberRepo.FetchByNikAndBirtDate("1234567890", "1995-05-20")
	assert.NoError(t, err)
	assert.Equal(t, "Jane Smith", result.Fullname)
}

func TestFetchById(t *testing.T) {
	db := setupTestDB(t)
	memberRepo = &MemberRepository{}
	member := seedMemberData(db)

	oldDB := config.DB
	config.DB = db
	defer func() { config.DB = oldDB }()

	result, err := memberRepo.FetchById(member.ID)
	assert.NoError(t, err)
	assert.Equal(t, "John Doe", result.Fullname)
}

func TestFetchByRTProfileId(t *testing.T) {
	db := setupTestDB(t)
	memberRepo = &MemberRepository{}
	member := seedMemberData(db)

	oldDB := config.DB
	config.DB = db
	defer func() { config.DB = oldDB }()

	result, err := memberRepo.FetchByRTProfileId(member.RTProfileId)
	assert.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, "John Doe", result[0].Fullname)
}

func TestStore(t *testing.T) {
	db := setupTestDB(t)
	memberRepo = &MemberRepository{}

	newMember := &model.Member{
		ID:             "3",
		Fullname:       "Alice",
		BornPlace:      "Medan",
		BirthDate:      "1990-12-12",
		Gender:         "Female",
		ReligionId:     2,
		MemberStatusId: 2,
		UserId:         "user-3",
		Status:         "Resident",
		RTProfileId:    "rt-3",
	}

	tx := db.Begin()
	result, err := memberRepo.Store(tx, newMember)
	tx.Commit()

	assert.NoError(t, err)
	assert.Equal(t, "Alice", result.Fullname)
}

func TestUpdate(t *testing.T) {
	db := setupTestDB(t)
	memberRepo = &MemberRepository{}

	member := seedMemberData(db)

	oldDB := config.DB
	config.DB = db
	defer func() { config.DB = oldDB }()

	member.Fullname = "Updated Name"
	updated, err := memberRepo.Update(member, member.ID)
	assert.NoError(t, err)
	assert.Equal(t, "Updated Name", updated.Fullname)

	var check model.Member
	db.First(&check, "id = ?", member.ID)
	assert.Equal(t, "Updated Name", check.Fullname)
}

func TestDelete(t *testing.T) {
	db := setupTestDB(t)
	memberRepo = &MemberRepository{}

	member := seedMemberData(db)

	oldDB := config.DB
	config.DB = db
	defer func() { config.DB = oldDB }()

	_, err := memberRepo.Delete(member, member.ID)
	assert.NoError(t, err)

	var check model.Member
	tx := db.First(&check, "id = ?", member.ID)
	assert.True(t, errors.Is(tx.Error, gorm.ErrRecordNotFound))
}
