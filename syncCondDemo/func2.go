package main

import (
	"fmt"
	"sync"
	"time"
)

var mtx2 sync.Mutex
var cond2 = sync.NewCond(&mtx2)

func dummyGoroutine2(id int) {
	cond2.L.Lock()
	defer cond2.L.Unlock()

	fmt.Printf("Goroutine %d is waitting...\n", id)
	cond2.Wait()
	fmt.Printf("Goroutine %d received the signal.\n", id)
}

func func2() {
	go dummyGoroutine2(1)
	go dummyGoroutine2(2)

	time.Sleep(time.Second)
	cond2.Broadcast()
	time.Sleep(time.Second)
}
