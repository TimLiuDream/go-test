package main

import (
	"fmt"
	"time"
)

const (
	MaxBufferSize = 5
	NumProducers  = 2
	NumConsumers  = 3
	NumItems      = 10
)

func producer(id int, buffer chan<- int) {
	for i := 0; i < NumItems; i++ {
		item := id*10 + i
		buffer <- item
		fmt.Printf("produce id: %d, item: %d\n", id, item)
		time.Sleep(100 * time.Millisecond)
	}
}

func consumer(id int, buffer <-chan int) {
	for i := 0; i < NumItems; i++ {
		item := <-buffer
		fmt.Printf("comsumer id: %d, item: %d\n", id, item)
		time.Sleep(200 * time.Millisecond)
	}
}

func main() {
	buffer := make(chan int, MaxBufferSize)
	for i := 0; i < NumProducers; i++ {
		go producer(i, buffer)
	}
	for i := 0; i < NumConsumers; i++ {
		go consumer(i, buffer)
	}
	time.Sleep(5 * time.Second)
	close(buffer)
	time.Sleep(5 * time.Second)
}
