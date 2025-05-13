package entity

import (
	"github.com/gofiber/fiber/v2"
)

//------------------------------------------------------------------------------
// Standard Response

type Response struct {
	Status      string      `json:"status"`
	Code        int         `json:"code"`
	Message     string      `json:"message"`
	Description string      `json:"description,omitempty"`
	Data        interface{} `json:"data,omitempty"`
}

func Success(c *fiber.Ctx, data interface{}, message string) error {
	response := Response{
		Status:  "Berhasil",
		Code:    fiber.StatusOK,
		Message: message,
		Data:    data,
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

func Error(c *fiber.Ctx, statusCode int, message string, desc string) error {
	response := Response{
		Status:      "Gagal",
		Code:        statusCode,
		Description: desc,
		Message:     message,
		// Data:    nil,
	}

	return c.Status(statusCode).JSON(response)
}
