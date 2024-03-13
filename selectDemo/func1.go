package main

import "fmt"

func func1() {
	ch := make(chan int, 1)

	ch <- 1 // 向通道写入数据，此时通道未满，操作不会被阻塞
	fmt.Println("Data written to channel")

	select {
	case ch <- 2: // 尝试向已满的通道再次写入数据，由于通道已满，操作会被立即返回
		fmt.Println("Data written to channel")
	default:
		fmt.Println("Channel is full, unable to write data")
	}

	data, ok := <-ch // 尝试从通道读取数据，此时通道中有数据，操作不会被阻塞
	if ok {
		fmt.Println("Data read from channel:", data)
	}

	select {
	case data, ok := <-ch: // 尝试从空的通道读取数据，由于通道为空，操作会被立即返回
		if ok {
			fmt.Println("Data read from channel:", data)
		} else {
			fmt.Println("Channel is empty, unable to read data")
		}
	default:
		fmt.Println("Channel is empty, unable to read data default")
	}
}
