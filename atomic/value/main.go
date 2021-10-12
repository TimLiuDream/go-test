package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	type T struct{ a, b, c int }
	var ta = T{1, 2, 3}
	var v atomic.Value
	v.Store(ta)
	var tb = v.Load().(T)
	fmt.Println(tb)       // {1 2 3}
	fmt.Println(ta == tb) // true

	v.Store(T{0, 9, 8})
	var tc = v.Load().(T)
	fmt.Println(tc)
	v.Store("hello") // 将导致一个恐慌
}
