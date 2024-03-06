package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan bool)
	go func() {
		testSelectFor1(ch)
	}()

	close(ch)
	time.Sleep(5 * time.Second)
}
func testSelectFor1(chExit chan bool) {
	for {
		select {
		case v, ok := <-chExit:
			if !ok {
				fmt.Println("close channel2", v)
				goto Exit
			}
			fmt.Println("ch2val=", v)
		}
	}

Exit:
	fmt.Println("exit testSelectFor2")
}

func testSelectFor2(chExit chan bool) {
Exit:
	for {
		select {
		case v, ok := <-chExit:
			if !ok {
				fmt.Println("close channel2", v)
				break Exit
				//goto Exit
			}
			fmt.Println("ch2val=", v)
		}
	}

	//Exit:
	fmt.Println("exit testSelectFor2")
}
