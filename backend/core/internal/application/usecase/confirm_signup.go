package usecase

import (
	"context"

	"github.com/GabrielChaves1/course/internal/application/dto/request"
	"github.com/GabrielChaves1/course/internal/clients/idp"
	apptypes "github.com/GabrielChaves1/course/internal/domain/types"
	apperrors "github.com/GabrielChaves1/course/internal/errors"
	"github.com/GabrielChaves1/course/pkg/validator"
)

type ConfirmSignUp struct {
	idpClient idp.Client
}

func NewConfirmSignUp(idpClient idp.Client) *ConfirmSignUp {
	return &ConfirmSignUp{
		idpClient: idpClient,
	}
}

func (uc *ConfirmSignUp) Execute(ctx context.Context, cmd request.ConfirmSignUpDTO) error {
	validation := validator.NewValidator()

	email, err := apptypes.NewEmail(cmd.Email)
	if err != nil {
		validation.Add(validator.FieldEmail, err.Error())
	}

	code, err := apptypes.NewConfirmationCode(cmd.Code)
	if err != nil {
		validation.Add(validator.FieldConfirmationCode, err.Error())
	}

	if validation.HasErrors() {
		return apperrors.NewValidationError(ctx, validation)
	}

	if err := uc.idpClient.VerifyEmail(ctx, email.Value, code.Value); err != nil {
		return apperrors.NewSystemError(ctx, "There was a problem confirming account")
	}

	return nil
}
