package main

import (
	"fmt"
)

// 两种答案
// panic: hello

// goroutine 1 [running]:
// main.main()
// 	C:/Users/Administrator/go/src/github.com/timliudream/go-test/Interview/5/interview5.go:16 +0x2e

// 1

func main() {
	// runtime.GOMAXPROCS(1)
	int_chan := make(chan int, 1)
	string_chan := make(chan string, 1)

	select {
	case value := <-string_chan:
		panic(value)
	case value := <-int_chan:
		fmt.Println(value)
	}
}
