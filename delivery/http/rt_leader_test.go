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

type MockRTLeaderUsecase struct {
	mock.Mock
}

func (m *MockRTLeaderUsecase) Fetch(c *fiber.Ctx) error {
	return m.Called(c).Error(0)
}

func (m *MockRTLeaderUsecase) FetchById(c *fiber.Ctx, id string) error {
	return m.Called(c, id).Error(0)
}
func (m *MockRTLeaderUsecase) FetchByRTProfileId(c *fiber.Ctx, id string) error {
	return m.Called(c, id).Error(0)
}

func (m *MockRTLeaderUsecase) Update(c *fiber.Ctx, id string, data *model.RTLeader) error {
	return m.Called(c, id, data).Error(0)
}

func TestRTLeaderHttp_GetData_NoParam(t *testing.T) {
	app := fiber.New()
	mockUC := new(MockRTLeaderUsecase)
	handler := NewRTLeaderHttp(mockUC)

	app.Get("/rt-leaders", handler.GetData)

	mockUC.On("Fetch", mock.Anything).Return(nil)

	req := httptest.NewRequest("GET", "/rt-leaders", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	mockUC.AssertCalled(t, "Fetch", mock.Anything)
}

func TestRTLeaderHttp_GetData_WithRTProfileId(t *testing.T) {
	app := fiber.New()
	mockUC := new(MockRTLeaderUsecase)
	handler := NewRTLeaderHttp(mockUC)

	app.Get("/rt-leaders/:rt_profile_id", handler.GetData)

	mockUC.On("FetchByRTProfileId", mock.Anything, "abc123").Return(nil)

	req := httptest.NewRequest("GET", "/rt-leaders/abc123", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	mockUC.AssertCalled(t, "FetchByRTProfileId", mock.Anything, "abc123")
}

func TestRTLeaderHttp_GetDataById(t *testing.T) {
	app := fiber.New()
	mockUC := new(MockRTLeaderUsecase)
	handler := NewRTLeaderHttp(mockUC)

	app.Get("/rt-leaders/detail/:id", handler.GetDataById)

	mockUC.On("FetchByRTProfileId", mock.Anything, "xyz456").Return(nil)

	req := httptest.NewRequest("GET", "/rt-leaders/detail/xyz456", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	mockUC.AssertCalled(t, "FetchByRTProfileId", mock.Anything, "xyz456")
}

func TestRTLeaderHttp_UpdateData_Success(t *testing.T) {
	app := fiber.New()
	mockUC := new(MockRTLeaderUsecase)
	handler := NewRTLeaderHttp(mockUC)

	app.Put("/rt-leaders/:id", handler.UpdateData)

	jsonStr := []byte(`{"name":"Test Leader"}`)

	mockUC.On("Update", mock.Anything, "id123", mock.AnythingOfType("*model.RTLeader")).Return(nil)

	req := httptest.NewRequest("PUT", "/rt-leaders/id123", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	resp, _ := app.Test(req)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	mockUC.AssertCalled(t, "Update", mock.Anything, "id123", mock.AnythingOfType("*model.RTLeader"))
}

func TestRTLeaderHttp_UpdateData_BodyParserError(t *testing.T) {
	app := fiber.New()
	mockUC := new(MockRTLeaderUsecase)
	handler := NewRTLeaderHttp(mockUC)

	app.Put("/rt-leaders/:id", handler.UpdateData)

	req := httptest.NewRequest("PUT", "/rt-leaders/id123", bytes.NewBuffer([]byte(`invalid-json`)))
	req.Header.Set("Content-Type", "application/json")

	resp, _ := app.Test(req)
	assert.Equal(t, fiber.StatusUnprocessableEntity, resp.StatusCode)
	mockUC.AssertNotCalled(t, "Update")
}
