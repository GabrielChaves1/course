package types

import (
	"errors"
	"fmt"
	"strings"
)

type ConfirmationCode struct {
	Value string
}

const (
	length = 6
)

func (c ConfirmationCode) Validate() error {
	if strings.TrimSpace(c.Value) == "" {
		return errors.New("Confirmation code cannot be empty")
	}

	if len(c.Value) < length {
		return fmt.Errorf("Confirmation code must be at least %d digits", length)
	}

	return nil
}

func NewConfirmationCode(code string) (*ConfirmationCode, error) {
	e := &ConfirmationCode{Value: code}
	if err := e.Validate(); err != nil {
		return nil, err
	}

	return e, nil
}
