package usecase

import (
	"context"

	"github.com/GabrielChaves1/course/internal/clients/idp"
	apperrors "github.com/GabrielChaves1/course/internal/errors"
)

type Logout struct {
	idpClient idp.Client
}

func NewLogout(idpClient idp.Client) *Logout {
	return &Logout{
		idpClient: idpClient,
	}
}

func (uc *Logout) Execute(ctx context.Context, accessToken string) error {
	if err := uc.idpClient.Logout(ctx, accessToken); err != nil {
		return apperrors.NewSystemError(ctx, "Error while logging out")
	}

	return nil
}
