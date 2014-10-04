package goliath

import (
	"fmt"
	"runtime"
	"testing"
)

func BenchmarkSyncValidation(b *testing.B) {
	validation, data := data()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validation.Validate(data)
	}
}

func BenchmarkAsyncValidation(b *testing.B) {
	defer runtime.GOMAXPROCS(1)
	cpus := runtime.NumCPU()
	runtime.GOMAXPROCS(cpus)

	validation, data := data()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validation.ValidateAsync(data)
	}
}

func data() (*Validation, map[string]interface{}) {
	validation := NewValidation()
	data := make(map[string]interface{})
	n := 100
	s := make([]int, 0)
	for i := 0; i < n; i++ {
		s = append(s, i)
	}
	for i := 0; i < n; i++ {
		name := fmt.Sprintf("test%d", i)
		validation.Field(name).Required().MaxLength(n).Each().MaxInt(n)
		data[name] = s
	}
	return validation, data
}
