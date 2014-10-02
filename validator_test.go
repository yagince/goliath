package goliath

import (
	"testing"
)

func TestChoiceMessage(t *testing.T) {
	m := ChoiceMessage("", "hoge", "foo")
	Assert(t, Equal{m, "hoge"})
}
