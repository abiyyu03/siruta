package role

import (
	"errors"
	"testing"

	"github.com/abiyyu03/siruta/config"
	"github.com/abiyyu03/siruta/entity/model"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var repo *RoleRepository

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)

	err = db.AutoMigrate(&model.Role{})
	assert.NoError(t, err)

	return db
}

func TestFetch(t *testing.T) {
	db := setupTestDB(t)

	// Seed data
	db.Create(&model.Role{ID: 1, Name: "Admin"})
	db.Create(&model.Role{ID: 2, Name: "RT"})

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

	db.Create(&model.Role{ID: 1, Name: "Admin"})

	oldDB := config.DB
	config.DB = db
	defer func() { config.DB = oldDB }()

	result, err := repo.FetchById(1)
	assert.NoError(t, err)
	assert.Equal(t, "Admin", result.Name)
}

func TestStore(t *testing.T) {
	db := setupTestDB(t)

	oldDB := config.DB
	config.DB = db
	defer func() { config.DB = oldDB }()

	newRole := &model.Role{Name: "Member"}

	result, err := repo.Store(newRole)
	assert.NoError(t, err)
	assert.NotZero(t, result.ID)
}

func TestUpdate(t *testing.T) {
	db := setupTestDB(t)

	role := &model.Role{Name: "RT"}
	db.Create(role)

	oldDB := config.DB
	config.DB = db
	defer func() { config.DB = oldDB }()

	role.Name = "RT"
	updatedRole, err := repo.Update(role, int(role.ID))
	assert.NoError(t, err)

	var check model.Role
	db.First(&check, role.ID)
	assert.Equal(t, "RT", check.Name)
	assert.Equal(t, updatedRole.Name, "RT")
}

func TestDelete(t *testing.T) {
	db := setupTestDB(t)

	role := &model.Role{Name: "ToDelete"}
	db.Create(role)

	oldDB := config.DB
	config.DB = db
	defer func() { config.DB = oldDB }()

	_, err := repo.Delete(int(role.ID))
	assert.NoError(t, err)

	var check model.Role
	tx := db.First(&check, role.ID)
	assert.True(t, errors.Is(tx.Error, gorm.ErrRecordNotFound))
}
