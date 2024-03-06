package main

import (
	"fmt"
	"runtime"
	"time"
)

func Foo() {
	fmt.Println("打印1")
	defer fmt.Println("打印2")
	runtime.Goexit() // 加入这行
	fmt.Println("打印3")
}

func main() {
	go Foo()
	fmt.Println("打印4")
	time.Sleep(2 * time.Second)
}
