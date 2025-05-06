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

type MockRWLeaderUsecase struct {
	mock.Mock
}

func (m *MockRWLeaderUsecase) Fetch(c *fiber.Ctx) error {
	return m.Called(c).Error(0)
}

func (m *MockRWLeaderUsecase) FetchById(c *fiber.Ctx, id string) error {
	return m.Called(c, id).Error(0)
}

func (m *MockRWLeaderUsecase) Update(c *fiber.Ctx, id string, data *model.RWLeader) error {
	return m.Called(c, id, data).Error(0)
}

func TestRWLeaderHttp_GetData_NoParam(t *testing.T) {
	app := fiber.New()
	mockUC := new(MockRWLeaderUsecase)
	handler := NewRWLeaderHttp(mockUC)

	app.Get("/rw-leaders", handler.GetData)

	mockUC.On("Fetch", mock.Anything).Return(nil)

	req := httptest.NewRequest("GET", "/rw-leaders", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	mockUC.AssertCalled(t, "Fetch", mock.Anything)
}

func TestRWLeaderHttp_GetDataById(t *testing.T) {
	app := fiber.New()
	mockUC := new(MockRWLeaderUsecase)
	handler := NewRWLeaderHttp(mockUC)

	app.Get("/rw-leaders/:id", handler.GetDataById)

	mockUC.On("FetchById", mock.Anything, "xyz456").Return(nil)

	req := httptest.NewRequest("GET", "/rw-leaders/xyz456", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	mockUC.AssertCalled(t, "FetchById", mock.Anything, "xyz456")
}

func TestRWLeaderHttp_UpdateData_Success(t *testing.T) {
	app := fiber.New()
	mockUC := new(MockRWLeaderUsecase)
	handler := NewRWLeaderHttp(mockUC)

	app.Put("/rw-leaders/:id", handler.UpdateData)

	jsonStr := []byte(`{"name":"Test Leader"}`)

	mockUC.On("Update", mock.Anything, "id123", mock.AnythingOfType("*model.RWLeader")).Return(nil)

	req := httptest.NewRequest("PUT", "/rw-leaders/id123", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	resp, _ := app.Test(req)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	mockUC.AssertCalled(t, "Update", mock.Anything, "id123", mock.AnythingOfType("*model.RWLeader"))
}

func TestRWLeaderHttp_UpdateData_BodyParserError(t *testing.T) {
	app := fiber.New()
	mockUC := new(MockRWLeaderUsecase)
	handler := NewRWLeaderHttp(mockUC)

	app.Put("/rw-leaders/:id", handler.UpdateData)

	req := httptest.NewRequest("PUT", "/rw-leaders/id123", bytes.NewBuffer([]byte(`invalid-json`)))
	req.Header.Set("Content-Type", "application/json")

	resp, _ := app.Test(req)
	assert.Equal(t, fiber.StatusUnprocessableEntity, resp.StatusCode)
	mockUC.AssertNotCalled(t, "Update")
}
