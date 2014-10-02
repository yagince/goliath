package goliath

import (
	"fmt"
)

// value >= Threshold
type MinInt struct {
	Threshold int
	message   string
}

func (validator MinInt) IsSatisfied(value interface{}) bool {
	if value == nil {
		return true
	}

	switch v := value.(type) {
	case int:
		return v >= validator.Threshold
	case int32:
		return v >= int32(validator.Threshold)
	case int64:
		return v >= int64(validator.Threshold)
	default:
		return false
	}
}

func (validator MinInt) Message() string {
	return ChoiceMessage(validator.message, fmt.Sprintf("must %d or more", validator.Threshold))
}
