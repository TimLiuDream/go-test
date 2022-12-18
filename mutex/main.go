package main

import (
	"fmt"
	"sync"
	"time"
)

func func1() {

	const loop = 100
	var wg sync.WaitGroup
	wg.Add(loop * 2)

	// declaring a shared value
	var n int = 0

	for i := 0; i < loop; i++ {
		go func() {
			time.Sleep(time.Second / 10)
			n++
			wg.Done()
		}()
		go func() {
			time.Sleep(time.Second / 10)
			n--
			wg.Done()
		}()
	}
	wg.Wait()
	// printing the final value of n
	if n != 0 {
		fmt.Println("The Final value of n should be 0. But found ", n)
		return
	}
	fmt.Printf("\nFinal value of n is %d\n\n", n) // the final of n should be 0
}
