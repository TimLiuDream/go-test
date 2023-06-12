package main

import (
	"os"
	"runtime/pprof"
	"testing"
)

// becnchmark test
func BenchmarkPreAllocStringsBuilder(b *testing.B) {
	f, err := os.Create("myfunction.pprof")
	if err != nil {
		b.Fatal(err)
	}
	defer f.Close()

	if err := pprof.StartCPUProfile(f); err != nil {
		b.Fatal(err)
	}
	defer pprof.StopCPUProfile()
	for i := 0; i < b.N; i++ {
		PreAllocStringsBuilder(1000, "aaaa")
	}
}
