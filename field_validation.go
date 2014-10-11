package goliath

import (
	. "reflect"
)

type FieldValidation interface {
	IsSatisfied(value interface{}) (message string, ok bool)
	AddValidator(v Validator) FieldValidation
	Required(customMessage ...string) FieldValidation
	MinLength(length int, customMessage ...string) FieldValidation
	MaxLength(length int, customMessage ...string) FieldValidation
	Min(threshold float64, customMessage ...string) FieldValidation
	Max(threshold float64, customMessage ...string) FieldValidation
	Each() FieldValidation
}

type BasicFieldValidation struct {
	Name       string
	validators []Validator
	next       FieldValidation
}

func NewBasicFieldValidation(name string) *BasicFieldValidation {
	return &BasicFieldValidation{
		Name:       name,
		validators: make([]Validator, 0),
		next:       nil,
	}
}

func (field *BasicFieldValidation) IsSatisfied(value interface{}) (message string, ok bool) {
	if message, ok = field.execValidators(value); !ok {
		return
	}

	if field.next != nil {
		return field.next.IsSatisfied(value)
	}

	return
}

func (field *BasicFieldValidation) execValidators(value interface{}) (message string, ok bool) {
	for _, validator := range field.validators {
		ok = validator.IsSatisfied(value)
		if !ok {
			message = validator.Message()
			return
		}
	}
	return "", true
}

func (field *BasicFieldValidation) AddValidator(v Validator) FieldValidation {
	field.validators = append(field.validators, v)
	return field
}

func (field *BasicFieldValidation) Required(customMessage ...string) FieldValidation {
	return field.AddValidator(wrapCustomMessageValidator(Required{}, customMessage...))
}

func (field *BasicFieldValidation) MinLength(length int, customMessage ...string) FieldValidation {
	return field.AddValidator(wrapCustomMessageValidator(MinLength{length}, customMessage...))
}

func (field *BasicFieldValidation) MaxLength(length int, customMessage ...string) FieldValidation {
	return field.AddValidator(wrapCustomMessageValidator(MaxLength{length}, customMessage...))
}

func (field *BasicFieldValidation) Min(threshold float64, customMessage ...string) FieldValidation {
	return field.AddValidator(wrapCustomMessageValidator(Min{threshold}, customMessage...))
}

func (field *BasicFieldValidation) Max(threshold float64, customMessage ...string) FieldValidation {
	return field.AddValidator(wrapCustomMessageValidator(Max{threshold}, customMessage...))
}

func (field *BasicFieldValidation) Each() FieldValidation {
	nested := &ArrayFieldValidation{NewBasicFieldValidation(field.Name)}
	field.next = nested
	return nested
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

type ArrayFieldValidation struct {
	*BasicFieldValidation
}

func (field ArrayFieldValidation) IsSatisfied(value interface{}) (message string, ok bool) {
	arrayOrSlice := ArrayOrSlice{}
	if !arrayOrSlice.IsSatisfied(value) {
		return arrayOrSlice.Message(), false
	}

	v := ValueOf(value)
	for i := 0; i < v.Len(); i++ {
		if message, ok = field.execValidators(v.Index(i).Interface()); !ok {
			return
		}
	}

	if field.next != nil {
		return field.next.IsSatisfied(value)
	}

	return "", true
}
