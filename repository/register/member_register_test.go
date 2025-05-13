package register

import (
	"errors"
	"testing"

	"github.com/abiyyu03/siruta/config"
	"github.com/abiyyu03/siruta/entity/model"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Dummy user repo
type MockUserRepository struct {
	ShouldFail bool
}

func (r *MockUserRepository) RegisterUser(tx *gorm.DB, user *model.User, roleId int) (*model.User, error) {
	if r.ShouldFail {
		return nil, errors.New("user register error")
	}
	return user, nil
}

// Dummy member repo
type MockMemberRepository struct {
	ShouldFail bool
}

func setupTestDatabase(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)
	_ = db.AutoMigrate(&model.User{}, &model.Member{})
	config.DB = db
	return db
}

func (r *MockMemberRepository) Store(tx *gorm.DB, member *model.Member) (*model.Member, error) {
	if r.ShouldFail {
		return nil, errors.New("member store error")
	}
	return member, nil
}

func TestRegisterMember_Success(t *testing.T) {
	db := setupTestDatabase(t)

	// Inject mock
	register.UserRepository = &MockUserRepository{}
	register.MemberRepository = &MockMemberRepository{}

	repo := register.MemberRegisterRepository{}

	member := &model.Member{
		Fullname: "John Doe",
	}
	user := &model.User{
		Email:  "john@example.com",
		RoleID: 4,
	}

	err := repo.RegisterMember(member, user)

	assert.Nil(t, err)
}

func TestRegisterMember_UserFail(t *testing.T) {
	setupTestDatabase(t)

	register.UserRepository = &MockUserRepository{ShouldFail: true}
	register.MemberRepository = &MockMemberRepository{}

	repo := register.MemberRegisterRepository{}

	member := &model.Member{Fullname: "John"}
	user := &model.User{Email: "fail@example.com", RoleID: 4}

	err := repo.RegisterMember(member, user)

	assert.NotNil(t, err)
	assert.EqualError(t, err, "user register error")
}

func TestRegisterMember_MemberFail(t *testing.T) {
	setupTestDatabase(t)

	register.UserRepository = &MockUserRepository{}
	register.MemberRepository = &MockMemberRepository{ShouldFail: true}

	repo := register.MemberRegisterRepository{}

	member := &model.Member{Fullname: "Fail Member"}
	user := &model.User{Email: "memberfail@example.com", RoleID: 4}

	err := repo.RegisterMember(member, user)

	assert.NotNil(t, err)
	assert.EqualError(t, err, "member store error")
}
