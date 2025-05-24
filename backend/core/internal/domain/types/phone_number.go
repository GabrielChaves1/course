package types

import (
	"errors"
	"regexp"
	"strings"
)

type PhoneNumber struct {
	Value string
}

var phoneRegex = regexp.MustCompile(`^\+55\d{10,11}$`)

func (p PhoneNumber) Validate() error {
	if strings.TrimSpace(p.Value) == "" {
		return errors.New("Phone number should not be empty")
	}

	if !phoneRegex.MatchString(p.Value) {
		return errors.New("Invalid phone number")
	}

	return nil
}

func NewPhoneNumber(phoneNumber string) (*PhoneNumber, error) {
	p := &PhoneNumber{Value: phoneNumber}
	if err := p.Validate(); err != nil {
		return nil, err
	}
	return p, nil
}
