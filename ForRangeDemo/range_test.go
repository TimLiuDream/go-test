package main

import "testing"

const N = 1000

func initSlice() []string {
	s := make([]string, N)
	for i := 0; i < N; i++ {
		s[i] = "www.flysnow.org"
	}
	return s
}

func BenchmarkForSlice(b *testing.B) {
	s := initSlice()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ForSlice(s)
	}
}

func BenchmarkRangeForSlice(b *testing.B) {
	s := initSlice()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RangeForSlice(s)
	}
}
