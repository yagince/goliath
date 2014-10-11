package goliath

import (
	. "./test_util"
	"fmt"
	"testing"
)

func TestMin_IsSatisfied_int(t *testing.T) {
	var validator Validator = Min{Threshold: 5}
	Verify(t, IsTrue{validator.IsSatisfied(5)})
	Verify(t, IsTrue{validator.IsSatisfied(5.1)})
	Verify(t, IsFalse{validator.IsSatisfied(4)})
	Verify(t, IsFalse{validator.IsSatisfied(4.1)})

	Verify(t, IsTrue{validator.IsSatisfied(int32(5))})
	Verify(t, IsFalse{validator.IsSatisfied(int32(4))})

	Verify(t, IsTrue{validator.IsSatisfied(int64(5))})
	Verify(t, IsFalse{validator.IsSatisfied(int64(4))})
}

func TestMin_IsSatisfied_Other(t *testing.T) {
	var validator Validator = Min{Threshold: 5}

	Verify(t, IsFalse{validator.IsSatisfied("")})
	Verify(t, IsFalse{validator.IsSatisfied("12345")})

	Verify(t, IsFalse{validator.IsSatisfied([]int{1, 2, 3})})
	Verify(t, IsFalse{validator.IsSatisfied(map[int]int{1: 1})})
}

func TestMin_Message(t *testing.T) {
	test := func(threshold float64) {
		Verify(t, Equal{
			(Min{Threshold: threshold}).Message(),
			fmt.Sprintf("must %d or more", threshold),
		})
	}

	test(1)
	test(2)
	test(100)
}
