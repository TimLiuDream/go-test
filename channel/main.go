package main

import (
	"fmt"
	"sync"
)

func main() {
	//go func1()
	//time.Sleep(5 * time.Second)
	//go func2()
	//time.Sleep(5 * time.Second)
	func5()
}

var testChan chan bool

func func1() {
	fmt.Println("in func1")
	fmt.Printf("1:%p\n", testChan)
	v := <-testChan
	fmt.Println("in func1 get value")
	fmt.Println(v)
}

func func2() {
	fmt.Println("in func2")
	testChan = make(chan bool)
	fmt.Printf("2:%p\n", testChan)
	testChan <- true
	fmt.Printf("3:%p\n", testChan)
	fmt.Println("func2 init chan an send value")
}

func func3() {
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Println(x)
		case ch <- i:
		}
	}
}

func printOdd(wg *sync.WaitGroup, oddCh chan int, evenCh chan int) {
	defer wg.Done()

	for i := 1; i <= 100; i += 2 {
		<-oddCh // 阻塞等待偶数协程发送信号
		fmt.Println("奇数:", i)
		evenCh <- i // 发送信号给偶数协程
	}
}

func printEven(wg *sync.WaitGroup, oddCh chan int, evenCh chan int) {
	defer wg.Done()

	for i := 2; i <= 100; i += 2 {
		<-evenCh // 阻塞等待奇数协程发送信号
		fmt.Println("偶数:", i)
		if i != 100 {
			oddCh <- i // 发送信号给奇数协程
		}
	}
}

func func4() {
	var wg sync.WaitGroup
	oddCh := make(chan int)
	evenCh := make(chan int)

	wg.Add(2)

	// 启动奇数协程
	go printOdd(&wg, oddCh, evenCh)

	// 启动偶数协程
	go printEven(&wg, oddCh, evenCh)

	// 发送初始信号给奇数协程
	oddCh <- 1

	wg.Wait()
}

func printNumber(wg *sync.WaitGroup, oddChan chan int, evenChan chan int) {
	defer wg.Done()
	for i := 0; i <= 100; i++ {
		if i%2 != 0 {
			oddChan <- i
		} else {
			evenChan <- i
		}
	}

	close(oddChan)
	close(evenChan)
}

func func5() {
	var wg sync.WaitGroup
	oddChan := make(chan int)
	evenChan := make(chan int)
	wg.Add(1)

	go printNumber(&wg, oddChan, evenChan)

	for {
		select {
		case oddV, ok := <-oddChan:
			if !ok {
				oddChan = nil
				break
			}
			fmt.Println("奇数：", oddV)
		case evenV, ok := <-evenChan:
			if !ok {
				evenChan = nil
				break
			}
			fmt.Println("偶数：", evenV)
		}
		if oddChan == nil || evenChan == nil {
			break
		}
	}

	wg.Wait()
}
