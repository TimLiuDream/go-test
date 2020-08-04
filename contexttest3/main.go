package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	fmt.Println(time.Now())
	NewT(ctx, 10).run()
	fmt.Println(time.Now())
	select {}
}

type t struct {
	ctx context.Context

	interval int
}

func NewT(ctx context.Context, interval int) *t {
	return &t{ctx: ctx, interval: interval}
}

func (t *t) run() {
	go func() {
		select {
		case <-t.ctx.Done():
			fmt.Println("done2!")
			fmt.Println(t.ctx.Err())
			return
		default:
		}
	}()
	go t.process()
	select {
	case <-t.ctx.Done():
		fmt.Println(t.ctx.Err())
	case <-time.After(5 * time.Second):
		fmt.Println("timeout")
	}
}

func (t *t) process() {
	for i := 0; i < t.interval; i++ {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}
