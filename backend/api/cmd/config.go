package main

import (
	"fmt"
	"log"
	"os"

	"github.com/GabrielChaves1/course/internal/domain/types"
	"github.com/joho/godotenv"
)

type Cognito struct {
	appClientID string
	userPoolID  string
}

type Config struct {
	cognito     Cognito
	environment types.Environment
}

func NewConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

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

	cognitoAppClientID := os.Getenv("COGNITO_APP_CLIENT_ID")
	if cognitoAppClientID == "" {
		return nil, fmt.Errorf("env var COGNITO_APP_CLIENT_ID not defined")
	}

	cognitoUserPoolID := os.Getenv("COGNITO_USER_POOL_ID")
	if cognitoUserPoolID == "" {
		return nil, fmt.Errorf("env var COGNITO_USER_POOL_ID not defined")
	}

	return &Config{
		environment: environment,
		cognito: Cognito{
			// issuer:          cognitoIssuer,
			appClientID: cognitoAppClientID,
			userPoolID:  cognitoUserPoolID,
		},
	}, nil
}
