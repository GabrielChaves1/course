package validator

type ValidatorField string

const (
	FieldEmail    ValidatorField = "email"
	FieldPassword ValidatorField = "password"
)

type ValidationError struct {
	Field   ValidatorField
	Message string
}

type Validator struct {
	Errors []*ValidationError
}

func NewValidator() *Validator {
	return &Validator{
		Errors: make([]*ValidationError, 0),
	}
}

func (v *Validator) Add(field ValidatorField, message string) {
	v.Errors = append(v.Errors, &ValidationError{
		Field:   field,
		Message: message,
	})
}

func (v *Validator) Error() string {
	return ""
}

func (v *Validator) HasErrors() bool {
	return len(v.Errors) > 0
}
