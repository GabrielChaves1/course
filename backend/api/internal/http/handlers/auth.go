package handlers

import (
	"net/http"

	"github.com/GabrielChaves1/course/internal/application/dto/request"
	"github.com/GabrielChaves1/course/internal/application/usecase"
	"github.com/gin-gonic/gin"
)

type AuthenticationHandlers struct {
	signInUseCase *usecase.SignInUseCase
	signUpUseCase *usecase.SignUpUseCase
}

func NewAuthenticationHandlers(signInUseCase *usecase.SignInUseCase, signUpUseCase *usecase.SignUpUseCase) *AuthenticationHandlers {
	return &AuthenticationHandlers{
		signInUseCase: signInUseCase,
		signUpUseCase: signUpUseCase,
	}
}

func (h AuthenticationHandlers) SignIn(c *gin.Context) {
	var dto request.SignInDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		// middleware.HandleError(c, err)
		return
	}

	err := h.signInUseCase.Execute(c.Request.Context(), dto)
	if err != nil {
		// middleware.HandleError(c, err)
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"message": "User logged successfully",
	})
}

func (h AuthenticationHandlers) SignUp(c *gin.Context) {
	var dto request.SignUpDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		// middleware.HandleError(c, err)
		return
	}

	err := h.signUpUseCase.Execute(c.Request.Context(), dto)
	if err != nil {
		// middleware.HandleError(c, err)
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"message": "User logged successfully",
	})
}
