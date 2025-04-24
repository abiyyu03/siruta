package auth

import (
	"testing"

	"github.com/abiyyu03/siruta/config"
	"github.com/abiyyu03/siruta/entity/model"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var repo *AuthRepository

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)

	// AutoMigrate model user dan member
	err = db.AutoMigrate(&model.User{}, &model.Member{}, &model.Role{})
	assert.NoError(t, err)

	return db
}

func TestMain(m *testing.M) {
	repo = &AuthRepository{}
	m.Run()
}

func TestFetchLogin_SuccessWithMember(t *testing.T) {
	db := setupTestDB(t)
	config.DB = db

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("secret123"), bcrypt.DefaultCost)

	role := model.Role{ID: 1, Name: "User"}
	db.Create(&role)

	user := model.User{Email: "user@example.com", Password: string(hashedPassword), RoleID: uint(role.ID)}
	db.Create(&user)

	member := model.Member{UserId: user.ID, Fullname: "Test User"}
	db.Create(&member)

	u, m, err := repo.FetchLogin("user@example.com", "secret123")
	assert.NoError(t, err)
	assert.Equal(t, "user@example.com", u.Email)
	assert.NotNil(t, m)
	assert.Equal(t, "Test User", m.Fullname)
}

func TestFetchLogin_SuccessWithoutMember(t *testing.T) {
	db := setupTestDB(t)
	config.DB = db

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("adminpass"), bcrypt.DefaultCost)

	role := model.Role{Name: "Admin"}
	db.Create(&role)

	user := model.User{Email: "admin@example.com", Password: string(hashedPassword), RoleID: uint(role.ID)}
	db.Create(&user)

	u, m, err := repo.FetchLogin("admin@example.com", "adminpass")
	assert.NoError(t, err)
	assert.Equal(t, "admin@example.com", u.Email)
	assert.Nil(t, m)
}

func TestFetchLogin_WrongPassword(t *testing.T) {
	db := setupTestDB(t)
	config.DB = db

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("correctpass"), bcrypt.DefaultCost)
	role := model.Role{Name: "Guest"}
	db.Create(&role)

	user := model.User{Email: "guest@example.com", Password: string(hashedPassword), RoleID: uint(role.ID)}
	db.Create(&user)

	u, m, err := repo.FetchLogin("guest@example.com", "wrongpass")
	assert.Error(t, err)
	assert.Nil(t, u)
	assert.Nil(t, m)
}

func TestFetchLogin_UserNotFound(t *testing.T) {
	db := setupTestDB(t)
	config.DB = db

	u, m, err := repo.FetchLogin("notfound@example.com", "any")
	assert.Error(t, err)
	assert.Nil(t, u)
	assert.Nil(t, m)
}
