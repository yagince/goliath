package goliath

import (
	"fmt"
)

// value <= Threshold
type MaxInt struct {
	Threshold int
}

func (validator MaxInt) IsSatisfied(value interface{}) bool {
	if value == nil {
		return true
	}

	switch v := value.(type) {
	case int:
		return v <= validator.Threshold
	case int32:
		return v <= int32(validator.Threshold)
	case int64:
		return v <= int64(validator.Threshold)
	default:
		return false
	}
}

func (validator MaxInt) Message() string {
	return fmt.Sprintf("must %d or less", validator.Threshold)
}
