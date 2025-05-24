package idp

import (
	"context"
)

type AuthenticateResult struct {
	AccessToken  string
	IDToken      string
	RefreshToken string
	ExpiresIn    int32
}

type Client interface {
	Authenticate(ctx context.Context, username, password string) (*AuthenticateResult, error)
	VerifyEmail(ctx context.Context, username, confirmationCode string) error
	SignUp(ctx context.Context, username, email, phoneNumber, password string) (bool, error)
	ForgotPassword(ctx context.Context, username string) error
	ResetPassword(ctx context.Context, username, confirmationCode, newPassword string) error
	ResendConfirmationCode(ctx context.Context, username string) error
	Logout(ctx context.Context, token string) error
}
