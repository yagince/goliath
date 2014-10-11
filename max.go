package goliath

import (
	"fmt"
)

// value <= Threshold
type Max struct {
	Threshold float64
}

func (validator Max) IsSatisfied(value interface{}) bool {
	if value == nil {
		return true
	}

	switch v := value.(type) {
	case int:
		return float64(v) <= validator.Threshold
	case int8:
		return float64(v) <= validator.Threshold
	case int16:
		return float64(v) <= validator.Threshold
	case int32:
		return float64(v) <= validator.Threshold
	case int64:
		return float64(v) <= validator.Threshold
	case float32:
		return float64(v) <= validator.Threshold
	case float64:
		return v <= validator.Threshold
	default:
		return false
	}
}

func (validator Max) Message() string {
	return fmt.Sprintf("must %d or less", validator.Threshold)
}
