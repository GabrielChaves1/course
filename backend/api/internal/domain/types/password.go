package types

import (
	"errors"
	"fmt"
	"strings"
)

type Password struct {
	Value string
}

const (
	minLength = 8
	maxLength = 100
)

func (e Password) Validate() error {
	if strings.TrimSpace(e.Value) == "" {
		return errors.New("Password should not be empty")
	}

	if len(e.Value) < minLength {
		return fmt.Errorf("Password must be at least %d characters", minLength)
	}

	if len(e.Value) > maxLength {
		return fmt.Errorf("Password cannot exceed %d characters", maxLength)
	}

	return nil
}

func NewPassword(password string) (*Password, error) {
	e := &Password{Value: password}
	if err := e.Validate(); err != nil {
		return nil, err
	}

	return e, nil
}
