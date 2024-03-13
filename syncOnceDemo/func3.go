package main

import (
	"fmt"
	"sync"
	"time"
)

func func3() {
	var i int32
	var mut sync.Mutex
	pool := &sync.Pool{
		New: func() interface{} {
			mut.Lock()
			defer mut.Unlock()
			i = i + 1
			time.Sleep(time.Millisecond)
			return i
		},
	}
	numberWorker := 10
	var wg sync.WaitGroup
	wg.Add(numberWorker)
	for j := 0; j < numberWorker; j++ {
		go func() {
			defer wg.Done()
			obj := pool.Get()
			fmt.Printf("obj=%d\n", obj)
			defer pool.Put(obj)
		}()
	}
	wg.Wait()
}
