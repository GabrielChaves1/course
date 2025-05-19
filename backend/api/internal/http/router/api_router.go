package router

import (
	"net/http"

	"github.com/GabrielChaves1/course/internal/domain/types"
	"github.com/GabrielChaves1/course/internal/http/handlers"
	"github.com/GabrielChaves1/course/internal/http/middleware"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
)

type APIRouterConfig struct {
	Environment types.Environment
}

func SetupAPIRouter(config APIRouterConfig, authHandlers *handlers.AuthenticationHandlers) *gin.Engine {
	if config.Environment == types.Production {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	router := gin.New()

	router.Use(middleware.CurrentTime())
	router.Use(requestid.New())
	router.Use(gin.Recovery())

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	auth := router.Group("/auth")
	{
		auth.POST("/sign-in", authHandlers.SignIn)
		auth.POST("/sign-up", authHandlers.SignUp)
	}

	return router
}
