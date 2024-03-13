package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var n int32
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			fmt.Printf("i: %d\n", i)
			atomic.AddInt32(&n, 1)
			fmt.Printf("n: %d\n", atomic.LoadInt32(&n))
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println(atomic.LoadInt32(&n))
}
