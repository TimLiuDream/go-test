package main

import (
	"fmt"
)

// 打印后
// 打印中
// 打印前
// panic: 触发异常

// goroutine 1 [running]:
// main.defer_call()
// 	C:/Users/Administrator/go/src/github.com/timliudream/go-test/Interview/1/interview1.go:16 +0xe7
// main.main()
// 	C:/Users/Administrator/go/src/github.com/timliudream/go-test/Interview/1/interview1.go:8 +0x27

func main() {
	defer_call()
}

func defer_call() {
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()

	panic("触发异常")
}
