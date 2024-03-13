package main

import (
	"fmt"
	"sync"
)

func func3() {
	m := map[int]int{}
	mux := &sync.RWMutex{}

	go writeLoop(m, mux)
	go readLoop(m, mux)
	go readLoop(m, mux)
	go readLoop(m, mux)

	block := make(chan int)
	<-block
}

func writeLoop(m map[int]int, mux *sync.RWMutex) {
	for {
		for i := 0; i < 100; i++ {
			mux.Lock()
			m[i] = i
			mux.Unlock()
		}
	}
}

func readLoop(m map[int]int, mux *sync.RWMutex) {
	for {
		mux.RLock()
		for k, v := range m {
			fmt.Println(k, v)
		}
		mux.RUnlock()
	}
}
