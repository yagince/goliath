package goliath

import (
	"fmt"
	"reflect"
)

// (length of value) >= Length
type MinLength struct {
	Length int
}

func (validator MinLength) IsSatisfied(value interface{}) bool {
	if value == nil {
		return true
	}

	if str, ok := value.(string); ok {
		return len(str) >= validator.Length
	}

	reflectValue := reflect.ValueOf(value)
	if reflectValue.Kind() == reflect.Slice || reflectValue.Kind() == reflect.Map {
		return reflectValue.Len() >= validator.Length
	}

	return false
}

func (validator MinLength) Message() string {
	return fmt.Sprintf("must have %d or more elements", validator.Length)
}
