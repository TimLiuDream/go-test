package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()

	go process(ctx)
	for {
		// cancel()
		time.Sleep(time.Second)
	}
}

func process(ctx context.Context) {
	myCtx, cancel := context.WithCancel(ctx)
	defer cancel()
	go func(myCtx context.Context) {
		for i := 0; i < 5; i++ {
			select {
			case <-myCtx.Done():
				fmt.Println("bye~", myCtx.Err())
				return
			default:
				time.Sleep(time.Second)
				fmt.Println(i)
			}
		}
	}(myCtx)

	select {
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	case <-time.After(5 * time.Second):
		fmt.Println("timeout")
	}
}
