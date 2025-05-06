package http

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockReferalCodeUsecase struct {
	mock.Mock
}

func (m *MockReferalCodeUsecase) Fetch(ctx *fiber.Ctx) error {
	args := m.Called(ctx)
	return args.Error(0)
}

func (m *MockReferalCodeUsecase) FetchById(ctx *fiber.Ctx, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockReferalCodeUsecase) FetchByRTProfileId(ctx *fiber.Ctx, rtProfileId string) error {
	args := m.Called(ctx, rtProfileId)
	return args.Error(0)
}

func (m *MockReferalCodeUsecase) RegenerateReferalCode(ctx *fiber.Ctx, profileId string, code string) error {
	args := m.Called(ctx, profileId, code)
	return args.Error(0)
}

func (m *MockReferalCodeUsecase) Validate(ctx *fiber.Ctx, code string) (error, string) {
	args := m.Called(ctx, code)
	return args.Error(0), args.String(1)
}

func (m *MockReferalCodeUsecase) Delete(ctx *fiber.Ctx, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func TestReferalCodeHttp_GetData(t *testing.T) {
	app := fiber.New()
	mockUsecase := new(MockReferalCodeUsecase)
	handler := &ReferalCodeHttp{referalCodeUsecase: mockUsecase}

	mockUsecase.On("Fetch", mock.Anything).Return(nil)

	app.Get("/", handler.GetData)

	req := httptest.NewRequest("GET", "/", nil)
	res, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, 200, res.StatusCode)
	mockUsecase.AssertExpectations(t)
}

func TestReferalCodeHttp_GetDataById(t *testing.T) {
	app := fiber.New()
	mockUsecase := new(MockReferalCodeUsecase)
	handler := &ReferalCodeHttp{referalCodeUsecase: mockUsecase}

	mockUsecase.On("FetchById", mock.Anything, 123).Return(nil)

	app.Get("/id/:id", handler.GetDataById)

	req := httptest.NewRequest("GET", "/id/123", nil)
	res, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, 200, res.StatusCode)
	mockUsecase.AssertExpectations(t)
}

func TestReferalCodeHttp_GetDataByRTProfileId(t *testing.T) {
	app := fiber.New()
	mockUsecase := new(MockReferalCodeUsecase)
	handler := &ReferalCodeHttp{referalCodeUsecase: mockUsecase}

	mockUsecase.On("FetchByRTProfileId", mock.Anything, "10").Return(nil)

	app.Get("/profile/:profile_id", handler.GetDataByRTProfileId)

	req := httptest.NewRequest("GET", "/profile/10", nil)
	res, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, 200, res.StatusCode)
	mockUsecase.AssertExpectations(t)
}

func TestReferalCodeHttp_RegenerateCode(t *testing.T) {
	app := fiber.New()
	mockUsecase := new(MockReferalCodeUsecase)
	handler := &ReferalCodeHttp{referalCodeUsecase: mockUsecase}

	mockUsecase.On("RegenerateReferalCode", mock.Anything, "20", "ABC123").Return(nil)

	app.Put("/regen/:profile_id/:code", handler.RegenerateCode)

	req := httptest.NewRequest("PUT", "/regen/20/ABC123", nil)
	res, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, 200, res.StatusCode)
	mockUsecase.AssertExpectations(t)
}

func TestReferalCodeHttp_ValidateReferalCode(t *testing.T) {
	app := fiber.New()
	mockUsecase := new(MockReferalCodeUsecase)
	handler := &ReferalCodeHttp{referalCodeUsecase: mockUsecase}

	mockUsecase.On("Validate", mock.Anything, "REF123").Return(nil, "valid")

	app.Get("/validate", handler.ValidateReferalCode)

	req := httptest.NewRequest("GET", "/validate?code=REF123", nil)
	res, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, 200, res.StatusCode)
	mockUsecase.AssertExpectations(t)
}

func TestReferalCodeHttp_DeleteData(t *testing.T) {
	app := fiber.New()
	mockUsecase := new(MockReferalCodeUsecase)
	handler := &ReferalCodeHttp{referalCodeUsecase: mockUsecase}

	mockUsecase.On("Delete", mock.Anything, 123).Return(nil)

	app.Delete("/:id", handler.DeleteData)

	req := httptest.NewRequest("DELETE", "/123", nil)
	res, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, 200, res.StatusCode)
	mockUsecase.AssertExpectations(t)
}
