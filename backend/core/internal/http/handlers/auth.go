package handlers

import (
	"net/http"

	"github.com/GabrielChaves1/course/internal/application/dto/request"
	"github.com/GabrielChaves1/course/internal/application/usecase"
	"github.com/gin-gonic/gin"
)

type AuthenticationHandlers struct {
	signIn                 *usecase.SignIn
	signUp                 *usecase.SignUp
	confirmSignUp          *usecase.ConfirmSignUp
	resendConfirmationCode *usecase.ResendConfirmationCode
	forgotPassword         *usecase.ForgotPassword
	resetPassword          *usecase.ResetPassword
	logout                 *usecase.Logout
}

func NewAuthenticationHandlers(
	signIn *usecase.SignIn,
	signUp *usecase.SignUp,
	confirmSignUp *usecase.ConfirmSignUp,
	resendConfirmationCode *usecase.ResendConfirmationCode,
	forgotPassword *usecase.ForgotPassword,
	resetPassword *usecase.ResetPassword,
	logout *usecase.Logout,
) *AuthenticationHandlers {
	return &AuthenticationHandlers{
		signIn:                 signIn,
		signUp:                 signUp,
		confirmSignUp:          confirmSignUp,
		resendConfirmationCode: resendConfirmationCode,
		forgotPassword:         forgotPassword,
		resetPassword:          resetPassword,
		logout:                 logout,
	}
}

func (h AuthenticationHandlers) SignIn(c *gin.Context) {
	var dto request.SignInDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.Error(err)
		return
	}

	result, err := h.signIn.Execute(c.Request.Context(), dto)
	if err != nil {
		c.Error(err)
		return
	}

	c.SetCookie("access_token", result.AccessToken, int(result.ExpiresIn), "/", "", false, true)
	c.SetCookie("id_token", result.IDToken, int(result.ExpiresIn), "/", "", false, true)
	c.SetCookie("refresh_token", result.RefreshToken, 30*24*60*60, "/refresh", "", false, true)

	c.JSON(http.StatusAccepted, gin.H{
		"message": "User signed in successfully",
	})
}

func (h AuthenticationHandlers) SignUp(c *gin.Context) {
	var dto request.SignUpDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.Error(err)
		return
	}

	err := h.signUp.Execute(c.Request.Context(), dto)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"message": "User created successfully, check your email to activate your account",
	})
}

func (h AuthenticationHandlers) ConfirmSignUp(c *gin.Context) {
	var dto request.ConfirmSignUpDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.Error(err)
		return
	}

	err := h.confirmSignUp.Execute(c.Request.Context(), dto)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"message": "User email verified successfully",
	})
}

func (h AuthenticationHandlers) ResendConfirmationCode(c *gin.Context) {
	var dto request.ResendConfirmationCodeDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.Error(err)
		return
	}

	err := h.resendConfirmationCode.Execute(c.Request.Context(), dto)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"message": "Confirmation code resent successfully",
	})
}

func (h AuthenticationHandlers) ForgotPassword(c *gin.Context) {
	var dto request.ForgotPasswordDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.Error(err)
		return
	}

	err := h.forgotPassword.Execute(c.Request.Context(), dto)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Confirmation code sent successfully",
	})
}

func (h AuthenticationHandlers) ResetPassword(c *gin.Context) {
	var dto request.ResetPasswordDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.Error(err)
		return
	}

	err := h.resetPassword.Execute(c.Request.Context(), dto)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Password changed successfully",
	})
}

func (h AuthenticationHandlers) Logout(c *gin.Context) {
	accessToken, _ := c.Cookie("access_token")

	err := h.logout.Execute(c.Request.Context(), accessToken)
	if err != nil {
		c.Error(err)
		return
	}
}
