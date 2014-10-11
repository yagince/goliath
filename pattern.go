package goliath

import (
	"fmt"
	"regexp"
)

type Pattern struct {
	pattern *regexp.Regexp
}

func (validator Pattern) IsSatisfied(value interface{}) bool {
	if value == nil {
		return false
	}

	if str, ok := value.(string); ok {
		return validator.pattern.MatchString(str)
	} else {
		return false
	}
}

func (validator Pattern) Message() string {
	return fmt.Sprintf("must match pattern `%s`", validator.pattern.String())
}
