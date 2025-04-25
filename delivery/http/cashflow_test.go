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

type MockCashflowUsecase struct {
	mock.Mock
}

func (m *MockCashflowUsecase) Fetch(ctx *fiber.Ctx, queryType string) error {
	args := m.Called(ctx, queryType)
	return args.Error(0)
}
func (m *MockCashflowUsecase) FetchByRTProfileId(ctx *fiber.Ctx, rtProfileId string, queryType string) error {
	args := m.Called(ctx, rtProfileId, queryType)
	return args.Error(0)
}
func (m *MockCashflowUsecase) FetchById(ctx *fiber.Ctx, id int, queryType string) error {
	args := m.Called(ctx, id, queryType)
	return args.Error(0)
}
func (m *MockCashflowUsecase) Store(ctx *fiber.Ctx, cashflow *model.Cashflow) error {
	args := m.Called(ctx, cashflow)
	return args.Error(0)
}
func (m *MockCashflowUsecase) Update(ctx *fiber.Ctx, cashflow *model.Cashflow, id int) error {
	args := m.Called(ctx, cashflow, id)
	return args.Error(0)
}
func (m *MockCashflowUsecase) Delete(ctx *fiber.Ctx, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func TestCashflowHttp_GetData(t *testing.T) {
	app := fiber.New()
	mockUsecase := new(MockCashflowUsecase)
	handler := &CashflowHttp{cashflowUsecase: mockUsecase}

	app.Get("/cashflow", handler.GetData)

	mockUsecase.On("Fetch", mock.Anything, "masuk").Return(nil)

	req := httptest.NewRequest("GET", "/cashflow?type=masuk", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, 200, resp.StatusCode)
	mockUsecase.AssertExpectations(t)
}

func TestCashflowHttp_GetDataByRTProfileId(t *testing.T) {
	mockUsecase := new(MockCashflowUsecase)
	handler := &CashflowHttp{cashflowUsecase: mockUsecase}

	app := fiber.New()
	app.Get("/cashflow/rt/:rt_profile_id", handler.GetDataByRTProfileId)

	mockUsecase.On("FetchByRTProfileId", mock.Anything, "1", "expense").Return(nil)

	req := httptest.NewRequest("GET", "/cashflow/rt/1?type=expense", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	mockUsecase.AssertExpectations(t)
}

func TestCashflowHttp_StoreData(t *testing.T) {
	mockUsecase := new(MockCashflowUsecase)
	handler := &CashflowHttp{cashflowUsecase: mockUsecase}

	app := fiber.New()
	app.Post("/cashflow", handler.StoreData)

	cashflow := &model.Cashflow{Amount: 1000}
	body, _ := json.Marshal(cashflow)

	mockUsecase.On("Store", mock.Anything, mock.MatchedBy(func(cf *model.Cashflow) bool {
		return cf.Amount == 1000
	})).Return(nil)

	req := httptest.NewRequest("POST", "/cashflow", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	mockUsecase.AssertExpectations(t)
}

func TestCashflowHttp_DeleteData(t *testing.T) {
	mockUsecase := new(MockCashflowUsecase)
	handler := &CashflowHttp{cashflowUsecase: mockUsecase}

	app := fiber.New()
	app.Delete("/cashflow/:id", handler.DeleteData)

	mockUsecase.On("Delete", mock.Anything, 1).Return(nil)

	req := httptest.NewRequest("DELETE", "/cashflow/1", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	mockUsecase.AssertExpectations(t)
}
