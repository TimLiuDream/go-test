package main

import (
	"context"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/ping", pingTest)
	_ = r.Run()
}

func pingTest(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()
	if process(ctx) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	} else {
		c.JSON(500, gin.H{
			"message": "fail",
		})
	}
}

func process(ctx context.Context) bool {
	myCtx, cancel := context.WithCancel(ctx)
	defer cancel()
	go func(myCtx context.Context) {
		time.Sleep(7 * time.Second)
		select {
		case <-myCtx.Done():
			fmt.Println(time.Now())
			fmt.Println("bye~", myCtx.Err())
			return
		default:
			time.Sleep(time.Second)
			fmt.Println("i")
		}
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
	}(myCtx)

	select {
	case <-ctx.Done():
		fmt.Println(ctx.Err())
		fmt.Println("done")
		return false
	case <-time.After(5 * time.Second):
		fmt.Println(time.Now())
		fmt.Println("timeout")
		cancel()
	}
	return true
}
