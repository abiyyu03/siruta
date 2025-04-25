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

type MockLetterTypeUsecase struct {
	mock.Mock
}

func (m *MockLetterTypeUsecase) Fetch(ctx *fiber.Ctx) error {
	args := m.Called(ctx)
	return args.Error(0)
}
func (m *MockLetterTypeUsecase) FetchById(ctx *fiber.Ctx, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}
func (m *MockLetterTypeUsecase) Store(letterType *model.LetterType, ctx *fiber.Ctx) error {
	args := m.Called(ctx, letterType)
	return args.Error(0)
}
func (m *MockLetterTypeUsecase) Update(letterType *model.LetterType, ctx *fiber.Ctx, id int) error {
	args := m.Called(ctx, letterType, id)
	return args.Error(0)
}
func (m *MockLetterTypeUsecase) Delete(ctx *fiber.Ctx, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func TestLetterTypeHttp_GetData(t *testing.T) {
	app := fiber.New()
	mockUsecase := new(MockLetterTypeUsecase)
	handler := &LetterTypeHttp{letterTypeUsecase: mockUsecase}

	app.Get("/letter-types", handler.GetData)

	mockUsecase.On("Fetch", mock.Anything).Return(nil)

	req := httptest.NewRequest("GET", "/letter-types", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, 200, resp.StatusCode)
	mockUsecase.AssertExpectations(t)
}

func TestLetterTypeHttp_GetDataById(t *testing.T) {
	mockUsecase := new(MockLetterTypeUsecase)
	handler := &LetterTypeHttp{letterTypeUsecase: mockUsecase}

	app := fiber.New()
	app.Get("/letter-types/:id", handler.GetDataById)

	mockUsecase.On("FetchById", mock.Anything, 1).Return(nil)

	req := httptest.NewRequest("GET", "/letter-types/1", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	mockUsecase.AssertExpectations(t)
}

func TestLetterTypeHttp_StoreData(t *testing.T) {
	mockUsecase := new(MockLetterTypeUsecase)
	handler := &LetterTypeHttp{letterTypeUsecase: mockUsecase}

	app := fiber.New()
	app.Post("/letter-types", handler.StoreData)

	letterType := &model.LetterType{
		TypeName:           "Surat Pengantar Nikah",
		IsForLocalResident: false,
		Code:               "SPN",
	}
	body, _ := json.Marshal(letterType)

	mockUsecase.On("Store", mock.Anything, mock.MatchedBy(func(cf *model.LetterType) bool {
		return cf.Code == "SPN"
	})).Return(nil)

	req := httptest.NewRequest("POST", "/letter-types", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	mockUsecase.AssertExpectations(t)
}

func TestLetterTypeHttp_UpdateData(t *testing.T) {
	mockUsecase := new(MockLetterTypeUsecase)
	handler := &LetterTypeHttp{letterTypeUsecase: mockUsecase}
	app := fiber.New()

	app.Put("/letter-types/:id", handler.UpdateData)

	body := &model.LetterType{
		ID:                 1,
		TypeName:           "Surat Pengantar KTP",
		IsForLocalResident: false,
		Code:               "SPK",
	}
	updatedBody, _ := json.Marshal(body)

	mockUsecase.On("Update", mock.MatchedBy(func(data *model.LetterType) bool {
		return data.TypeName == "Komputer High End"
	}), mock.Anything, 1).Return(nil)

	req := httptest.NewRequest("PUT", "/letter-types/1", bytes.NewReader(updatedBody))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	mockUsecase.AssertExpectations(t)
}

func TestLetterTypeHttp_DeleteData(t *testing.T) {
	mockUsecase := new(MockLetterTypeUsecase)
	handler := &LetterTypeHttp{letterTypeUsecase: mockUsecase}

	app := fiber.New()
	app.Delete("/letter-types/:id", handler.DeleteData)

	mockUsecase.On("Delete", mock.Anything, 1).Return(nil)

	req := httptest.NewRequest("DELETE", "/letter-types/1", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	mockUsecase.AssertExpectations(t)
}
