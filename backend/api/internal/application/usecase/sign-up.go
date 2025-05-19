package usecase

import (
	"context"

	"github.com/GabrielChaves1/course/internal/application/dto/request"
	"github.com/GabrielChaves1/course/internal/clients/idp"
)

type SignUpUseCase struct {
	idpClient idp.Client
}

func NewSignUpUseCase(idpClient idp.Client) *SignUpUseCase {
	return &SignUpUseCase{
		idpClient: idpClient,
	}
}

func (uc *SignUpUseCase) Execute(ctx context.Context, cmd request.SignUpDTO) error {
	if ok, err := uc.idpClient.SignUp(ctx, cmd.Email, cmd.Email, cmd.Password); !ok {
		return err
	}

	return nil
}
