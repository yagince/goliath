package golidation

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
	result := ValidationResult{make(Errors)}

	for name, validators := range validation.validators {
		message, ok := validators.IsSatisfied(values[name])
		if !ok {
			result.AddError(Error{Key: name, Message: message})
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

func (field *FieldValidator) Required() *FieldValidator {
	field.validators = append(field.validators, Required{})
	return field
}

func (field *FieldValidator) MinLength(length int) *FieldValidator {
	field.validators = append(field.validators, MinLength{length})
	return field
}

func (field *FieldValidator) MaxLength(length int) *FieldValidator {
	field.validators = append(field.validators, MaxLength{length})
	return field
}

type ValidationResult struct {
	errors Errors
}

func (result ValidationResult) AddError(error Error) {
	result.errors[error.Key] = error
}

func (result ValidationResult) HasError() bool {
	return len(result.errors) > 0
}

func (result ValidationResult) Errors() Errors {
	return result.errors
}

type Error struct {
	Key     string
	Message string
}

type Errors map[string]Error
