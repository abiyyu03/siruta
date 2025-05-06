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

// MockMemberUsecase implements the mocked usecase
type MockMemberUsecase struct {
	mock.Mock
}

func (m *MockMemberUsecase) Fetch(ctx *fiber.Ctx) error {
	args := m.Called(ctx)
	return args.Error(0)
}

func (m *MockMemberUsecase) FetchById(ctx *fiber.Ctx, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockMemberUsecase) FetchByRTProfileId(ctx *fiber.Ctx, rtProfileId string) error {
	args := m.Called(ctx, rtProfileId)
	return args.Error(0)
}

func (m *MockMemberUsecase) Update(ctx *fiber.Ctx, id string, data *model.Member) error {
	args := m.Called(ctx, id, data)
	return args.Error(0)
}

// -------------------- TEST CASES --------------------

func TestMemberHttp_GetData(t *testing.T) {
	app := fiber.New()
	mockUsecase := new(MockMemberUsecase)
	handler := &MemberHttp{memberUsecase: mockUsecase}

	app.Get("/members", handler.GetData)

	mockUsecase.On("Fetch", mock.Anything).Return(nil)

	req := httptest.NewRequest("GET", "/members", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	mockUsecase.AssertExpectations(t)
}

func TestMemberHttp_GetDataById(t *testing.T) {
	app := fiber.New()
	mockUsecase := new(MockMemberUsecase)
	handler := &MemberHttp{memberUsecase: mockUsecase}

	app.Get("/members/:id", handler.GetDataById)

	mockUsecase.On("FetchById", mock.Anything, "1").Return(nil)

	req := httptest.NewRequest("GET", "/members/1", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	mockUsecase.AssertExpectations(t)
}

func TestMemberHttp_GetDataByRTProfileId(t *testing.T) {
	app := fiber.New()
	mockUsecase := new(MockMemberUsecase)
	handler := &MemberHttp{memberUsecase: mockUsecase}

	app.Get("/members/rt-profile/:rt_profile_id", handler.GetDataByRTProfileId)

	mockUsecase.On("FetchByRTProfileId", mock.Anything, "123").Return(nil)

	req := httptest.NewRequest("GET", "/members/rt-profile/123", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	mockUsecase.AssertExpectations(t)
}

func TestMemberHttp_UpdateData(t *testing.T) {
	app := fiber.New()
	mockUsecase := new(MockMemberUsecase)
	handler := &MemberHttp{memberUsecase: mockUsecase}

	app.Put("/members/:id", handler.UpdateData)

	memberData := &model.Member{
		Fullname:       "Budi Santoso",
		BornPlace:      "Jakarta",
		BirthDate:      "2023-12-12",
		Gender:         "male",
		ReligionId:     1,
		MemberStatusId: 2,
		UserId:         "usear023",
		Status:         "active",
		RTProfileId:    "rt-21321321",
	}

	body, _ := json.Marshal(memberData)

	mockUsecase.On("Update", mock.Anything, "1", mock.MatchedBy(func(m *model.Member) bool {
		return m.Fullname == "Budi Santoso" && m.Status == "active"
	})).Return(nil)

	req := httptest.NewRequest("PUT", "/members/1", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)

	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	mockUsecase.AssertExpectations(t)
}
