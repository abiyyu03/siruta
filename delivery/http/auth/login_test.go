package auth

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockAuthUsecase is a mock implementation of the authentication use case.
type MockAuthUsecase struct {
	mock.Mock
}

func (m *MockAuthUsecase) IssueAuthToken(ctx *fiber.Ctx, email, password string) (string, error) {
	args := m.Called(ctx, email, password)
	return args.String(0), args.Error(1)
}

// Function to load private key
func loadPrivateKey(relativePath string) ([]byte, error) {
	basePath, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	absolutePath := filepath.Join(basePath, relativePath)
	return ioutil.ReadFile(absolutePath)
}

func TestLoginHandler(t *testing.T) {
	// Load private key (Adjust the path if necessary)
	privateKeyPath := "../../../keys/private.pem"
	privateKey, err := loadPrivateKey(privateKeyPath)
	if err != nil {
		t.Fatalf("Failed to load private key: %v", err)
	}
	log.Printf("Private key loaded successfully: %s", privateKeyPath)

	tests := []struct {
		name            string
		inputBody       string
		mockSetup       func(mockAuth *MockAuthUsecase)
		expectedCode    int
		expectedMessage string
	}{
		{
			name:      "Successful login",
			inputBody: `{"email": "test@example.com", "password": "password123"}`,
			mockSetup: func(mockAuth *MockAuthUsecase) {
				mockAuth.On("IssueAuthToken", mock.Anything, "test@example.com", "password123").Return("mocked_token", nil)
			},
			expectedCode:    http.StatusOK,
			expectedMessage: "login successfully",
		},
		{
			name:      "Invalid credentials",
			inputBody: `{"email": "test@example.com", "password": "wrongpassword"}`,
			mockSetup: func(mockAuth *MockAuthUsecase) {
				mockAuth.On("IssueAuthToken", mock.Anything, "test@example.com", "wrongpassword").Return("", assert.AnError)
			},
			expectedCode:    http.StatusBadRequest,
			expectedMessage: "Bad Request",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := fiber.New()
			mockAuth := new(MockAuthUsecase)

			// Set up the mock behavior
			if tt.mockSetup != nil {
				tt.mockSetup(mockAuth)
			}

			// Register the login handler
			app.Post("/login", func(ctx *fiber.Ctx) error {
				var request struct {
					Email    string `json:"email"`    // Corrected field name to start with uppercase
					Password string `json:"password"` // Corrected field name to start with uppercase
				}
				if err := ctx.BodyParser(&request); err != nil {
					return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
						"message": fiber.ErrBadRequest.Message,
					})
				}

				// Use private key for signing (log for debugging)
				log.Printf("Using private key for signing: %s", privateKey)

				generatedToken, err := mockAuth.IssueAuthToken(ctx, request.Email, request.Password) // Use correct field names
				if err != nil {
					return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
						"message": fiber.ErrBadRequest.Message,
					})
				}

				return ctx.JSON(fiber.Map{
					"message": "login successfully",
					"token":   generatedToken,
				})
			})

			// Create a test request
			req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(tt.inputBody))
			req.Header.Set("Content-Type", "application/json")
			resp, _ := app.Test(req)

			assert.Equal(t, tt.expectedCode, resp.StatusCode)

			// Parse response body
			var response map[string]interface{}
			json.NewDecoder(resp.Body).Decode(&response)
			assert.Contains(t, response["message"], tt.expectedMessage)
		})
	}
}
