package domainerrors

type ValidationRule string

const (
	Required     ValidationRule = "required"
	MinLength    ValidationRule = "min_length"
	MaxLength    ValidationRule = "max_length"
	InvalidValue ValidationRule = "invalid_value"
)

type ValidationError struct {
	Field string
	Value interface{}
	Rule  ValidationRule
}

func NewValidationError(field string, value interface{}, rule ValidationRule) ValidationError {
	return ValidationError{
		Field: field,
		Value: value,
		Rule:  rule,
	}
}

type ValidationErrors []ValidationError

func (ves *ValidationErrors) Add(field string, value interface{}, rule ValidationRule) {
	*ves = append(*ves, NewValidationError(field, value, rule))
}

func (ves *ValidationErrors) HasErrors() bool {
	return len(*ves) > 0
}
