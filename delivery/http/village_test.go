package http

import (
	"bytes"
	"net/http/httptest"
	"testing"

	"github.com/abiyyu03/siruta/entity/model"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockVillageUsecase struct {
	mock.Mock
}

func (m *MockVillageUsecase) Fetch(c *fiber.Ctx) error {
	return m.Called(c).Error(0)
}

func (m *MockVillageUsecase) FetchById(c *fiber.Ctx, id int) error {
	return m.Called(c, id).Error(0)
}

func (m *MockVillageUsecase) Store(village *model.Village, c *fiber.Ctx) error {
	return m.Called(village, c).Error(0)
}

func (m *MockVillageUsecase) Update(village *model.Village, c *fiber.Ctx, id int) error {
	return m.Called(village, c, id).Error(0)
}

func (m *MockVillageUsecase) Delete(c *fiber.Ctx, id int) error {
	return m.Called(c, id).Error(0)
}

func TestVillageHttp_GetData(t *testing.T) {
	app := fiber.New()
	mockUC := new(MockVillageUsecase)
	handler := NewVillageHttp(mockUC)

	app.Get("/villages", handler.GetData)

	mockUC.On("Fetch", mock.Anything).Return(nil)

	req := httptest.NewRequest("GET", "/villages", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	mockUC.AssertCalled(t, "Fetch", mock.Anything)
}

func TestVillageHttp_GetDataById(t *testing.T) {
	app := fiber.New()
	mockUC := new(MockVillageUsecase)
	handler := NewVillageHttp(mockUC)

	app.Get("/villages/:id", handler.GetDataById)

	mockUC.On("FetchById", mock.Anything, 1).Return(nil)

	req := httptest.NewRequest("GET", "/villages/1", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	mockUC.AssertCalled(t, "FetchById", mock.Anything, 1)
}

func TestVillageHttp_StoreData(t *testing.T) {
	app := fiber.New()
	mockUC := new(MockVillageUsecase)
	handler := NewVillageHttp(mockUC)

	app.Post("/villages", handler.StoreData)

	jsonStr := []byte(`{"name":"Test Village"}`)

	mockUC.On("Store", mock.Anything, mock.Anything).Return(nil)

	req := httptest.NewRequest("POST", "/villages", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	resp, _ := app.Test(req)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	mockUC.AssertCalled(t, "Store", mock.Anything, mock.Anything)
}

func TestVillageHttp_UpdateData(t *testing.T) {
	app := fiber.New()
	mockUC := new(MockVillageUsecase)
	handler := NewVillageHttp(mockUC)

	app.Put("/villages/:id", handler.UpdateData)

	jsonStr := []byte(`{"name":"Updated Village"}`)

	mockUC.On("Update", mock.Anything, mock.Anything, 2).Return(nil)

	req := httptest.NewRequest("PUT", "/villages/2", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	resp, _ := app.Test(req)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	mockUC.AssertCalled(t, "Update", mock.Anything, mock.Anything, 2)
}

func TestVillageHttp_DeleteData(t *testing.T) {
	app := fiber.New()
	mockUC := new(MockVillageUsecase)
	handler := NewVillageHttp(mockUC)

	app.Delete("/villages/:id", handler.DeleteData)

	mockUC.On("Delete", mock.Anything, 5).Return(nil)

	req := httptest.NewRequest("DELETE", "/villages/5", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	mockUC.AssertCalled(t, "Delete", mock.Anything, 5)
}
