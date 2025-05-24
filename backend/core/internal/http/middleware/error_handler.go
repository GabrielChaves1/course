package middleware

import (
	"errors"
	"net/http"
	"time"

	apperrors "github.com/GabrielChaves1/course/internal/errors"
	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Code      string                 `json:"code"`
	Detail    string                 `json:"detail"`
	Context   map[string]interface{} `json:"context,omitempty"`
	Timestamp time.Time              `json:"timestamp"`
}

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err

			var appError *apperrors.AppError
			if ok := errors.As(err, &appError); ok {
				resp := ErrorResponse{
					Code:      appError.Code,
					Detail:    appError.Detail,
					Context:   appError.Context,
					Timestamp: appError.Timestamp,
				}

				c.JSON(httpStatusFromCode(appError.Code), resp)
				c.Abort()
				return
			}

			c.JSON(http.StatusInternalServerError, ErrorResponse{
				Code:      apperrors.SystemError,
				Detail:    "An internal server error has occurred",
				Timestamp: time.Now(),
			})
		}
	}
}

func httpStatusFromCode(code string) int {
	switch code {
	case apperrors.AuthenticationError:
		return http.StatusUnauthorized
	case apperrors.ValidationError:
		return http.StatusBadRequest
	case apperrors.NotFoundError:
		return http.StatusNotFound
	case apperrors.SystemError:
		return http.StatusInternalServerError
	case apperrors.AuthorizationError:
		return http.StatusForbidden
	default:
		return http.StatusInternalServerError
	}
}
