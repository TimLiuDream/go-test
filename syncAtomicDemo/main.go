package main

import (
	"fmt"
	"sync"
)

func main() {
	var iteration uint64
	var waitgroup sync.WaitGroup
	for i := 0; i < 1000; i++ {
		waitgroup.Add(1)
		go func() {
			iteration++
			waitgroup.Done()
		}()
	}
	waitgroup.Wait()
	fmt.Println("Counter:", iteration)
}
