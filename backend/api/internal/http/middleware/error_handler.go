package middleware

import (
	"github.com/GabrielChaves1/course/internal/errors"
	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err

			if appErr, ok := err.(*errors.AppError); ok {
			
			}
		}
	}
}
