package main

import (
	"fmt"
	"sync"
	"time"
)

func func1() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		fmt.Println(time.Now(), "start")
		time.Sleep(time.Second)
		fmt.Println(time.Now(), "done")
	}()

	wg.Wait()
	fmt.Println(time.Now(), "exiting...")
}
