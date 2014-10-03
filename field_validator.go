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

func (field *FieldValidator) Required(customMessage ...string) *FieldValidator {
	return field.AddValidator(wrapCustomMessageValidator(Required{}, customMessage...))
}

func (field *FieldValidator) MinLength(length int, customMessage ...string) *FieldValidator {
	return field.AddValidator(wrapCustomMessageValidator(MinLength{length}, customMessage...))
}

func (field *FieldValidator) MaxLength(length int, customMessage ...string) *FieldValidator {
	return field.AddValidator(wrapCustomMessageValidator(MaxLength{length}, customMessage...))
}

func (field *FieldValidator) MinInt(threshold int, customMessage ...string) *FieldValidator {
	return field.AddValidator(wrapCustomMessageValidator(MinInt{threshold}, customMessage...))
}

func (field *FieldValidator) MaxInt(threshold int, customMessage ...string) *FieldValidator {
	return field.AddValidator(wrapCustomMessageValidator(MaxInt{threshold}, customMessage...))
}

func extractMessage(m ...string) (message string, ok bool) {
	if len(m) > 0 {
		return m[0], true
	}
	return "", false
}

func wrapCustomMessageValidator(validator Validator, message ...string) Validator {
	if m, ok := extractMessage(message...); ok {
		return CustomMessageValidator{validator, m}
	} else {
		return validator
	}
}
