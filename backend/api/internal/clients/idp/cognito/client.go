package cognitoidp

import (
	"context"

	"github.com/GabrielChaves1/course/internal/clients/idp"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
)

type CognitoProvider struct {
	client *cognitoidentityprovider.Client
}

func NewCognitoProvider(client *cognitoidentityprovider.Client) idp.Client {
	return &CognitoProvider{
		client: client,
	}
}

func Authenticate(ctx context.Context, username, password string) error {
	return nil
}
