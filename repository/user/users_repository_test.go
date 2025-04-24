package user_test

import (
	"testing"

	"github.com/abiyyu03/siruta/config"
	"github.com/abiyyu03/siruta/entity/model"
	"github.com/abiyyu03/siruta/repository/user"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func setupTestDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&model.User{}, &model.Role{})
	if err != nil {
		panic(err)
	}

	return db
}

func TestFetch(t *testing.T) {
	db := setupTestDB()
	config.DB = db // override config.DB temporarily

	repo := user.UserRepository{}

	role := model.Role{ID: 1, Name: "admin"}
	db.Create(&role)

	db.Create(&model.User{
		ID:           "user-1",
		Email:        "user1@example.com",
		Password:     "hashed",
		IsAuthorized: true,
		RoleID:       uint(role.ID),
	})

	users, err := repo.Fetch()
	assert.NoError(t, err)
	assert.Len(t, users, 1)
	assert.Equal(t, "user1@example.com", users[0].Email)
}

func TestFetchByEmail(t *testing.T) {
	db := setupTestDB()
	config.DB = db

	repo := user.UserRepository{}

	db.Create(&model.User{
		ID:           "user-2",
		Email:        "user2@example.com",
		Password:     "pw2",
		IsAuthorized: true,
	})

	found, err := repo.FetchByEmail("user2@example.com")
	assert.NoError(t, err)
	assert.NotNil(t, found)
	assert.Equal(t, "user2@example.com", found.Email)
}

func TestFetchById(t *testing.T) {
	db := setupTestDB()
	config.DB = db

	repo := user.UserRepository{}

	db.Create(&model.User{
		ID:           "user-3",
		Email:        "user3@example.com",
		Password:     "pw3",
		IsAuthorized: true,
	})

	found, err := repo.FetchById("user-3")
	assert.NoError(t, err)
	assert.NotNil(t, found)
	assert.Equal(t, "user3@example.com", found.Email)
}

func TestRegisterUser(t *testing.T) {
	db := setupTestDB()
	config.DB = db

	repo := user.UserRepository{}

	tx := db.Begin()
	defer tx.Rollback()

	newUser := &model.User{
		ID:           "user-4",
		Email:        "user4@example.com",
		Password:     "pw4",
		IsAuthorized: true,
		RoleID:       1,
	}

	created, err := repo.RegisterUser(tx, newUser, 1)
	assert.NoError(t, err)
	assert.Equal(t, "user4@example.com", created.Email)
}

func TestUpdatePassword(t *testing.T) {
	db := setupTestDB()
	config.DB = db

	repo := user.UserRepository{}

	db.Create(&model.User{
		ID:       "user-5",
		Email:    "user5@example.com",
		Password: "oldpass",
	})

	tx := db.Begin()
	defer tx.Rollback()

	err := repo.UpdatePassword(tx, "user-5", "newpass")
	assert.NoError(t, err)

	tx.Commit()

	var updated model.User
	db.First(&updated, "id = ?", "user-5")
	assert.Equal(t, "newpass", updated.Password)
}

func TestRevokeUserAccess(t *testing.T) {
	db := setupTestDB()
	config.DB = db

	repo := user.UserRepository{}

	db.Create(&model.User{
		ID:           "user-6",
		Email:        "user6@example.com",
		IsAuthorized: true,
	})

	err := repo.RevokeUserAccess("user-6")
	assert.NoError(t, err)

	var revoked model.User
	db.First(&revoked, "id = ?", "user-6")
	assert.False(t, revoked.IsAuthorized)
}
