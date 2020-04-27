package main

import "testing"

func BenchmarkCrcHash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		index := i % len(UUIDs)
		CrcHash(UUIDs[index])
	}
}
