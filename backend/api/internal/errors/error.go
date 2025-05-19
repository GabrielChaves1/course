package errors

import (
	"context"
	"time"
)

const (
	ValidationError string = "VALIDATION_ERROR"
	NotFoundError string = "RESOURCE_NOT_FOUND"
	AuthorizationError string = "UNAUTHORIZED_ACCESS"
	SystemError string = "SYSTEM_ERROR"
)

type AppError struct {
	Code string
	Detail string
	Context map[string]interface{}
	Timestamp time.Time
}

func (e AppError) Error() string {
	return e.Detail
}

func NewValidationAppError(ctx context.Context, validationErrs *error)
