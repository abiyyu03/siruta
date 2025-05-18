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

type MockUserRepository struct {
	ShouldFail bool
}

func (r *MockUserRepository) RegisterUser(tx *gorm.DB, user *model.User, roleId int) (*model.User, error) {
	if r.ShouldFail {
		return nil, errors.New("user register error")
	}
	return user, nil
}

type MockMemberRepository struct {
	ShouldFail bool
}

func (r *MockMemberRepository) Store(tx *gorm.DB, member *model.Member) (*model.Member, error) {
	if r.ShouldFail {
		return nil, errors.New("member store error")
	}
	return member, nil
}

func setup(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)
	assert.NoError(t, db.AutoMigrate(&model.User{}, &model.Member{}))

	config.DB = db
	return db
}

func TestRegisterMember_Success(t *testing.T) {
	setup(t)

	repo := MemberRegisterRepository{}

	member := &model.Member{Fullname: "John Doe"}
	user := &model.User{Email: "john@example.com", RoleID: 4}

	err := repo.RegisterMember(member, user)
	assert.NoError(t, err)
}

func TestRegisterMember_UserFail(t *testing.T) {
	setup(t)
	repo := MemberRegisterRepository{}

	member := &model.Member{Fullname: "Fail User"}
	user := &model.User{Email: "fail@example.com", RoleID: 4}

	err := repo.RegisterMember(member, user)
	assert.EqualError(t, err, "user register error")
}

func TestRegisterMember_MemberFail(t *testing.T) {
	setup(t)

	repo := MemberRegisterRepository{}

	member := &model.Member{Fullname: "Fail Member"}
	user := &model.User{Email: "failmember@example.com", RoleID: 4}

	err := repo.RegisterMember(member, user)
	assert.EqualError(t, err, "member store error")
}
