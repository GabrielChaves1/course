package idp

import "context"

type AuthenticateResult struct {
	AccessToken  string
	IDToken      string
	RefreshToken string
	ExpiresIn    int32
}

type Client interface {
	Authenticate(ctx context.Context, username, password string) (*AuthenticateResult, error)
	VerifyEmail(ctx context.Context, username, confirmationCode string) error
	SignUp(ctx context.Context, username, email, password string) (bool, error)
	ForgotPassword(ctx context.Context, username string) error
	ConfirmForgotPassword(ctx context.Context, username, confirmationCode, newPassword string) error
}
