package http

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/abiyyu03/siruta/entity/model"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRoleUsecase struct {
	mock.Mock
}

func (m *MockRoleUsecase) Fetch(ctx *fiber.Ctx) error {
	args := m.Called(ctx)
	return args.Error(0)
}

func (m *MockRoleUsecase) FetchById(ctx *fiber.Ctx, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockRoleUsecase) Store(role *model.Role, ctx *fiber.Ctx) error {
	args := m.Called(ctx, role)
	return args.Error(0)
}
func (m *MockRoleUsecase) Update(role *model.Role, ctx *fiber.Ctx, id int) error {
	args := m.Called(role, ctx, id)
	return args.Error(0)
}
func (m *MockRoleUsecase) Delete(ctx *fiber.Ctx, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func TestRoleHttp_GetData(t *testing.T) {
	app := fiber.New()
	mockUsecase := new(MockRoleUsecase)
	handler := &RoleHttp{roleUsecase: mockUsecase}

	app.Get("/roles", handler.GetData)

	mockUsecase.On("Fetch", mock.Anything).Return(nil)

	req := httptest.NewRequest("GET", "/roles", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, 200, resp.StatusCode)
	mockUsecase.AssertExpectations(t)
}

func TestRoleHttp_GetDataById(t *testing.T) {
	mockUsecase := new(MockRoleUsecase)
	handler := &RoleHttp{roleUsecase: mockUsecase}

	app := fiber.New()
	app.Get("/roles/:id", handler.GetDataById)

	mockUsecase.On("FetchById", mock.Anything, 1).Return(nil)

	req := httptest.NewRequest("GET", "/roles/1", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	mockUsecase.AssertExpectations(t)
}

func TestRoleHttp_StoreData(t *testing.T) {
	mockUsecase := new(MockRoleUsecase)
	handler := &RoleHttp{roleUsecase: mockUsecase}

	app := fiber.New()
	app.Post("/roles", handler.StoreData)

	role := &model.Role{
		Name: "Admin aa",
	}
	body, _ := json.Marshal(role)

	mockUsecase.On("Store", mock.Anything, mock.MatchedBy(func(cf *model.Role) bool {
		return cf.Name == "Admin aa"
	})).Return(nil)

	req := httptest.NewRequest("POST", "/roles", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	mockUsecase.AssertExpectations(t)
}

func TestRoleHttp_UpdateData(t *testing.T) {
	mockUsecase := new(MockRoleUsecase)
	handler := &RoleHttp{roleUsecase: mockUsecase}
	app := fiber.New()

	app.Put("/roles/:id", handler.UpdateData)

	role := &model.Role{
		Name: "Admin aa",
	}
	updatedBody, _ := json.Marshal(role)

	mockUsecase.On("Update", mock.MatchedBy(func(data *model.Role) bool {
		return data.Name == "Admin aa"
	}), mock.Anything, 1).Return(nil)

	req := httptest.NewRequest("PUT", "/roles/1", bytes.NewReader(updatedBody))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	mockUsecase.AssertExpectations(t)
}

func TestRoleHttp_DeleteData(t *testing.T) {
	mockUsecase := new(MockRoleUsecase)
	handler := &RoleHttp{roleUsecase: mockUsecase}

	app := fiber.New()
	app.Delete("/roles/:id", handler.DeleteData)

	mockUsecase.On("Delete", mock.Anything, 1).Return(nil)

	req := httptest.NewRequest("DELETE", "/roles/1", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	mockUsecase.AssertExpectations(t)
}
