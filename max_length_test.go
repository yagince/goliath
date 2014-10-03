package goliath

import (
	. "./test_util"
	"fmt"
	"testing"
)

func TestMaxLength_IsSatisfied_String(t *testing.T) {
	var validator Validator = MaxLength{Length: 2}

	Verify(t, IsTrue{validator.IsSatisfied("")})
	Verify(t, IsTrue{validator.IsSatisfied("aa")})
	Verify(t, IsFalse{validator.IsSatisfied("aaa")})
}

func TestMaLength_IsSatisfied_Slice(t *testing.T) {
	var validator Validator = MaxLength{Length: 2}

	Verify(t, IsFalse{validator.IsSatisfied([]string{"a", "b", "c"})})
}

func TestMaxLength_IsSatisfied_Map(t *testing.T) {
	var validator Validator = MaxLength{Length: 2}

	Verify(t, IsFalse{validator.IsSatisfied(map[int]int{1: 1, 2: 2, 3: 3})})
}

func TestMaxLength_IsSatisfied_Others(t *testing.T) {
	var validator Validator = MaxLength{Length: 1}

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

func TestMaxLength_Message(t *testing.T) {
	test := func(length int) {
		message := fmt.Sprintf("must have %d or less elements", length)
		if (MaxLength{Length: length}).Message() != message {
			t.Errorf("message must be %s", message)
		}
	}

	test(1)
	test(2)
	test(100)
}
