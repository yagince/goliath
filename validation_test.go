package goliath

import (
	. "./test_util"
	"testing"
)

func TestValidation_OK(t *testing.T) {
	validation := NewValidation()
	validation.Field("name").Required().MinLength(2).MaxLength(3)

	result := validation.Validate(map[string]interface{}{"name": "aa"})
	Assert(t, Equal{false, result.HasError()})
}

func TestValidation_EmptyMap(t *testing.T) {
	validation := NewValidation()
	validation.Field("name").Required().MinLength(2).MaxLength(3)

	var (
		result ValidationResult
		errors ValidationErrors
		error  ValidationError
	)

	result = validation.Validate(map[string]interface{}{})
	Verify(t, IsTrue{result.HasError()})
	if !result.HasError() {
		t.Error("should has some errors")
	}

	errors = result.Errors()
	error, ok := errors["name"]
	Assert(t, IsTrue{ok})
	Verify(t, Equal{error.Message, (Required{}).Message()})

}

func TestValidation_MinLength(t *testing.T) {
	validation := NewValidation()
	validation.Field("name").Required().MinLength(2).MaxLength(3)

	var (
		result ValidationResult
		errors ValidationErrors
		error  ValidationError
	)

	result = validation.Validate(map[string]interface{}{"name": "1"})
	Verify(t, IsTrue{result.HasError()})

	errors = result.Errors()
	error, ok := errors["name"]
	Assert(t, IsTrue{ok})
	Verify(t, Equal{error.Message, (MinLength{Length: 2}).Message()})

}

func TestValidation_MaxLength(t *testing.T) {
	validation := NewValidation()
	validation.Field("name").Required().MinLength(2).MaxLength(3)

	var (
		result ValidationResult
		errors ValidationErrors
		error  ValidationError
	)

	result = validation.Validate(map[string]interface{}{"name": "1234"})
	Verify(t, IsTrue{result.HasError()})

	errors = result.Errors()
	error, ok := errors["name"]
	Assert(t, IsTrue{ok})
	Verify(t, Equal{error.Message, (MaxLength{Length: 3}).Message()})

}

func TestValidation_Pattern(t *testing.T) {
	validation := NewValidation()
	validation.Field("name").Pattern("^\\w$")

	result := validation.Validate(map[string]interface{}{"name": "aaa hoge"})
	Verify(t, IsTrue{result.HasError()})

	errors := result.Errors()
	_, ok := errors["name"]
	Assert(t, IsTrue{ok})
}

func TestValidation_MultiValue(t *testing.T) {
	validation := NewValidation()
	validation.Field("name").Required().MinLength(2).MaxLength(3)
	validation.Field("password").MinLength(6).MaxLength(10)

	var (
		result ValidationResult
		errors ValidationErrors
	)

	result = validation.Validate(map[string]interface{}{"name": "1234", "password": "12345"})
	Verify(t, IsTrue{result.HasError()})

	errors = result.Errors()
	{
		error, ok := errors["name"]
		Assert(t, IsTrue{ok})
		Verify(t, Equal{error.Message, (MaxLength{Length: 3}).Message()})
	}
	{
		error, ok := errors["password"]
		Assert(t, IsTrue{ok})
		Verify(t, Equal{error.Message, (MinLength{Length: 6}).Message()})
	}
}

type CustomValidator struct{}

func (v CustomValidator) IsSatisfied(value interface{}) bool {
	_, ok := value.(string)
	return ok
}
func (v CustomValidator) Message() string {
	return "must be String"
}

func TestValidation_Custom(t *testing.T) {
	validation := NewValidation()
	validation.Field("name").AddValidator(CustomValidator{})

	result := validation.Validate(map[string]interface{}{"name": 100})
	Verify(t, IsTrue{result.HasError()})

	errors := result.Errors()
	error, ok := errors["name"]
	Assert(t, IsTrue{ok})
	Verify(t, Equal{error.Message, (CustomValidator{}).Message()})
}

func TestValidation_CustomMessage(t *testing.T) {
	validation := NewValidation()
	validation.Field("name").Required("req").MinLength(2, "min").MaxLength(4, "max")

	{
		errors := validation.Validate(map[string]interface{}{}).Errors()
		err, ok := errors["name"]
		Assert(t, IsTrue{ok})
		Verify(t, Equal{err.Message, "req"})
	}

	{
		errors := validation.Validate(map[string]interface{}{"name": "a"}).Errors()
		err, ok := errors["name"]
		Assert(t, IsTrue{ok})
		Verify(t, Equal{err.Message, "min"})
	}

	{
		errors := validation.Validate(map[string]interface{}{"name": "aaaaa"}).Errors()
		err, ok := errors["name"]
		Assert(t, IsTrue{ok})
		Verify(t, Equal{err.Message, "max"})
	}
}

func TestValidation_Each(t *testing.T) {
	validation := NewValidation()
	validation.Field("items").Required().MinLength(1).Each().Max(3)

	{
		params := map[string]interface{}{
			"items": 100,
		}

		result := validation.Validate(params)
		Verify(t, IsTrue{result.HasError()})

		errors := result.Errors()
		err, ok := errors["items"]
		Assert(t, IsTrue{ok})
		Verify(t, Equal{(MinLength{1}).Message(), err.Message})
	}

	{
		params := map[string]interface{}{
			"items": "1",
		}

		result := validation.Validate(params)
		Verify(t, IsTrue{result.HasError()})

		errors := result.Errors()
		err, ok := errors["items"]
		Assert(t, IsTrue{ok})
		Verify(t, Equal{(ArrayOrSlice{}).Message(), err.Message})
	}

	{
		params := map[string]interface{}{
			"items": []int{1, 4},
		}

		result := validation.Validate(params)
		Verify(t, IsTrue{result.HasError()})

		errors := result.Errors()
		err, ok := errors["items"]
		Assert(t, IsTrue{ok})
		Verify(t, Equal{(Max{3}).Message(), err.Message})
	}

	{
		params := map[string]interface{}{
			"items": []interface{}{1, 1},
		}

		result := validation.Validate(params)
		Verify(t, IsFalse{result.HasError()})

	}
}
