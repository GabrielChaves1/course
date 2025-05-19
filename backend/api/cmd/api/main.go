package main

import (
	"context"
	"net/http"
	"time"

	"github.com/GabrielChaves1/course/internal/http/handlers/router"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
)

func initializeDependencies(cfg *Config) (*cognitoidentityprovider.Client, error) {
	ctx := context.Background()
	sdk, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, err
	}

	cognitoClient := cognitoidentityprovider.NewFromConfig(sdk)

	return cognitoClient, nil
}

func main() {
	config, err := NewConfig()
	if err != nil {
		panic(err)
	}

	_, err = initializeDependencies(config)
	if err != nil {
		panic(err)
	}

	routerConfig := router.APIRouterConfig{
		Environment: config.environment,
	}

	ginRouter := router.SetupAPIRouter(routerConfig)

	server := &http.Server{
		Addr:           ":8080",
		Handler:        ginRouter,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
