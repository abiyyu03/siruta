package http_test

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/abiyyu03/siruta/delivery/http"
	"github.com/abiyyu03/siruta/entity/model"
	"github.com/abiyyu03/siruta/entity/request"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock untuk LetterReqUsecaseInterface
type MockLetterReqUsecase struct {
	mock.Mock
}

func (m *MockLetterReqUsecase) StoreOutcomingLetter(ctx *fiber.Ctx, member *model.Member, letter *model.OutcomingLetter, memberStatus string, birthDate string, nik string) error {
	args := m.Called(ctx, member, letter, memberStatus, birthDate, nik)
	return args.Error(0)
}

func (m *MockLetterReqUsecase) UpdateApprovalStatusByRT(ctx *fiber.Ctx, letterReqId string) error {
	args := m.Called(ctx, letterReqId)
	return args.Error(0)
}

func TestLetterReqHttp_CreateData(t *testing.T) {
	app := fiber.New()
	mockUsecase := new(MockLetterReqUsecase)
	handler := http.NewLetterReqHttp(mockUsecase)

	app.Post("/letter-requests", handler.CreateData)

	// Sample data
	member := &model.Member{
		Fullname:       "John Doe",
		BornPlace:      "Depok",
		BirthDate:      "1990-01-01",
		Gender:         "L",
		ReligionId:     1,
		MemberStatusId: 2,
		UserId:         "user-12321",
		Status:         "active",
		RTProfileId:    "rtprofile-123123",
	}

	letter := &model.OutcomingLetter{
		LetterTypeId: 1,
		Date:         "2025-01-01",
	}

	check := &request.CheckResidentMember{
		MemberStatus: "tetap",
		BirthDate:    "1990-01-01",
		NikNumber:    "3201010101010001",
	}

	body := http.CombinedRequest{
		Member:          member,
		OutcomingLetter: letter,
		CheckResident:   check,
	}

	jsonBody, _ := json.Marshal(body)

	mockUsecase.
		On("StoreOutcomingLetter", mock.Anything, member, letter, check.MemberStatus, check.BirthDate, check.NikNumber).
		Return(nil)

	req := httptest.NewRequest("POST", "/letter-requests", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	mockUsecase.AssertExpectations(t)
}
