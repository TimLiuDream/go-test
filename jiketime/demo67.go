package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	coordinateWithWaitGroup1()
}

func coordinateWithWaitGroup1() {
	total := 12
	stride := 3
	var num int32
	fmt.Printf("the number:%d [with sync.WaitGroup]\n", num)
	var wg sync.WaitGroup
	for i := 1; i <= total; i = i + stride {
		wg.Add(stride)
		for j := 0; j < stride; j++ {
			go addNum1(&num, i+j, wg.Done)
		}
		wg.Wait()
	}
	fmt.Println("end")
}

func addNum1(nump *int32, id int, deferFunc func()) {
	defer func() {
		deferFunc()
	}()
	for i := 0; ; i++ {
		currNum := atomic.LoadInt32(nump)
		newNum := currNum + 1
		time.Sleep(time.Millisecond * 200)
		if atomic.CompareAndSwapInt32(nump, currNum, newNum) {
			fmt.Printf("the number :%d [%d-%d]\n", newNum, id, i)
			break
		} else {
			//fmt.Printf("the cas operation faild.[%d-%d]\n", id, i)
		}
	}
}
