package http

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/abiyyu03/siruta/entity/model"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockIncomingLetterUsecase struct {
	mock.Mock
}

func (m *MockIncomingLetterUsecase) Fetch(ctx *fiber.Ctx) error {
	args := m.Called(ctx)
	return args.Error(0)
}
func (m *MockIncomingLetterUsecase) FetchByRTProfileId(ctx *fiber.Ctx, rtProfileId string) error {
	args := m.Called(ctx, rtProfileId)
	return args.Error(0)
}
func (m *MockIncomingLetterUsecase) FetchById(ctx *fiber.Ctx, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}
func (m *MockIncomingLetterUsecase) Store(incomingLetter *model.IncomingLetter, ctx *fiber.Ctx) error {
	args := m.Called(ctx, incomingLetter)
	return args.Error(0)
}
func (m *MockIncomingLetterUsecase) Update(incomingLetter *model.IncomingLetter, ctx *fiber.Ctx, id int) error {
	args := m.Called(incomingLetter, ctx, id)
	return args.Error(0)
}
func (m *MockIncomingLetterUsecase) Delete(ctx *fiber.Ctx, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func TestIncomingLetterHttp_GetData(t *testing.T) {
	app := fiber.New()
	mockUsecase := new(MockIncomingLetterUsecase)
	handler := &IncomingLetterHttp{incomingLetterUsecase: mockUsecase}

	app.Get("/incoming-letters", handler.GetData)

	mockUsecase.On("Fetch", mock.Anything).Return(nil)

	req := httptest.NewRequest("GET", "/incoming-letters", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, 200, resp.StatusCode)
	mockUsecase.AssertExpectations(t)
}

func TestIncomingLetterHttp_GetDataByRTProfileId(t *testing.T) {
	mockUsecase := new(MockIncomingLetterUsecase)
	handler := &IncomingLetterHttp{incomingLetterUsecase: mockUsecase}

	app := fiber.New()
	app.Get("/incoming-letters/rt/:rt_profile_id", handler.GetDataByRTProfileId)

	mockUsecase.On("FetchByRTProfileId", mock.Anything, "1").Return(nil)

	req := httptest.NewRequest("GET", "/incoming-letters/rt/1", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	mockUsecase.AssertExpectations(t)
}

func TestIncomingLetterHttp_GetDataById(t *testing.T) {
	mockUsecase := new(MockIncomingLetterUsecase)
	handler := &IncomingLetterHttp{incomingLetterUsecase: mockUsecase}

	app := fiber.New()
	app.Get("/incoming-letters/:id", handler.GetDataById)

	mockUsecase.On("FetchById", mock.Anything, 1).Return(nil)

	req := httptest.NewRequest("GET", "/incoming-letters/1", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	mockUsecase.AssertExpectations(t)
}
func TestIncomingLetterHttp_UpdateData(t *testing.T) {
	mockUsecase := new(MockIncomingLetterUsecase)
	handler := &IncomingLetterHttp{incomingLetterUsecase: mockUsecase}

	app := fiber.New()
	app.Put("/incoming-letters/:id", handler.UpdateData)

	updatedBody := &model.IncomingLetter{
		Title:        "Updated Title",
		LetterDate:   time.Now(),
		OriginLetter: "KPK",
		RTProfileId:  "user-999",
	}
	jsonBody, _ := json.Marshal(updatedBody)

	mockUsecase.On("Update", mock.MatchedBy(func(data *model.IncomingLetter) bool {
		return data.Title == "Updated Title"
	}), mock.Anything, 1).Return(nil)

	req := httptest.NewRequest("PUT", "/incoming-letters/1", bytes.NewReader(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)

	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	mockUsecase.AssertExpectations(t)
}

func TestIncomingLetterHttp_StoreData(t *testing.T) {
	mockUsecase := new(MockIncomingLetterUsecase)
	handler := &IncomingLetterHttp{incomingLetterUsecase: mockUsecase}

	app := fiber.New()
	app.Post("/incoming-letters", handler.StoreData)

	incomingLetter := &model.IncomingLetter{
		Title:        "Test Surat",
		LetterDate:   time.Now(),
		OriginLetter: "KPK",
		RTProfileId:  "user-123-321",
	}
	body, _ := json.Marshal(incomingLetter)

	mockUsecase.On("Store", mock.Anything, mock.MatchedBy(func(cf *model.IncomingLetter) bool {
		return cf.Title == "Test Surat"
	})).Return(nil)

	req := httptest.NewRequest("POST", "/incoming-letters", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	mockUsecase.AssertExpectations(t)
}

func TestIncomingLetterHttp_DeleteData(t *testing.T) {
	mockUsecase := new(MockIncomingLetterUsecase)
	handler := &IncomingLetterHttp{incomingLetterUsecase: mockUsecase}

	app := fiber.New()
	app.Delete("/incoming-letters/:id", handler.DeleteData)

	mockUsecase.On("Delete", mock.Anything, 1).Return(nil)

	req := httptest.NewRequest("DELETE", "/incoming-letters/1", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	mockUsecase.AssertExpectations(t)
}
