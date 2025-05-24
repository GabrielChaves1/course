package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/GabrielChaves1/course/internal/application/dto/request"
	"github.com/GabrielChaves1/course/internal/clients/idp"
	apptypes "github.com/GabrielChaves1/course/internal/domain/types"
	apperrors "github.com/GabrielChaves1/course/internal/errors"
	"github.com/GabrielChaves1/course/pkg/validator"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
)

type SignUp struct {
	idpClient idp.Client
}

func NewSignUp(idpClient idp.Client) *SignUp {
	return &SignUp{
		idpClient: idpClient,
	}
}

func (uc *SignUp) Execute(ctx context.Context, cmd request.SignUpDTO) error {
	validation := validator.NewValidator()

	email, err := apptypes.NewEmail(cmd.Email)
	if err != nil {
		validation.Add(validator.FieldEmail, err.Error())
	}

	password, err := apptypes.NewPassword(cmd.Password)
	if err != nil {
		validation.Add(validator.FieldPassword, err.Error())
	}

	phoneNumber, err := apptypes.NewPhoneNumber(cmd.PhoneNumber)
	if err != nil {
		validation.Add(validator.FieldPhoneNumber, err.Error())
	}

	if cmd.Password != cmd.ConfirmPassword {
		validation.Add(validator.FieldPassword, "password and confirm password must be equal")
	}

	if validation.HasErrors() {
		return apperrors.NewValidationError(ctx, validation)
	}

	if _, err := uc.idpClient.SignUp(ctx, email.Value, email.Value, phoneNumber.Value, password.Value); err != nil {
		var alreadyExists *types.UsernameExistsException
		if errors.As(err, &alreadyExists) {
			return apperrors.NewAuthenticationError(ctx, fmt.Sprintf("User %s already exists", email.Value))
		}

		return apperrors.NewSystemError(ctx, "There was a problem registering the user")
	}

	return nil
}
