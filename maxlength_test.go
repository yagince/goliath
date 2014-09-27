package golidation

import (
	"fmt"
	"testing"
)

func TestMaxLength_IsSatisfied_String(t *testing.T) {
	var validator Validator = MaxLength{2}

	if !validator.IsSatisfied("") {
		t.Error("empty string is 0 length value")
	}

	if !validator.IsSatisfied("aa") {
		t.Error("aa is 1 length")
	}

	if validator.IsSatisfied("aaa") {
		t.Error("aaa is over 2 length")
	}
}

func TestMaxLength_IsSatisfied_Slice(t *testing.T) {
	var validator Validator = MaxLength{2}
	value := []string{"a", "b", "c"}

	if validator.IsSatisfied(value) {
		t.Errorf("%v is over 2 length", value)
	}
}

func TestMaxLength_IsSatisfied_Map(t *testing.T) {
	var validator Validator = MaxLength{2}
	value := map[int]int{1: 1, 2: 2, 3: 3}

	if validator.IsSatisfied(value) {
		t.Errorf("%v is over 2 length", value)
	}
}

func TestMaxLength_IsSatisfied_Others(t *testing.T) {
	var validator Validator = MaxLength{1}

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
		if (MaxLength{length}).Message() != message {
			t.Errorf("message must be %s", message)
		}
	}

	test(1)
	test(2)
	test(100)
}
