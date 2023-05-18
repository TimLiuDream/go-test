package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {
	func6()
}

func func1(s string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched()
		fmt.Println(s)
	}
}

func func2() {
	a := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)
}

func sum(a []int, c chan int) {
	total := 0
	for _, v := range a {
		total += v
	}
	c <- total
}

func func3() {
	c := make(chan int, 2)

	c <- 1
	c <- 2

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(c)
}

func func4() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}

func fibonacci(n int, c chan int) {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func func6() {
	// 使用 goroutine 和 channel 实现一个每隔一秒输出 "hello, world" 的程序。
	var length = 3
	c := make(chan string)
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		for i := 0; i < length; i++ {
			c <- "hello, world"
			time.Sleep(time.Second)
		}
		cancel()
	}()
	for {
		select {
		case <-ctx.Done():
			return
		case s := <-c:
			fmt.Printf("time is %d, %s\n", time.Now().Unix(), s)
		}
	}
}
