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

type MockMemberStatusUsecase struct {
	mock.Mock
}

func (m *MockMemberStatusUsecase) Fetch(ctx *fiber.Ctx) error {
	args := m.Called(ctx)
	return args.Error(0)
}

func (m *MockMemberStatusUsecase) FetchById(ctx *fiber.Ctx, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}
func (m *MockMemberStatusUsecase) Store(inventory *model.MemberStatus, ctx *fiber.Ctx) error {
	args := m.Called(ctx, inventory)
	return args.Error(0)
}
func (m *MockMemberStatusUsecase) Update(inventory *model.MemberStatus, ctx *fiber.Ctx, id int) error {
	args := m.Called(inventory, ctx, id)
	return args.Error(0)
}
func (m *MockMemberStatusUsecase) Delete(ctx *fiber.Ctx, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func TestMemberStatusHttp_GetData(t *testing.T) {
	app := fiber.New()
	mockUsecase := new(MockMemberStatusUsecase)
	handler := &MemberStatusHttp{memberStatusUsecase: mockUsecase}

	app.Get("/member-status", handler.GetData)

	mockUsecase.On("Fetch", mock.Anything).Return(nil)

	req := httptest.NewRequest("GET", "/member-status", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, 200, resp.StatusCode)
	mockUsecase.AssertExpectations(t)
}

func TestMemberStatusHttp_GetDataById(t *testing.T) {
	mockUsecase := new(MockMemberStatusUsecase)
	handler := &MemberStatusHttp{memberStatusUsecase: mockUsecase}

	app := fiber.New()
	app.Get("/member-status/:id", handler.GetDataById)

	mockUsecase.On("FetchById", mock.Anything, 1).Return(nil)

	req := httptest.NewRequest("GET", "/member-status/1", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	mockUsecase.AssertExpectations(t)
}

func TestMemberStatusHttp_StoreData(t *testing.T) {
	mockUsecase := new(MockMemberStatusUsecase)
	handler := &MemberStatusHttp{memberStatusUsecase: mockUsecase}

	app := fiber.New()
	app.Post("/member-status", handler.StoreData)

	inventory := &model.MemberStatus{
		Status: "Kepala Keluarga",
	}
	body, _ := json.Marshal(inventory)

	mockUsecase.On("Store", mock.Anything, mock.MatchedBy(func(cf *model.MemberStatus) bool {
		return cf.Status == "Kepala Keluarga"
	})).Return(nil)

	req := httptest.NewRequest("POST", "/member-status", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	mockUsecase.AssertExpectations(t)
}

func TestMemberStatusHttp_UpdateData(t *testing.T) {
	mockUsecase := new(MockMemberStatusUsecase)
	handler := &MemberStatusHttp{memberStatusUsecase: mockUsecase}
	app := fiber.New()

	app.Put("/member-status/:id", handler.UpdateData)

	inventory := &model.MemberStatus{
		Status: "Ibu Rumah Tangga",
	}
	updatedBody, _ := json.Marshal(inventory)

	mockUsecase.On("Update", mock.MatchedBy(func(data *model.MemberStatus) bool {
		return data.Status == "Ibu Rumah Tangga"
	}), mock.Anything, 1).Return(nil)

	req := httptest.NewRequest("PUT", "/member-status/1", bytes.NewReader(updatedBody))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	mockUsecase.AssertExpectations(t)
}

func TestMemberStatusHttp_DeleteData(t *testing.T) {
	mockUsecase := new(MockMemberStatusUsecase)
	handler := &MemberStatusHttp{memberStatusUsecase: mockUsecase}

	app := fiber.New()
	app.Delete("/member-status/:id", handler.DeleteData)

	mockUsecase.On("Delete", mock.Anything, 1).Return(nil)

	req := httptest.NewRequest("DELETE", "/member-status/1", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	mockUsecase.AssertExpectations(t)
}
