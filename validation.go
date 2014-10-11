package goliath

import ()

type Validation struct {
	fieldValidations map[string]FieldValidation
}

func NewValidation() *Validation {
	return &Validation{make(map[string]FieldValidation)}
}

func (validation *Validation) Field(name string) FieldValidation {
	fieldValidator, ok := validation.fieldValidations[name]
	if !ok {
		fieldValidator = NewBasicFieldValidation(name)
		validation.fieldValidations[name] = fieldValidator
	}
	return fieldValidator
}

func (validation *Validation) Validate(values map[string]interface{}) ValidationResult {
	result := ValidationResult{make(ValidationErrors)}

	for name, fieldValidation := range validation.fieldValidations {
		message, ok := fieldValidation.IsSatisfied(values[name])
		if !ok {
			result.AddError(ValidationError{Key: name, Message: message})
		}
	}
	return result
}

type asycValidtaionResult struct {
	name    string
	message string
	ok      bool
}

func (validation *Validation) ValidateParallel(values map[string]interface{}) ValidationResult {
	result := ValidationResult{make(ValidationErrors)}

	validationCount := len(validation.fieldValidations)
	c := make(chan asycValidtaionResult, validationCount)

	for name, fieldValidation := range validation.fieldValidations {
		go func(n string, fV FieldValidation, v interface{}) {
			message, ok := fieldValidation.IsSatisfied(values[name])
			c <- asycValidtaionResult{name, message, ok}
		}(name, fieldValidation, values[name])
	}

	for i := 0; i < validationCount; i++ {
		r := <-c
		if !r.ok {
			result.AddError(ValidationError{Key: r.name, Message: r.message})
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
