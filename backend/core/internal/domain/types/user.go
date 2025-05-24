package types

import (
	"database/sql/driver"

	"github.com/google/uuid"
)

type UserID uuid.UUID

func NewUserIDFromString(rawId string) (UserID, error) {
	id, err := uuid.Parse(rawId)
	if err != nil {
		return UserID{}, err
	}

	return UserID(id), nil
}

func (id UserID) Value() (driver.Value, error) {
	return uuid.UUID(id).Value()
}
