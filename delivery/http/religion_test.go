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

type MockReligionUsecase struct {
	mock.Mock
}

func (m *MockReligionUsecase) Fetch(ctx *fiber.Ctx) error {
	args := m.Called(ctx)
	return args.Error(0)
}

func (m *MockReligionUsecase) FetchById(ctx *fiber.Ctx, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockReligionUsecase) Store(religion *model.Religion, ctx *fiber.Ctx) error {
	args := m.Called(ctx, religion)
	return args.Error(0)
}
func (m *MockReligionUsecase) Update(religion *model.Religion, ctx *fiber.Ctx, id int) error {
	args := m.Called(religion, ctx, id)
	return args.Error(0)
}
func (m *MockReligionUsecase) Delete(ctx *fiber.Ctx, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func TestReligionHttp_GetData(t *testing.T) {
	app := fiber.New()
	mockUsecase := new(MockReligionUsecase)
	handler := &ReligionHttp{religionUsecase: mockUsecase}

	app.Get("/religions", handler.GetData)

	mockUsecase.On("Fetch", mock.Anything).Return(nil)

	req := httptest.NewRequest("GET", "/religions", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, 200, resp.StatusCode)
	mockUsecase.AssertExpectations(t)
}

func TestReligionHttp_GetDataById(t *testing.T) {
	mockUsecase := new(MockReligionUsecase)
	handler := &ReligionHttp{religionUsecase: mockUsecase}

	app := fiber.New()
	app.Get("/religions/:id", handler.GetDataById)

	mockUsecase.On("FetchById", mock.Anything, 1).Return(nil)

	req := httptest.NewRequest("GET", "/religions/1", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	mockUsecase.AssertExpectations(t)
}

func TestReligionHttp_StoreData(t *testing.T) {
	mockUsecase := new(MockReligionUsecase)
	handler := &ReligionHttp{religionUsecase: mockUsecase}

	app := fiber.New()
	app.Post("/religions", handler.StoreData)

	religion := &model.Religion{
		ReligionName: "Islam",
	}
	body, _ := json.Marshal(religion)

	mockUsecase.On("Store", mock.Anything, mock.MatchedBy(func(cf *model.Religion) bool {
		return cf.ReligionName == "Islam"
	})).Return(nil)

	req := httptest.NewRequest("POST", "/religions", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	mockUsecase.AssertExpectations(t)
}

func TestReligionHttp_UpdateData(t *testing.T) {
	mockUsecase := new(MockReligionUsecase)
	handler := &ReligionHttp{religionUsecase: mockUsecase}
	app := fiber.New()

	app.Put("/religions/:id", handler.UpdateData)

	religion := &model.Religion{
		ReligionName: "Islam",
	}
	updatedBody, _ := json.Marshal(religion)

	mockUsecase.On("Update", mock.MatchedBy(func(data *model.Religion) bool {
		return data.ReligionName == "Islam"
	}), mock.Anything, 1).Return(nil)

	req := httptest.NewRequest("PUT", "/religions/1", bytes.NewReader(updatedBody))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	mockUsecase.AssertExpectations(t)
}

func TestReligionHttp_DeleteData(t *testing.T) {
	mockUsecase := new(MockReligionUsecase)
	handler := &ReligionHttp{religionUsecase: mockUsecase}

	app := fiber.New()
	app.Delete("/religions/:id", handler.DeleteData)

	mockUsecase.On("Delete", mock.Anything, 1).Return(nil)

	req := httptest.NewRequest("DELETE", "/religions/1", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	mockUsecase.AssertExpectations(t)
}
