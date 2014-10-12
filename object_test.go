package goliath

import (
	. "./test_util"
	"testing"
)

func TestJsonObject_IsSatisfied_Map_KeyString(t *testing.T) {
	var validator Validator = JsonObject{}
	m := map[string]interface{}{}
	Verify(t, IsTrue{validator.IsSatisfied(m)})
}

func TestJsonObject_IsSatisfied_Map_KeyInt(t *testing.T) {
	var validator Validator = JsonObject{}
	m := map[int]interface{}{}
	Verify(t, IsFalse{validator.IsSatisfied(m)})
}

func TestJsonObject_IsSatisfied_Nil(t *testing.T) {
	var validator Validator = JsonObject{}
	Verify(t, IsTrue{validator.IsSatisfied(nil)})
}

func TestJsonObject_Cast(t *testing.T) {
	validator := JsonObject{}
	var value interface{} = map[string]interface{}{}
	var casted map[string]interface{} = validator.cast(value)
	Verify(t, Equal{value, casted})
}
