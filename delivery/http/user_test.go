package http

import (
	"net/http/httptest"
	"testing"

	"github.com/abiyyu03/siruta/entity/model"
	"github.com/abiyyu03/siruta/entity/request"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserUsecase struct {
	mock.Mock
}

func (m *MockUserUsecase) Fetch(c *fiber.Ctx) error {
	return m.Called(c).Error(0)
}

func (m *MockUserUsecase) FetchById(c *fiber.Ctx, id string) error {
	return m.Called(c, id).Error(0)
}

func (m *MockUserUsecase) Register(c *fiber.Ctx) error {
	return m.Called(c).Error(0)
}

func (m *MockUserUsecase) RegisterUserWithTokenVerification(c *fiber.Ctx, user *model.User, token string) error {
	return m.Called(c, user, token).Error(0)
}
func (m *MockUserUsecase) UpdateProfilePhoto(c *fiber.Ctx, userId string, profileType string, req *request.UpdateProfilePhoto) error {
	return m.Called(c, userId, profileType, req).Error(0)
}

func (m *MockUserUsecase) TokenVerification(user *model.User, roleId uint, token string) (*model.User, string, error) {
	args := m.Called(user, roleId, token)

	var userResult *model.User
	if args.Get(0) != nil {
		userResult = args.Get(0).(*model.User)
	}

	var tokenResult string
	if args.Get(1) != nil {
		tokenResult = args.String(1)
	}

	return userResult, tokenResult, args.Error(2)
}

func (m *MockUserUsecase) RevokeUserAccess(c *fiber.Ctx, userId string) error {
	return m.Called(c, userId).Error(0)
}

func TestUserHttp_GetData(t *testing.T) {
	app := fiber.New()
	mockUC := new(MockUserUsecase)
	handler := &UserHttp{userUsecase: mockUC}

	app.Get("/users", handler.GetData)

	mockUC.On("Fetch", mock.Anything).Return(nil)

	req := httptest.NewRequest("GET", "/users", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	mockUC.AssertCalled(t, "Fetch", mock.Anything)
}

func TestUserHttp_GetDataById(t *testing.T) {
	app := fiber.New()
	mockUC := new(MockUserUsecase)
	handler := &UserHttp{userUsecase: mockUC}

	app.Get("/users/:id", handler.GetDataById)

	mockUC.On("FetchById", mock.Anything, "abc123").Return(nil)

	req := httptest.NewRequest("GET", "/users/abc123", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	mockUC.AssertCalled(t, "FetchById", mock.Anything, "abc123")
}

func TestUserHttp_RevokeUser(t *testing.T) {
	app := fiber.New()
	mockUC := new(MockUserUsecase)
	handler := &UserHttp{userUsecase: mockUC}

	app.Delete("/users/:id", handler.RevokeUser)

	mockUC.On("RevokeUserAccess", mock.Anything, "revoke-id").Return(nil)

	req := httptest.NewRequest("DELETE", "/users/revoke-id", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	mockUC.AssertCalled(t, "RevokeUserAccess", mock.Anything, "revoke-id")
}
