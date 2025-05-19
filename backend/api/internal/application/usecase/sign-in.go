package usecase

import (
	"context"

	"github.com/GabrielChaves1/course/internal/application/dto/request"
	"github.com/GabrielChaves1/course/internal/clients/idp"
)

type SignInUseCase struct {
	idpClient idp.Client
}

func NewSignInUseCase(idpClient idp.Client) *SignInUseCase {
	return &SignInUseCase{
		idpClient: idpClient,
	}
}

func (uc *SignInUseCase) Execute(ctx context.Context, cmd request.SignInDTO) error {
	return nil
}
