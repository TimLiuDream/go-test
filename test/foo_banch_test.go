package test

import "testing"

func BenchmarkFoo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		foo()
	}
}

func foo() {
	arr := make([]int, 0)
	arr = append(arr, 0)
}
