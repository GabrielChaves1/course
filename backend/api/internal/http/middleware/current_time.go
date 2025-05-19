package middleware

import (
	"time"

	appcontext "github.com/GabrielChaves1/course/internal/application/context"
	"github.com/gin-gonic/gin"
)

func CurrentTime() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := appcontext.NewContextWithCurrentTime(c.Request.Context(), time.Now().UTC())
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
