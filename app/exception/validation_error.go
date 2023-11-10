package exception

import (
	"github.com/go-playground/validator/v10"
)

type (
	ValidationError struct {
		FieldErrors []validator.FieldError
		Old         interface{}
	}
)

func NewValidationError(fieldErrors []validator.FieldError, old interface{}) *ValidationError {
	return &ValidationError{
		FieldErrors: fieldErrors,
		Old:         old,
	}
}

func (e *ValidationError) Error() string {
	return "Validation Error!"
}

func (e *ValidationError) GetFieldErrors() []string {
	var errors []string

	for _, fieldError := range e.FieldErrors {
		errors = append(errors, formatError(fieldError))
	}

	return errors
}

func formatError(fieldError validator.FieldError) string {
	switch fieldError.Tag() {
	case "required":
		return fieldError.Field() + " is required!"
	case "email":
		return fieldError.Field() + " must be a valid email address!"
	case "min":
		return fieldError.Field() + " must be at least " + fieldError.Param() + " characters!"
	default:
		return fieldError.Field() + " is not valid!"
	}
}
