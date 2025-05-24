package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/GabrielChaves1/course/internal/application/usecase"
	"github.com/GabrielChaves1/course/internal/clients/idp"
	cognitoidp "github.com/GabrielChaves1/course/internal/clients/idp/cognito"
	"github.com/GabrielChaves1/course/internal/http/handlers"
	"github.com/GabrielChaves1/course/internal/http/router"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

func initializeDependencies(cfg *Config) (idp.Client, error) {
	ctx := context.Background()
	sdk, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, err
	}

	ssmClient := ssm.NewFromConfig(sdk)

	output, err := ssmClient.GetParameter(ctx, &ssm.GetParameterInput{
		Name:           aws.String("/cognito/client_secret"),
		WithDecryption: aws.Bool(true),
	})
	if err != nil {
		return nil, err
	}

	cognitoClientSecret := *output.Parameter.Value

	cognitoClient, err := cognitoidp.NewCognitoProvider(sdk, cfg.cognito.appClientID, cfg.cognito.userPoolID, cognitoClientSecret)
	if err != nil {
		return nil, err
	}

	return cognitoClient, nil
}

func main() {
	config, err := NewConfig()
	if err != nil {
		panic(err)
	}

	cognitoClient, err := initializeDependencies(config)
	if err != nil {
		panic(err)
	}

	signInUseCase := usecase.NewSignIn(cognitoClient)
	signUpUseCase := usecase.NewSignUp(cognitoClient)
	confirmSignUpUseCase := usecase.NewConfirmSignUp(cognitoClient)
	resendConfirmationCodeUseCase := usecase.NewResendConfirmationCode(cognitoClient)
	forgotPasswordUseCase := usecase.NewForgotPassword(cognitoClient)
	resetPasswordUseCase := usecase.NewResetPassword(cognitoClient)
	logoutUseCase := usecase.NewLogout(cognitoClient)

	authHandlers := handlers.NewAuthenticationHandlers(
		signInUseCase,
		signUpUseCase,
		confirmSignUpUseCase,
		resendConfirmationCodeUseCase,
		forgotPasswordUseCase,
		resetPasswordUseCase,
		logoutUseCase,
	)

	routerConfig := router.APIRouterConfig{
		Environment: config.environment,
	}

	ginRouter := router.SetupAPIRouter(routerConfig, authHandlers)

	server := &http.Server{
		Addr:           ":8080",
		Handler:        ginRouter,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		fmt.Println("Failed to shutdown HTTP server")
	} else {
		fmt.Println("HTTP server shutdown")
	}
}
