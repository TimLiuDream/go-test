package main

import (
	"fmt"
	"time"
)

func func3() {
	ch := make(chan int)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- 1
	}()

	select {
	case <-ch:
		fmt.Println("Received from channel")
	case <-time.After(time.Second):
		fmt.Println("Timeout")
	}
}
