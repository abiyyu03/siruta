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

type MockInventoryUsecase struct {
	mock.Mock
}

func (m *MockInventoryUsecase) Fetch(ctx *fiber.Ctx) error {
	args := m.Called(ctx)
	return args.Error(0)
}
func (m *MockInventoryUsecase) FetchByRTProfileId(ctx *fiber.Ctx, rtProfileId string) error {
	args := m.Called(ctx, rtProfileId)
	return args.Error(0)
}

func (m *MockInventoryUsecase) FetchById(ctx *fiber.Ctx, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}
func (m *MockInventoryUsecase) Store(inventory *model.Inventory, ctx *fiber.Ctx) error {
	args := m.Called(ctx, inventory)
	return args.Error(0)
}
func (m *MockInventoryUsecase) Update(inventory *model.Inventory, ctx *fiber.Ctx, id int) error {
	args := m.Called(inventory, ctx, id)
	return args.Error(0)
}
func (m *MockInventoryUsecase) Delete(ctx *fiber.Ctx, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func TestInventoryHttp_GetData(t *testing.T) {
	app := fiber.New()
	mockUsecase := new(MockInventoryUsecase)
	handler := &InventoryHttp{inventoryUsecase: mockUsecase}

	app.Get("/inventories", handler.GetData)

	mockUsecase.On("Fetch", mock.Anything).Return(nil)

	req := httptest.NewRequest("GET", "/inventories", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, 200, resp.StatusCode)
	mockUsecase.AssertExpectations(t)
}

func TestInventoryHttp_GetDataByRTProfileId(t *testing.T) {
	mockUsecase := new(MockInventoryUsecase)
	handler := &InventoryHttp{inventoryUsecase: mockUsecase}

	app := fiber.New()
	app.Get("/inventories/rt/:rt_profile_id", handler.GetDataByRTProfileId)

	mockUsecase.On("FetchByRTProfileId", mock.Anything, "1").Return(nil)

	req := httptest.NewRequest("GET", "/inventories/rt/1", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	mockUsecase.AssertExpectations(t)
}

func TestInventoryHttp_GetDataById(t *testing.T) {
	mockUsecase := new(MockInventoryUsecase)
	handler := &InventoryHttp{inventoryUsecase: mockUsecase}

	app := fiber.New()
	app.Get("/inventories/:id", handler.GetDataById)

	mockUsecase.On("FetchById", mock.Anything, 1).Return(nil)

	req := httptest.NewRequest("GET", "/inventories/1", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	mockUsecase.AssertExpectations(t)
}

func TestInventoryHttp_StoreData(t *testing.T) {
	mockUsecase := new(MockInventoryUsecase)
	handler := &InventoryHttp{inventoryUsecase: mockUsecase}

	app := fiber.New()
	app.Post("/inventories", handler.StoreData)

	image := "laptop.jpg"

	inventory := &model.Inventory{
		Name:        "Laptop",
		Quantity:    10,
		Image:       &image,
		RTProfileId: "LAP001",
	}
	body, _ := json.Marshal(inventory)

	mockUsecase.On("Store", mock.Anything, mock.MatchedBy(func(cf *model.Inventory) bool {
		return cf.Quantity == 10
	})).Return(nil)

	req := httptest.NewRequest("POST", "/inventories", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	mockUsecase.AssertExpectations(t)
}

func TestInventoryHttp_UpdateData(t *testing.T) {
	mockUsecase := new(MockInventoryUsecase)
	handler := &InventoryHttp{inventoryUsecase: mockUsecase}
	app := fiber.New()

	app.Put("/inventories/:id", handler.UpdateData)

	inventory := &model.Inventory{
		ID:          1,
		Name:        "Komputer High End",
		Quantity:    10,
		RTProfileId: "LAP001",
	}
	updatedBody, _ := json.Marshal(inventory)

	mockUsecase.On("Update", mock.MatchedBy(func(data *model.Inventory) bool {
		return data.Name == "Komputer High End"
	}), mock.Anything, 1).Return(nil)

	req := httptest.NewRequest("PUT", "/inventories/1", bytes.NewReader(updatedBody))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	mockUsecase.AssertExpectations(t)
}

func TestInventoryHttp_DeleteData(t *testing.T) {
	mockUsecase := new(MockInventoryUsecase)
	handler := &InventoryHttp{inventoryUsecase: mockUsecase}

	app := fiber.New()
	app.Delete("/inventories/:id", handler.DeleteData)

	mockUsecase.On("Delete", mock.Anything, 1).Return(nil)

	req := httptest.NewRequest("DELETE", "/inventories/1", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	mockUsecase.AssertExpectations(t)
}
