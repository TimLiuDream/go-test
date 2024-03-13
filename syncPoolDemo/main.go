package main

import (
	"fmt"
	"sync"
)

var pool = sync.Pool{New: func() interface{} {
	return 0
}}

func main() {
	pool.Put(1)
	pool.Put(2)
	pool.Put(3)

	a := pool.Get().(int)
	b := pool.Get().(int)
	c := pool.Get().(int)
	fmt.Println(a, b, c)
}
