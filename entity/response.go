package entity

import (
	"github.com/gofiber/fiber/v2"
)

func Success(c *fiber.Ctx, data interface{}, message string) error {
	response := Response{
		Status:  fiber.StatusOK,
		Message: message,
		Data:    data,
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

func Error(c *fiber.Ctx, statusCode int, message string) error {
	response := Response{
		Status:  statusCode,
		Message: message,
		// Data:    nil,
	}

	return c.Status(statusCode).JSON(response)
}