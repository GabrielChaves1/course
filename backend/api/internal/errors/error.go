package errors

import (
	"context"
	"time"

	"github.com/GabrielChaves1/course/internal/errors/validator"
)

const (
	ValidationError    string = "VALIDATION_ERROR"
	NotFoundError      string = "RESOURCE_NOT_FOUND"
	AuthorizationError string = "UNAUTHORIZED_ACCESS"
	SystemError        string = "SYSTEM_ERROR"
)

type AppError struct {
	Code      string
	Detail    string
	Context   map[string]interface{}
	Timestamp time.Time
}

func (e AppError) Error() string {
	return e.Detail
}

func NewValidationError(ctx context.Context, validation *validator.Validator) AppError {
	return AppError{
		Code:   ValidationError,
		Detail: "Validation error",
		Context: map[string]interface{}{
			"errors": validation.Errors,
		},
		Timestamp: time.Now(),
	}
}
