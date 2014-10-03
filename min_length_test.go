package goliath

import (
	. "./test_util"
	"fmt"
	"testing"
)

func TestMinLength_IsSatisfied_String(t *testing.T) {
	var validator Validator = MinLength{Length: 1}

	Verify(t, IsFalse{validator.IsSatisfied("")})
	Verify(t, IsTrue{validator.IsSatisfied("a")})
	Verify(t, IsTrue{validator.IsSatisfied("aaa")})
}

func TestMinLength_IsSatisfied_Slice(t *testing.T) {
	var validator Validator = MinLength{Length: 1}

	Verify(t, IsFalse{validator.IsSatisfied([]string{})})
}

func TestMinLength_IsSatisfied_Map(t *testing.T) {
	var validator Validator = MinLength{Length: 1}

	Verify(t, IsFalse{validator.IsSatisfied(map[int]int{})})
}

func TestMinLength_IsSatisfied_Others(t *testing.T) {
	var validator Validator = MinLength{Length: 1}

	if !validator.IsSatisfied(nil) {
		t.Error("nil is ignore")
	}

	if validator.IsSatisfied(1) {
		t.Error("number has no length")
	}

	if validator.IsSatisfied(true) {
		t.Error("bool has no length")
	}

	type Sample struct{}

	if validator.IsSatisfied(Sample{}) {
		t.Error("struct has no length")
	}
}

func TestMinLength_Message(t *testing.T) {
	test := func(length int) {
		message := fmt.Sprintf("must have %d or more elements", length)
		if (MinLength{Length: length}).Message() != message {
			t.Errorf("message must be %s", message)
		}
	}

	test(1)
	test(2)
	test(100)
}
