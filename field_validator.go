package goliath

type FieldValidator struct {
	Name       string
	validators []Validator
}

func NewFieldValidator(name string) *FieldValidator {
	return &FieldValidator{name, make([]Validator, 0)}
}

func (field FieldValidator) IsSatisfied(value interface{}) (message string, ok bool) {
	for _, validator := range field.validators {
		ok = validator.IsSatisfied(value)
		if !ok {
			message = validator.Message()
			return
		}
	}
	return "", true
}

func (field *FieldValidator) AddValidator(v Validator) *FieldValidator {
	field.validators = append(field.validators, v)
	return field
}

func extractMessage(m ...string) string {
	if len(m) > 0 {
		return m[0]
	}
	return ""
}

func (field *FieldValidator) Required(customMessage ...string) *FieldValidator {
	validator := Required{}
	validator.message = extractMessage(customMessage...)
	return field.AddValidator(validator)
}

func (field *FieldValidator) MinLength(length int, customMessage ...string) *FieldValidator {
	validator := MinLength{Length: length}
	validator.message = extractMessage(customMessage...)
	return field.AddValidator(validator)
}

func (field *FieldValidator) MaxLength(length int, customMessage ...string) *FieldValidator {
	validator := MaxLength{Length: length}
	validator.message = extractMessage(customMessage...)
	return field.AddValidator(validator)
}

func (field *FieldValidator) MinInt(threshold int, customMessage ...string) *FieldValidator {
	validator := MinInt{Threshold: threshold}
	validator.message = extractMessage(customMessage...)
	return field.AddValidator(validator)
}

func (field *FieldValidator) MaxInt(threshold int, customMessage ...string) *FieldValidator {
	validator := MaxInt{Threshold: threshold}
	validator.message = extractMessage(customMessage...)
	return field.AddValidator(validator)
}
