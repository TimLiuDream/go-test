package main

import (
	"context"
	"testing"
	"time"
)

func TestProcess(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	isSuccess := false
	go func(ctx context.Context) {
		isSuccess = process(t)
	}(ctx)
	select {
	case <-ctx.Done():
		if isSuccess {
			t.Log("process successfully!")
		} else {
			t.Log("timeout!")
		}
		return
	case <-time.After(6 * time.Second):
		t.Log("timeout!")
		return
	}
}

func process(t *testing.T) bool {
	t.Log("start process!")
	time.Sleep(5 * time.Second)
	t.Log("complete process function!")
	return true
}
