package usecase

import (
	"context"

	"github.com/GabrielChaves1/course/internal/application/dto/request"
	"github.com/GabrielChaves1/course/internal/clients/idp"
	apptypes "github.com/GabrielChaves1/course/internal/domain/types"
	apperrors "github.com/GabrielChaves1/course/internal/errors"
	"github.com/GabrielChaves1/course/pkg/validator"
)

type ForgotPassword struct {
	idpClient idp.Client
}

func NewForgotPassword(idpClient idp.Client) *ForgotPassword {
	return &ForgotPassword{
		idpClient: idpClient,
	}
}

func (uc *ForgotPassword) Execute(ctx context.Context, cmd request.ForgotPasswordDTO) error {
	validation := validator.NewValidator()

	email, err := apptypes.NewEmail(cmd.Email)
	if err != nil {
		validation.Add(validator.FieldEmail, err.Error())
	}

	if validation.HasErrors() {
		return apperrors.NewValidationError(ctx, validation)
	}

	err = uc.idpClient.ForgotPassword(ctx, email.Value)
	if err != nil {
		return apperrors.NewSystemError(ctx, "System error")
	}

	return nil
}
