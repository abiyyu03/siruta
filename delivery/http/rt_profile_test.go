package http

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Define a mock Rt-ProfilesUsecase to simulate the behavior of the usecase layer
type MockRTProfilesUsecase struct {
	mock.Mock
}

func (m *MockRTProfilesUsecase) Fetch(ctx *fiber.Ctx) error {
	return m.Called(ctx).Error(0)
}

func (m *MockRTProfilesUsecase) FetchById(ctx *fiber.Ctx, id string) error {
	return m.Called(ctx, id).Error(0)
}

func (m *MockRTProfilesUsecase) FetchByRWProfileId(ctx *fiber.Ctx, rwProfileId string) error {
	return m.Called(ctx, rwProfileId).Error(0)
}

// Test the GetData method
func TestRTProfile_GetData(t *testing.T) {
	// Set up the mock usecase
	mockUsecase := new(MockRTProfilesUsecase)
	handler := NewRTProfileHttp(mockUsecase)

	// Set up the mock to expect a call to Fetch
	mockUsecase.On("Fetch", mock.Anything).Return(nil)

	// Create a new Fiber app and define a route for testing
	app := fiber.New()
	app.Get("/rt-profiles", handler.GetData)

	// Send a GET request to the /rt-profiles route
	req := httptest.NewRequest("GET", "/rt-profiles", nil)
	resp, err := app.Test(req)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	mockUsecase.AssertExpectations(t)
}

// Test the GetDataById method
func TestRTProfile_GetDataById(t *testing.T) {
	// Set up the mock usecase
	mockUsecase := new(MockRTProfilesUsecase)
	handler := NewRTProfileHttp(mockUsecase)

	// Set up the mock to expect a call to FetchById
	mockUsecase.On("FetchById", mock.Anything, "user-123213").Return(nil)

	// Create a new Fiber app and define a route for testing
	app := fiber.New()
	app.Get("/rt-profiles/:id", handler.GetDataById)

	// Send a GET request to the /rt-profiles/:id route
	req := httptest.NewRequest("GET", "/rt-profiles/user-123213", nil)
	resp, err := app.Test(req)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	mockUsecase.AssertExpectations(t)
}

func TestRTProfile_FetchByRWProfileId(t *testing.T) {
	// Set up the mock usecase
	mockUsecase := new(MockRTProfilesUsecase)
	handler := NewRTProfileHttp(mockUsecase)

	// Set up the mock to expect a call to FetchById
	mockUsecase.On("FetchByRWProfileId", mock.Anything, "user").Return(nil)

	// Create a new Fiber app and define a route for testing
	app := fiber.New()
	app.Get("/rt-profiles/:rw_profile_id/rw", handler.GetDataByRWProfileId)

	// Send a GET request to the /rt-profiles/:id route
	req := httptest.NewRequest("GET", "/rt-profiles/user/rw", nil)
	resp, err := app.Test(req)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	mockUsecase.AssertExpectations(t)
}
