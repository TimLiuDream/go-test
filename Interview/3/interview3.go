package main

import (
	"fmt"
	"sync"
)

func main() {
	func2()
}

// A:  2
// A:  3
// A:  10
// A:  10
// A:  10
// A:  10
// A:  10
// A:  10
// A:  10
// A:  10
func func1() {
	//runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("A: ", i)
			wg.Done()
		}()
	}
	//for i := 0; i < 10; i++ {
	//	go func(i int) {
	//		fmt.Println("B: ", i)
	//		wg.Done()
	//	}(i)
	//}
	wg.Wait()
}

// B:  9
// B:  0
// B:  1
// B:  2
// B:  3
// B:  4
// B:  5
// B:  6
// B:  7
// B:  8
func func2() {
	// runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			fmt.Println("B: ", i)
			wg.Done()
		}()
	}
	wg.Wait()
}
