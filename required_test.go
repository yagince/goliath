package goliath

import (
	"testing"
)

func TestRequired_IsSatisfied_Nil(t *testing.T) {
	var validator Validator = Required{}

	if validator.IsSatisfied(nil) {
		t.Error("nil is undefined value")
	}
}

func TestRequired_IsSatisfied_String(t *testing.T) {
	var validator Validator = Required{}

	if validator.IsSatisfied("") {
		t.Error("empty string is undefined value")
	}

	if !validator.IsSatisfied("aaa") {
		t.Error("aaa is defined value")
	}
}

func TestRequired_IsSatisfied_Number(t *testing.T) {
	var validator Validator = Required{}

	if !validator.IsSatisfied(0) {
		t.Error("0is defined value")
	}
}

func TestRequired_Message(t *testing.T) {
	var validator Validator = Required{}
	message := "required"

	if validator.Message() != message {
		t.Errorf("message of required is %s", message)
	}
}
