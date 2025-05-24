package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPassword_Validate(t *testing.T) {
	tests := []struct {
		name     string
		password string
		wantErr  bool
	}{
		{"empty password", "", true},
		{"too short", "Abc123", true},
		{"no number", "Abcdefgh", true},
		{"no uppercase", "abcdefg1", true},
		{"no lowercase", "ABCDEFG1", true},
		{"valid password", "Abcdefg1", false},
		{"too long", string(make([]byte, 101)), true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Password{Value: tt.password}
			err := p.Validate()
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
