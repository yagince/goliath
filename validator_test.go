package goliath

import (
	. "./test_util"
	"testing"
)

func TestCustomMessageValidator(t *testing.T) {
	c := CustomMessageValidator{Required{}, "custom message"}
	Assert(t, Equal{c.Message(), "custom message"})
	Assert(t, IsFalse{c.IsSatisfied("")})
}
