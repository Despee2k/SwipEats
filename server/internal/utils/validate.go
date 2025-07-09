package utils

import (
	"github.com/go-playground/validator/v10"
)

var Validate = validator.New()

func GetValidationErrorDetails(err error) (string, map[string]string) {
	if err == nil {
		return "", nil
	}

	var errorMessage string

	validationErrors := err.(validator.ValidationErrors)
	details := make(map[string]string)

	for i, fieldError := range validationErrors {
		fieldName := fieldError.Field()
		details[fieldName] = GetFieldMessage(fieldName, fieldError.Tag())

		if i == 0 {
			errorMessage = details[fieldName]
		}
	}

	return errorMessage, details
}

func GetFieldMessage(fieldName string, tag string) string {
	capFieldName := Capitalize(fieldName)

	switch tag {
	case "required":
		return capFieldName + " is required"
	case "email":
		return capFieldName + " must be a valid email address"
	case "min":
		return capFieldName + " must be at least 8 characters long"
	case "max":
		return capFieldName + " must not exceed the maximum length"
	default:
		return capFieldName + " is invalid"
	}
}