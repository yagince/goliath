package goliath

import (
	. "./test_util"
	"regexp"
	"testing"
)

func compileReg(pattern string) *regexp.Regexp {
	return regexp.MustCompile(pattern)
}

func TestPattern_IsSatisfied_Nil(t *testing.T) {
	var validator Validator = Pattern{compileReg(".*")}

	if validator.IsSatisfied(nil) {
		t.Error("nil does not match any patterns")
	}
}

func TestPattern_IsSatisfied_Number(t *testing.T) {
	var validator Validator = Pattern{compileReg(".*")}

	Verify(t, IsFalse{validator.IsSatisfied(0)})
	Verify(t, IsFalse{validator.IsSatisfied(1)})
	Verify(t, IsFalse{validator.IsSatisfied(1.1)})
	Verify(t, IsFalse{validator.IsSatisfied(0.11)})
}

func TestPattern_IsSatisfied_AnyStringValue(t *testing.T) {
	var validator Validator = Pattern{compileReg(".*")}
	Verify(t, IsTrue{validator.IsSatisfied("")})
}

func TestPattern_IsSatisfied_MatchFoo(t *testing.T) {
	var validator Validator = Pattern{compileReg("foo")}
	Verify(t, IsFalse{validator.IsSatisfied("")})
	Verify(t, IsTrue{validator.IsSatisfied("foo")})
	Verify(t, IsTrue{validator.IsSatisfied("hogefoo")})
}

func TestPattern_IsSatisfied_MatchFullFoo(t *testing.T) {
	var validator Validator = Pattern{compileReg("^foo$")}
	Verify(t, IsFalse{validator.IsSatisfied("")})
	Verify(t, IsTrue{validator.IsSatisfied("foo")})
	Verify(t, IsFalse{validator.IsSatisfied("hogefoo")})
}

func TestPattern_Message(t *testing.T) {
	var validator Validator = Pattern{compileReg("^[0-9]+$")}
	Assert(t, Equal{"must match pattern `^[0-9]+$`", validator.Message()})
}
