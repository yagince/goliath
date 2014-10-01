package goliath

import ()

type Validation struct {
	validators map[string]*FieldValidator
}

func NewValidation() *Validation {
	return &Validation{make(map[string]*FieldValidator)}
}

func (validation *Validation) Field(name string) *FieldValidator {
	fieldValidator, ok := validation.validators[name]
	if !ok {
		fieldValidator = NewFieldValidator(name)
		validation.validators[name] = fieldValidator
	}
	return fieldValidator
}

func (validation *Validation) Validate(values map[string]interface{}) ValidationResult {
	result := ValidationResult{make(ValidationErrors)}

	for name, validators := range validation.validators {
		message, ok := validators.IsSatisfied(values[name])
		if !ok {
			result.AddError(ValidationError{Key: name, Message: message})
		}
	}
	return result
}

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

func (field *FieldValidator) Required() *FieldValidator {
	return field.AddValidator(Required{})
}

func (field *FieldValidator) MinLength(length int) *FieldValidator {
	return field.AddValidator(MinLength{length})
}

func (field *FieldValidator) MaxLength(length int) *FieldValidator {
	return field.AddValidator(MaxLength{length})
}

type ValidationResult struct {
	errors ValidationErrors
}

func (result ValidationResult) AddError(error ValidationError) {
	result.errors[error.Key] = error
}

func (result ValidationResult) HasError() bool {
	return len(result.errors) > 0
}

func (result ValidationResult) Errors() ValidationErrors {
	return result.errors
}

type ValidationError struct {
	Key     string
	Message string
}

type ValidationErrors map[string]ValidationError
