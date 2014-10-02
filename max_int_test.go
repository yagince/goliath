package goliath

import (
	"fmt"
	"testing"
)

func TestMaxInt_IsSatisfied_int(t *testing.T) {
	var validator Validator = MaxInt{Threshold: 5}
	Verify(t, IsTrue{validator.IsSatisfied(5)})
	Verify(t, IsFalse{validator.IsSatisfied(6)})

	Verify(t, IsTrue{validator.IsSatisfied(int32(5))})
	Verify(t, IsFalse{validator.IsSatisfied(int32(6))})

	Verify(t, IsTrue{validator.IsSatisfied(int64(5))})
	Verify(t, IsFalse{validator.IsSatisfied(int64(6))})
}

func TestMaxInt_IsSatisfied_Other(t *testing.T) {
	var validator Validator = MaxInt{Threshold: 5}
	Verify(t, IsFalse{validator.IsSatisfied(5.0)})
	Verify(t, IsFalse{validator.IsSatisfied(4.1)})

	Verify(t, IsFalse{validator.IsSatisfied("")})
	Verify(t, IsFalse{validator.IsSatisfied("12345")})

	Verify(t, IsFalse{validator.IsSatisfied([]int{1, 2, 3})})
	Verify(t, IsFalse{validator.IsSatisfied(map[int]int{1: 1})})
}

func TestMaxInt_Message(t *testing.T) {
	test := func(threshold int) {
		Verify(t, Equal{
			(MaxInt{Threshold: threshold}).Message(),
			fmt.Sprintf("must %d or less", threshold),
		})
	}

	test(1)
	test(2)
	test(100)
}
