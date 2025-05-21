package cognitoidp

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"log"

	"github.com/GabrielChaves1/course/internal/clients/idp"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
)

type CognitoProvider struct {
	client       *cognitoidentityprovider.Client
	clientID     string
	clientSecret string
	userPoolId   string
}

func NewCognitoProvider(sdk aws.Config, clientID, userPoolID, clientSecret string) (idp.Client, error) {
	client := cognitoidentityprovider.NewFromConfig(sdk)

	if client == nil {
		return nil, fmt.Errorf("fail to create cognito client")
	}

	return &CognitoProvider{
		client:       client,
		clientID:     clientID,
		userPoolId:   userPoolID,
		clientSecret: clientSecret,
	}, nil
}

func (c *CognitoProvider) Authenticate(ctx context.Context, username, password string) (*idp.AuthenticateResult, error) {
	var authResult *types.AuthenticationResultType
	secretHash := calculateSecretHash(username, c.clientID, c.clientSecret)

	input := &cognitoidentityprovider.InitiateAuthInput{
		AuthFlow: "USER_PASSWORD_AUTH",
		ClientId: aws.String(c.clientID),
		AuthParameters: map[string]string{
			"USERNAME":    username,
			"PASSWORD":    password,
			"SECRET_HASH": secretHash,
		},
	}

	output, err := c.client.InitiateAuth(ctx, input)

	if err != nil {
		var resetRequired *types.PasswordResetRequiredException
		if errors.As(err, &resetRequired) {
			log.Println(*resetRequired.Message)
		} else {
			log.Printf("Couldn't sign in user %v. Here's why: %v\n", username, err)
		}
	} else {
		authResult = output.AuthenticationResult
	}

	return &idp.AuthenticateResult{
		AccessToken:  *authResult.AccessToken,
		RefreshToken: *authResult.RefreshToken,
		IDToken:      *authResult.IdToken,
		ExpiresIn:    authResult.ExpiresIn,
	}, nil
}

func (c *CognitoProvider) SignUp(ctx context.Context, username, email, password string) (bool, error) {
	confirmed := false
	secretHash := calculateSecretHash(username, c.clientID, c.clientSecret)

	output, err := c.client.SignUp(ctx, &cognitoidentityprovider.SignUpInput{
		ClientId:   aws.String(c.clientID),
		Password:   aws.String(password),
		Username:   aws.String(username),
		SecretHash: aws.String(secretHash),
		UserAttributes: []types.AttributeType{
			{
				Name:  aws.String("email"),
				Value: aws.String(email),
			},
		},
	})

	if err != nil {
		var invalidPassword *types.InvalidPasswordException
		if errors.As(err, &invalidPassword) {
			log.Println(*invalidPassword.Message)
		} else {
			log.Printf("Couldn't sign up user %v. Here's why: %v\n", email, err)
		}
	} else {
		confirmed = output.UserConfirmed
	}

	return confirmed, err
}

func (c *CognitoProvider) VerifyEmail(ctx context.Context, username, confirmationCode string) error {
	secretHash := calculateSecretHash(username, c.clientID, c.clientSecret)

	params := &cognitoidentityprovider.ConfirmSignUpInput{
		ClientId:         aws.String(c.clientID),
		Username:         aws.String(username),
		ConfirmationCode: aws.String(confirmationCode),
		SecretHash:       aws.String(secretHash),
	}

	if _, err := c.client.ConfirmSignUp(ctx, params); err != nil {
		return err
	}

	return nil
}

func (c *CognitoProvider) ForgotPassword(ctx context.Context, username string) error {
	secretHash := calculateSecretHash(username, c.clientID, c.clientSecret)

	params := &cognitoidentityprovider.ForgotPasswordInput{
		ClientId:   aws.String(c.clientID),
		Username:   aws.String(username),
		SecretHash: aws.String(secretHash),
	}

	if _, err := c.client.ForgotPassword(ctx, params); err != nil {
		return err
	}

	return nil
}

func (c *CognitoProvider) ConfirmForgotPassword(ctx context.Context, username, confirmationCode, newPassword string) error {
	secretHash := calculateSecretHash(username, c.clientID, c.clientSecret)

	params := &cognitoidentityprovider.ConfirmForgotPasswordInput{
		ClientId:         aws.String(c.clientID),
		Username:         aws.String(username),
		Password:         aws.String(newPassword),
		ConfirmationCode: aws.String(confirmationCode),
		SecretHash:       aws.String(secretHash),
	}

	if _, err := c.client.ConfirmForgotPassword(ctx, params); err != nil {
		return err
	}

	return nil
}

func calculateSecretHash(username, clientID, clientSecret string) string {
	mac := hmac.New(sha256.New, []byte(clientSecret))
	mac.Write([]byte(username + clientID))
	return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}
