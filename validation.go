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
