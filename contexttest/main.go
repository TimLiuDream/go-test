package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	go process(ctx)
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		fmt.Printf("count:%d\n", i)
	}
	cancel()
	fmt.Println("process timeout!")
	for i := 0; i < 10; i++ {
		fmt.Printf("test:%d\n", i)
		time.Sleep(time.Second)
	}
}

func process(ctx context.Context) {
	fmt.Println("start process!")
	for i := 0; i < 10; i++ {
		select {
		case <-ctx.Done():
			return
		case <-time.After(time.Second):
			process1(i)
		}
	}
	fmt.Println("complete process function!")
}

func process1(i int) {
	fmt.Printf("process:%d\n", i)
}
