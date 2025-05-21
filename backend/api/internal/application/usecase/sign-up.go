package usecase

import (
	"context"

	"github.com/GabrielChaves1/course/internal/application/dto/request"
	"github.com/GabrielChaves1/course/internal/clients/idp"
	"github.com/GabrielChaves1/course/internal/domain/types"
	"github.com/GabrielChaves1/course/internal/errors"
	"github.com/GabrielChaves1/course/internal/errors/validator"
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
	validation := validator.NewValidator()

	email, err := types.NewEmail(cmd.Email)
	if err != nil {
		validation.Add(validator.FieldEmail, err.Error())
	}

	password, err := types.NewPassword(cmd.Password)
	if err != nil {
		validation.Add(validator.FieldPassword, err.Error())
	}

	if validation.HasErrors() {
		return errors.NewValidationError(ctx, validation)
	}

	if ok, err := uc.idpClient.SignUp(ctx, email.Value, email.Value, password.Value); !ok {
		return err
	}

	return nil
}
