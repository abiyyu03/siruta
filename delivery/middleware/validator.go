package middleware

import (
	"reflect"

	"github.com/abiyyu03/siruta/config"
	"github.com/abiyyu03/siruta/helper"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// ErrorResponse formats the error response
type ErrorResponse struct {
	Field          string `json:"field"`
	ValidationType string `json:"validation_type"`
	Message        string `json:"message"`
}

// uniqueValidation function that works dynamically across different models
func uniqueValidation(fl validator.FieldLevel) bool {
	fieldName := helper.ConvertPascalCaseToSnakeCase(fl.StructFieldName()) // Name of the struct field
	fieldValue := fl.Field().Interface()                                   // Value of the struct field
	modelType := reflect.TypeOf(fl.Parent().Interface())                   // Get the struct type

	// Create an instance of the struct dynamically
	model := reflect.New(modelType).Interface()

	// Query the database to check for uniqueness
	var count int64
	if err := config.DB.Model(model).Where(fieldName+" = ?", fieldValue).Count(&count).Error; err != nil {
		return false
	}

	if count == 0 {
		return true
	}
	return false
}

// ValidateStruct is a generic function to validate struct data
func ValidateStruct(data interface{}) []*ErrorResponse {
	validate := validator.New()
	var errors []*ErrorResponse

	// Register the unique validation function
	validate.RegisterValidation("unique", uniqueValidation)
	// validate.RegisterValidation("rt_number_unique", RtNumberUniqueValidation)

	// Run the validations
	err := validate.Struct(data)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, &ErrorResponse{
				Field:          err.StructNamespace(),
				ValidationType: err.Tag(),
				Message:        err.Error(),
			})
		}
	}
	return errors
}

// ValidateField is a middleware function to validate fields
func ValidateField[T any]() func(c *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		var data T

		if err := ctx.BodyParser(&data); err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		// Validate the parsed data
		errors := ValidateStruct(data)
		if len(errors) > 0 {
			return ctx.Status(fiber.StatusBadRequest).JSON(errors)
		}

		return ctx.Next()
	}
}
