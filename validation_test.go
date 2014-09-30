package goliath

import (
	"testing"
)

func TestValidation_EmptyMap(t *testing.T) {
	validation := NewValidation()
	validation.Field("name").Required().MinLength(2).MaxLength(3)

	var (
		result ValidationResult
		errors Errors
		error  Error
	)

	result = validation.Validate(map[string]interface{}{})
	if !result.HasError() {
		t.Error("should has some errors")
	}

	errors = result.Errors()
	error, ok := errors["name"]
	if !ok {
		t.Error("should has error of `name`")
		t.FailNow()
	}

	expect := (Required{}).Message()
	if error.Message != expect {
		t.Error("should be require error")
		t.Logf("expect %s", expect)
		t.Logf("got    %s", error.Message)
	}

}

func TestValidation_MinLength(t *testing.T) {
	validation := NewValidation()
	validation.Field("name").Required().MinLength(2).MaxLength(3)

	var (
		result ValidationResult
		errors Errors
		error  Error
	)

	result = validation.Validate(map[string]interface{}{"name": "1"})
	if !result.HasError() {
		t.Error("should has some errors")
	}

	errors = result.Errors()
	error, ok := errors["name"]
	if !ok {
		t.Error("should has error of `name`")
		t.FailNow()
	}

	expect := (MinLength{2}).Message()
	if error.Message != expect {
		t.Error("should be minlength error")
		t.Logf("expect %s", expect)
		t.Logf("got    %s", error.Message)
	}

}

func TestValidation_MaxLength(t *testing.T) {
	validation := NewValidation()
	validation.Field("name").Required().MinLength(2).MaxLength(3)

	var (
		result ValidationResult
		errors Errors
		error  Error
	)

	result = validation.Validate(map[string]interface{}{"name": "1234"})
	if !result.HasError() {
		t.Error("should has some errors")
	}

	errors = result.Errors()
	error, ok := errors["name"]
	if !ok {
		t.Error("should has error of `name`")
		t.FailNow()
	}

	expect := (MaxLength{3}).Message()
	if error.Message != expect {
		t.Error("should be maxlength error")
		t.Logf("expect %s", expect)
		t.Logf("got    %s", error.Message)
	}

}
