package goliath

import (
	"reflect"
)

type ArrayOrSlice struct{}

func (validator ArrayOrSlice) IsSatisfied(value interface{}) bool {
	if value == nil {
		return true
	}

	switch reflect.TypeOf(value).Kind() {
	case reflect.Slice, reflect.Array:
		return true
	default:
		return false
	}
}

func (validator ArrayOrSlice) Message() string {
	return "must be array"
}
