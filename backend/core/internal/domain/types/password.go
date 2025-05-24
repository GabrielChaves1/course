package types

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

type Password struct {
	Value string
}

const (
	minLength = 8
	maxLength = 100
)

func (p Password) Validate() error {
	if strings.TrimSpace(p.Value) == "" {
		return errors.New("Password should not be empty")
	}
	if len(p.Value) < minLength {
		return fmt.Errorf("Password must be at least %d characters", minLength)
	}

	if len(p.Value) > maxLength {
		return fmt.Errorf("Password cannot exceed %d characters", maxLength)
	}

	if match, _ := regexp.MatchString("[0-9]", p.Value); !match {
		return errors.New("Password must be at least one number")
	}

	if match, _ := regexp.MatchString("[A-Z]", p.Value); !match {
		return errors.New("Password must be at least one uppercase letter")
	}

	if match, _ := regexp.MatchString("[a-z]", p.Value); !match {
		return errors.New("Password must be at least one lowercase letter")
	}

	return nil
}

func NewPassword(password string) (*Password, error) {
	p := &Password{Value: password}
	if err := p.Validate(); err != nil {
		return nil, err
	}

	return p, nil
}
