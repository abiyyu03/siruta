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

type MockGuestListUsecase struct {
	mock.Mock
}

func (m *MockGuestListUsecase) Fetch(ctx *fiber.Ctx) error {
	args := m.Called(ctx)
	return args.Error(0)
}
func (m *MockGuestListUsecase) FetchByRTProfileId(ctx *fiber.Ctx, rtProfileId string) error {
	args := m.Called(ctx, rtProfileId)
	return args.Error(0)
}
func (m *MockGuestListUsecase) FetchById(ctx *fiber.Ctx, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}
func (m *MockGuestListUsecase) Store(ctx *fiber.Ctx, guestList *model.GuestList) error {
	args := m.Called(ctx, guestList)
	return args.Error(0)
}
func (m *MockGuestListUsecase) Update(ctx *fiber.Ctx, guestList *model.GuestList, id int) error {
	args := m.Called(ctx, guestList, id)
	return args.Error(0)
}
func (m *MockGuestListUsecase) Delete(ctx *fiber.Ctx, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func TestGuestListHttp_GetData(t *testing.T) {
	app := fiber.New()
	mockUsecase := new(MockGuestListUsecase)
	handler := &GuestListHttp{guestListUsecase: mockUsecase}

	app.Get("/guest-lists", handler.GetData)

	mockUsecase.On("Fetch", mock.Anything).Return(nil)

	req := httptest.NewRequest("GET", "/guest-lists", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, 200, resp.StatusCode)
	mockUsecase.AssertExpectations(t)
}

func TestGuestListHttp_GetDataByRTProfileId(t *testing.T) {
	mockUsecase := new(MockGuestListUsecase)
	handler := &GuestListHttp{guestListUsecase: mockUsecase}

	app := fiber.New()
	app.Get("/guest-lists/rt/:rt_profile_id", handler.GetDataByRTProfileId)

	mockUsecase.On("FetchByRTProfileId", mock.Anything, "1").Return(nil)

	req := httptest.NewRequest("GET", "/guest-lists/rt/1", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	mockUsecase.AssertExpectations(t)
}

func TestGuestListHttp_StoreData(t *testing.T) {
	mockUsecase := new(MockGuestListUsecase)
	handler := &GuestListHttp{guestListUsecase: mockUsecase}

	app := fiber.New()
	app.Post("/guest-lists", handler.CreateData)

	guestList := &model.GuestList{FullName: "Abiyyu", PhoneNumber: "08123123", VisitAt: time.Now(), RTProfileId: "user-222-123"}
	body, _ := json.Marshal(guestList)

	mockUsecase.On("Store", mock.Anything, mock.MatchedBy(func(cf *model.GuestList) bool {
		return cf.FullName == "Abiyyu"
	})).Return(nil)

	req := httptest.NewRequest("POST", "/guest-lists", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	mockUsecase.AssertExpectations(t)
}

func TestGuestListHttp_DeleteData(t *testing.T) {
	mockUsecase := new(MockGuestListUsecase)
	handler := &GuestListHttp{guestListUsecase: mockUsecase}

	app := fiber.New()
	app.Delete("/guest-lists/:id", handler.DeleteData)

	mockUsecase.On("Delete", mock.Anything, 1).Return(nil)

	req := httptest.NewRequest("DELETE", "/guest-lists/1", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	mockUsecase.AssertExpectations(t)
}
