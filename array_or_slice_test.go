package goliath

import (
	. "./test_util"
	"testing"
)

func TestArrayOrSlice_IsSatisfied_Array(t *testing.T) {
	var validator Validator = ArrayOrSlice{}
	arry := [1]int{1}
	Verify(t, IsTrue{validator.IsSatisfied(arry)})
}

func TestArrayOrSlice_IsSatisfied_Slice(t *testing.T) {
	var validator Validator = ArrayOrSlice{}
	slice := []int{1}
	Verify(t, IsTrue{validator.IsSatisfied(slice)})
}

func TestArrayOrSlice_IsSatisfied_Nil(t *testing.T) {
	var validator Validator = ArrayOrSlice{}
	Verify(t, IsTrue{validator.IsSatisfied(nil)})
}

func TestArrayOrSlice_IsSatisfied_String(t *testing.T) {
	var validator Validator = ArrayOrSlice{}
	Verify(t, IsFalse{validator.IsSatisfied("")})
}

func TestArrayOrSlice_IsSatisfied_Number(t *testing.T) {
	var validator Validator = ArrayOrSlice{}
	Verify(t, IsFalse{validator.IsSatisfied(0)})
}

func TestArrayOrSlice_Message(t *testing.T) {
	var validator Validator = ArrayOrSlice{}
	Verify(t, Equal{validator.Message(), "must be array"})
}
