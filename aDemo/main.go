package main

import (
	"fmt"
	"sync"
)

//
//func printOdd(wg *sync.WaitGroup, oddChan chan int, evenChan chan int) {
//	defer wg.Done()
//	for i := 1; i <= 100; i += 2 {
//		<-oddChan
//		fmt.Println("奇数：", i)
//		evenChan <- 1
//	}
//}
//
//func printEven(wg *sync.WaitGroup, oddChan chan int, evenChan chan int) {
//	defer wg.Done()
//	for i := 0; i <= 100; i += 2 {
//		<-evenChan
//		fmt.Println("偶数：", i)
//		if i != 100 {
//			oddChan <- 1
//		}
//	}
//}
//
//func main() {
//	var wg sync.WaitGroup
//	oddChan := make(chan int)
//	evenChan := make(chan int)
//	wg.Add(2)
//
//	go printOdd(&wg, oddChan, evenChan)
//	go printEven(&wg, oddChan, evenChan)
//
//	evenChan <- 1
//
//	wg.Wait()
//}

func printNumber(wg *sync.WaitGroup, oddChan chan int, evenChan chan int) {
	defer wg.Done()
	for i := 0; i <= 100; i++ {
		if i%2 == 0 {
			evenChan <- i
		} else {
			oddChan <- i
		}
	}
	close(oddChan)
	close(evenChan)
}

func main() {
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
			fmt.Println("奇数：:", oddV)
		case evenV, ok := <-evenChan:
			if !ok {
				evenChan = nil
				break
			}
			fmt.Println("偶数：:", evenV)
		}
		if oddChan == nil || evenChan == nil {
			break
		}
	}

	wg.Wait()
}
