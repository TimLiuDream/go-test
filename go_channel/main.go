package main

import (
	"fmt"
)

var (
	intChan    chan int
	tagChan    chan bool
	returnChan chan bool
	maxCount   = 200
)

func main() {
	intChan = make(chan int, 1)
	tagChan = make(chan bool)
	returnChan = make(chan bool)
	go productser()
	go cumsumer()
	<-returnChan
}

// 生产者
func productser() {
	for {
		select {
		case <-tagChan:
			fmt.Println("end product")
			returnChan <- true
			return
		default:
			intChan <- 2
		}
	}
}

// 消费者 满200
func cumsumer() {
	count := 0
	for {
		select {
		case <-intChan:
			if count < maxCount {
				fmt.Println(count)
				count++
			} else {
				close(intChan)
				tagChan <- true
				return
			}
		}
	}
}
