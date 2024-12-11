package middleware

import (
	"reflect"

	"github.com/abiyyu03/siruta/config"
	"github.com/abiyyu03/siruta/usecase/helper"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// ErrorResponse formats the error response
type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
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

// func rtNumberUniqueValidation(fl validator.FieldLevel) bool {
// 	var rtProfile *model.RTProfile
//     fieldValue := fl.Field().Interface()

// 	if err := config.DB.Where("rt_number_id = ? AND rw_email = ?").First(&rtProfile).Error; err != nil {
// 		return false
// 	}
// 	return true
// }

// ValidateStruct is a generic function to validate struct data
func ValidateStruct(data interface{}) []*ErrorResponse {
	validate := validator.New()
	var errors []*ErrorResponse

	// Register the unique validation function
	validate.RegisterValidation("unique", uniqueValidation)
	// validate.RegisterValidation("rt_number_unique", rtNumberUniqueValidation)

	// Run the validations
	err := validate.Struct(data)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, &ErrorResponse{
				FailedField: err.StructNamespace(),
				Tag:         err.Tag(),
				// Value:       err.Err,
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
