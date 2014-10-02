package goliath

import (
	"fmt"
	"testing"
)

func TestMinLength_IsSatisfied_String(t *testing.T) {
	var validator Validator = MinLength{Length: 1}

	if validator.IsSatisfied("") {
		t.Error("empty string is 0 length value")
	}

	if !validator.IsSatisfied("a") {
		t.Error("a is 1 length")
	}

	if !validator.IsSatisfied("aaa") {
		t.Error("aaa is over 1 length")
	}
}

func TestMinLength_IsSatisfied_Slice(t *testing.T) {
	var validator Validator = MinLength{Length: 1}

	if validator.IsSatisfied([]string{}) {
		t.Error("empty slice is 0 length value")
	}
}

func TestMinLength_IsSatisfied_Map(t *testing.T) {
	var validator Validator = MinLength{Length: 1}
	value := map[int]int{}

	if validator.IsSatisfied(value) {
		t.Errorf("%v is 0 length", value)
	}
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
