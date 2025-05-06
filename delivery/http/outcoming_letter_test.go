package http

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockOutcomingLetterUsecase struct {
	mock.Mock
}

func (m *MockOutcomingLetterUsecase) Fetch(c *fiber.Ctx) error {
	return m.Called(c).Error(0)
}

func (m *MockOutcomingLetterUsecase) FetchPreview(c *fiber.Ctx, id string) error {
	return m.Called(c, id).Error(0)
}

func (m *MockOutcomingLetterUsecase) FetchById(c *fiber.Ctx, id string) error {
	return m.Called(c, id).Error(0)
}

func (m *MockOutcomingLetterUsecase) FetchByRTProfileId(c *fiber.Ctx, id string) error {
	return m.Called(c, id).Error(0)
}

func (m *MockOutcomingLetterUsecase) Delete(c *fiber.Ctx, id string) error {
	return m.Called(c, id).Error(0)
}

func TestOutcomingLetterHttp_GetData(t *testing.T) {
	app := fiber.New()
	mockUC := new(MockOutcomingLetterUsecase)
	handler := NewOutcomingLetterHttp(mockUC)

	app.Get("/outcoming-letters", handler.GetData)

	mockUC.On("Fetch", mock.Anything).Return(nil)

	req := httptest.NewRequest("GET", "/outcoming-letters", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	mockUC.AssertCalled(t, "Fetch", mock.Anything)
}

func TestOutcomingLetterHttp_GetPreview(t *testing.T) {
	app := fiber.New()
	mockUC := new(MockOutcomingLetterUsecase)
	handler := NewOutcomingLetterHttp(mockUC)

	app.Get("/outcoming-letters/preview/:id", handler.GetPreview)

	mockUC.On("FetchPreview", mock.Anything, "abc").Return(nil)

	req := httptest.NewRequest("GET", "/outcoming-letters/preview/abc", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	mockUC.AssertCalled(t, "FetchPreview", mock.Anything, "abc")
}

func TestOutcomingLetterHttp_GetDataById(t *testing.T) {
	app := fiber.New()
	mockUC := new(MockOutcomingLetterUsecase)
	handler := NewOutcomingLetterHttp(mockUC)

	app.Get("/outcoming-letters/:id", handler.GetDataById)

	mockUC.On("FetchById", mock.Anything, "xyz").Return(nil)

	req := httptest.NewRequest("GET", "/outcoming-letters/xyz", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	mockUC.AssertCalled(t, "FetchById", mock.Anything, "xyz")
}

func TestOutcomingLetterHttp_GetDataByRTProfileId(t *testing.T) {
	app := fiber.New()
	mockUC := new(MockOutcomingLetterUsecase)
	handler := NewOutcomingLetterHttp(mockUC)

	app.Get("/outcoming-letters/rt/:rt_profile_id", handler.GetDataByRTProfileId)

	mockUC.On("FetchByRTProfileId", mock.Anything, "999").Return(nil)

	req := httptest.NewRequest("GET", "/outcoming-letters/rt/999", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	mockUC.AssertCalled(t, "FetchByRTProfileId", mock.Anything, "999")
}

func TestOutcomingLetterHttp_DeleteData(t *testing.T) {
	app := fiber.New()
	mockUC := new(MockOutcomingLetterUsecase)
	handler := NewOutcomingLetterHttp(mockUC)

	app.Delete("/outcoming-letters/:id", handler.DeleteData)

	mockUC.On("Delete", mock.Anything, "del-id").Return(nil)

	req := httptest.NewRequest("DELETE", "/outcoming-letters/del-id", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	mockUC.AssertCalled(t, "Delete", mock.Anything, "del-id")
}
