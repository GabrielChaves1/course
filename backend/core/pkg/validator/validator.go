package validator

type ValidatorField string

const (
	FieldEmail            ValidatorField = "email"
	FieldPassword         ValidatorField = "password"
	FieldConfirmationCode ValidatorField = "code"
	FieldPhoneNumber      ValidatorField = "phone_number"
)

type ValidationError struct {
	Field   ValidatorField `json:"field"`
	Message string         `json:"message"`
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

func (v *Validator) HasErrors() bool {
	return len(v.Errors) > 0
}
