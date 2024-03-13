package main

import (
	"fmt"
	"sync"
	"time"
)

var mtx1 sync.Mutex
var cond1 = sync.NewCond(&mtx1)

func dummyGoroutine(id int) {
	cond1.L.Lock()
	defer cond1.L.Unlock()

	fmt.Printf("Goroutine %d is waitting...\n", id)
	cond1.Wait()
	fmt.Printf("Goroutine %d received the signal.\n", id)
}

func func1() {
	go dummyGoroutine(1)
	time.Sleep(time.Second)
	fmt.Println("Sending signal...")
	cond1.Signal()
	time.Sleep(time.Second)
}
