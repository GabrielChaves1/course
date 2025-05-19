package main

import (
	"fmt"
	"os"

	"github.com/GabrielChaves1/course/internal/domain/types"
)

type Cognito struct {
	issuer          string
	appClientID     string
	appClientSecret string
	userPoolID      string
}

type Config struct {
	cognito     Cognito
	environment types.Environment
}

func NewConfig() (*Config, error) {
	env := os.Getenv("ENVIRONMENT")
	var environment types.Environment

	switch env {
	case "production":
		environment = types.Production
	case "staging":
		environment = types.Staging
	default:
		environment = types.Development
	}

	cognitoIssuer := os.Getenv("COGNITO_ISSUER")
	if cognitoIssuer == "" {
		return nil, fmt.Errorf("env var COGNITO_ISSUER not defined")
	}

	cognitoAppClientID := os.Getenv("COGNITO_APP_CLIENT_ID")
	if cognitoAppClientID == "" {
		return nil, fmt.Errorf("env var COGNITO_APP_CLIENT_ID not defined")
	}

	cognitoAppClientSecret := os.Getenv("COGNITO_APP_CLIENT_SECRET")
	if cognitoAppClientSecret == "" {
		return nil, fmt.Errorf("env var COGNITO_APP_CLIENT_SECRET not defined")
	}

	cognitoUserPoolID := os.Getenv("COGNITO_USER_POOL_ID")
	if cognitoUserPoolID == "" {
		return nil, fmt.Errorf("env var COGNITO_USER_POOL_ID not defined")
	}

	return &Config{
		environment: environment,
		cognito: Cognito{
			issuer:          cognitoIssuer,
			appClientID:     cognitoAppClientID,
			appClientSecret: cognitoAppClientSecret,
			userPoolID:      cognitoUserPoolID,
		},
	}, nil
}
