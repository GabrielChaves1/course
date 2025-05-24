package usecase

import (
	"context"

	"github.com/GabrielChaves1/course/internal/application/dto/request"
	"github.com/GabrielChaves1/course/internal/application/dto/response"
	"github.com/GabrielChaves1/course/internal/clients/idp"
	apptypes "github.com/GabrielChaves1/course/internal/domain/types"
	apperrors "github.com/GabrielChaves1/course/internal/errors"
	"github.com/GabrielChaves1/course/pkg/validator"
)

type SignIn struct {
	idpClient idp.Client
}

func NewSignIn(idpClient idp.Client) *SignIn {
	return &SignIn{
		idpClient: idpClient,
	}
}

func (uc *SignIn) Execute(ctx context.Context, cmd request.SignInDTO) (*response.SignInDTO, error) {
	validation := validator.NewValidator()

	email, err := apptypes.NewEmail(cmd.Email)
	if err != nil {
		validation.Add(validator.FieldEmail, err.Error())
	}

	password, err := apptypes.NewPassword(cmd.Password)
	if err != nil {
		validation.Add(validator.FieldPassword, err.Error())
	}

	if validation.HasErrors() {
		return nil, apperrors.NewValidationError(ctx, validation)
	}

	result, err := uc.idpClient.Authenticate(ctx, email.Value, password.Value)
	if err != nil {
		return nil, apperrors.NewSystemError(ctx, "System error")
	}

	return &response.SignInDTO{
		AccessToken:  result.AccessToken,
		RefreshToken: result.RefreshToken,
		IDToken:      result.IDToken,
		ExpiresIn:    result.ExpiresIn,
	}, nil
}
