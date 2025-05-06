package http

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Define a mock RWProfileUsecase to simulate the behavior of the usecase layer
type MockRWProfileUsecase struct {
	mock.Mock
}

func (m *MockRWProfileUsecase) Fetch(ctx *fiber.Ctx) error {
	args := m.Called(ctx)
	return args.Error(0)
}

func (m *MockRWProfileUsecase) FetchById(ctx *fiber.Ctx, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

// Test the GetData method
func TestGetData(t *testing.T) {
	// Set up the mock usecase
	mockUsecase := new(MockRWProfileUsecase)
	handler := NewRWProfileHttp(mockUsecase)

	// Set up the mock to expect a call to Fetch
	mockUsecase.On("Fetch", mock.Anything).Return(nil)

	// Create a new Fiber app and define a route for testing
	app := fiber.New()
	app.Get("/rwprofile", handler.GetData)

	// Send a GET request to the /rwprofile route
	req := httptest.NewRequest("GET", "/rwprofile", nil)
	resp, err := app.Test(req)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	mockUsecase.AssertExpectations(t)
}

// Test the GetDataById method
func TestGetDataById(t *testing.T) {
	// Set up the mock usecase
	mockUsecase := new(MockRWProfileUsecase)
	handler := NewRWProfileHttp(mockUsecase)

	// Set up the mock to expect a call to FetchById
	mockUsecase.On("FetchById", mock.Anything, "1").Return(nil)

	// Create a new Fiber app and define a route for testing
	app := fiber.New()
	app.Get("/rwprofile/:id", handler.GetDataById)

	// Send a GET request to the /rwprofile/:id route
	req := httptest.NewRequest("GET", "/rwprofile/1", nil)
	resp, err := app.Test(req)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	mockUsecase.AssertExpectations(t)
}
