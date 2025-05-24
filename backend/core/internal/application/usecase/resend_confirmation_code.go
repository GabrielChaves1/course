package usecase

import (
	"context"

	"github.com/GabrielChaves1/course/internal/application/dto/request"
	"github.com/GabrielChaves1/course/internal/clients/idp"
	apptypes "github.com/GabrielChaves1/course/internal/domain/types"
	apperrors "github.com/GabrielChaves1/course/internal/errors"
	"github.com/GabrielChaves1/course/pkg/validator"
)

type ResendConfirmationCode struct {
	idpClient idp.Client
}

func NewResendConfirmationCode(idpClient idp.Client) *ResendConfirmationCode {
	return &ResendConfirmationCode{
		idpClient: idpClient,
	}
}

func (uc *ResendConfirmationCode) Execute(ctx context.Context, cmd request.ResendConfirmationCodeDTO) error {
	validation := validator.NewValidator()

	email, err := apptypes.NewEmail(cmd.Email)
	if err != nil {
		validation.Add(validator.FieldEmail, err.Error())
	}

	if validation.HasErrors() {
		return apperrors.NewValidationError(ctx, validation)
	}

	if err := uc.idpClient.ResendConfirmationCode(ctx, email.Value); err != nil {
		return apperrors.NewSystemError(ctx, "There was a problem confirming account")
	}

	return nil
}
