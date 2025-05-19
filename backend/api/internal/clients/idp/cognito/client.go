package cognitoidp

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/GabrielChaves1/course/internal/clients/idp"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
)

type CognitoProvider struct {
	client     *cognitoidentityprovider.Client
	clientID   string
	userPoolId string
}

func NewCognitoProvider(sdk aws.Config, clientID, userPoolID string) (idp.Client, error) {
	client := cognitoidentityprovider.NewFromConfig(sdk)

	if client == nil {
		return nil, fmt.Errorf("fail to create cognito client")
	}

	return &CognitoProvider{
		client:     client,
		clientID:   clientID,
		userPoolId: userPoolID,
	}, nil
}

func (c *CognitoProvider) Authenticate(ctx context.Context, username, password string) (*idp.AuthenticateResult, error) {
	var authResult *types.AuthenticationResultType

	input := &cognitoidentityprovider.InitiateAuthInput{
		AuthFlow: "USER_PASSWORD_AUTH",
		ClientId: aws.String(c.clientID),
		AuthParameters: map[string]string{
			"USERNAME": username,
			"PASSWORD": password,
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

	output, err := c.client.SignUp(ctx, &cognitoidentityprovider.SignUpInput{
		ClientId: aws.String(c.clientID),
		Password: aws.String(password),
		Username: aws.String(username),
		UserAttributes: []types.AttributeType{
			{Name: aws.String("email"), Value: aws.String(email)},
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
