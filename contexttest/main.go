package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	//ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	//defer cancel()
	//
	//go process(ctx)
	//for {
	//	time.Sleep(time.Second)
	//}

	func2()
}

func process(ctx context.Context) {
	myCtx, cancel := context.WithCancel(ctx)
	defer cancel()
	go func(myCtx context.Context) {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
			time.Sleep(time.Second)
			// 	select {
			// 	case <-myCtx.Done():
			// 		fmt.Println("bye~", myCtx.Err())
			// 		return
			// 	default:
			// 		time.Sleep(time.Second)
			// 		fmt.Println(i)
			// 	}
		}
	}(myCtx)

	select {
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	case <-time.After(5 * time.Second):
		fmt.Println("timeout")
	}
}

func func2() {
	ctx, _ := context.WithTimeout(context.Background(), 0)
	<-ctx.Done()
	fmt.Println("timed out")
}

func func3() {

}
