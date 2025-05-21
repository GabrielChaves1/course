package types

import (
	"errors"
	"regexp"
	"strings"
)

const (
	emailRegex = `^[a-z0-9.]+@[a-z0-9]+\.[a-z]+\.([a-z]+)?$/i`
)

type Email struct {
	Value string
}

func (e Email) Validate() error {
	if strings.TrimSpace(e.Value) == "" {
		return errors.New("Email should not be empty")
	}

	re, _ := regexp.Compile(emailRegex)

	if !re.MatchString(e.Value) {
		return errors.New("Invalid email")
	}

	return nil
}

func NewEmail(email string) (*Email, error) {
	e := &Email{Value: email}
	if err := e.Validate(); err != nil {
		return nil, err
	}

	return e, nil
}
