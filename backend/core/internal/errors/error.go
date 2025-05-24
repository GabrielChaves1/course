package errors

import (
	"context"
	"time"

	"github.com/GabrielChaves1/course/pkg/validator"
)

const (
	ValidationError     string = "VALIDATION_ERROR"
	NotFoundError       string = "RESOURCE_NOT_FOUND"
	AuthorizationError  string = "UNAUTHORIZED_ACCESS"
	SystemError         string = "SYSTEM_ERROR"
	AuthenticationError string = "AUTHENTICATION_ERROR"
)

type AppError struct {
	Code      string
	Detail    string
	Context   map[string]any
	Timestamp time.Time
}

func (e AppError) Error() string {
	return e.Detail
}

func NewValidationError(ctx context.Context, validation *validator.Validator) *AppError {
	return &AppError{
		Code:   ValidationError,
		Detail: "Validation error",
		Context: map[string]any{
			"errors": validation.Errors,
		},
		Timestamp: time.Now(),
	}
}

func NewAuthenticationError(ctx context.Context, detail string) *AppError {
	return &AppError{
		Code:      AuthenticationError,
		Detail:    detail,
		Context:   map[string]any{},
		Timestamp: time.Now(),
	}
}

func NewSystemError(ctx context.Context, detail string) *AppError {
	return &AppError{
		Code:      SystemError,
		Detail:    detail,
		Context:   map[string]any{},
		Timestamp: time.Now(),
	}
}
