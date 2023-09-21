package main

import (
	"fmt"
	"time"
)

func main() {
	TimeOut()
}

func CreateTimer() {
	timer := time.NewTimer(2 * time.Second)
	fmt.Println("Timer created.")

	<-timer.C // 阻塞
	fmt.Println("Timer expired.")
}

func StopTimer() {
	timer := time.NewTimer(2 * time.Second)
	fmt.Println("Timer created.")

	go func() {
		<-timer.C
		fmt.Println("Timer expired.")
	}()

	time.Sleep(1 * time.Second)
	stopped := timer.Stop()
	if stopped {
		fmt.Println("Timer stopped.")
	} else {
		fmt.Println("Timer has already expired.")
	}
}

func ResetTimer() {
	timer := time.NewTimer(10 * time.Second)
	fmt.Printf("time: %d, Timer created.\n", time.Now().Unix())

	time.Sleep(2 * time.Second)
	reset := timer.Reset(3 * time.Second)
	if reset {
		fmt.Printf("time: %d, Timer reset.\n", time.Now().Unix())
	} else {
		fmt.Printf("time: %d, Timer has already expired.\n", time.Now().Unix())
	}
	<-timer.C // 阻塞
	fmt.Printf("time: %d, Timer expired again.\n", time.Now().Unix())
}

func CompareResetAndStop() {
	timer := time.NewTimer(5 * time.Second)
	fmt.Printf("time: %d, Timer created.\n", time.Now().Unix())

	go func() {
		<-timer.C
		fmt.Printf("time: %d, Timer expired.\n", time.Now().Unix())
	}()

	time.Sleep(2 * time.Second)
	timer.Reset(3 * time.Second)
	fmt.Printf("time: %d, Timer reset.\n", time.Now().Unix())

	time.Sleep(2 * time.Second)
	timer.Stop()
	fmt.Printf("time: %d, Timer stopped.\n", time.Now().Unix())
}

func Tick() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	go func() {
		for range ticker.C {
			fmt.Printf("time: %d, Ticker ticked!\n", time.Now().Unix())
		}
	}()

	time.Sleep(5 * time.Second)
}

func TimeOut() {
	ch := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- "Operation completed."
	}()

	select {
	case msg := <-ch:
		fmt.Println(msg)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout reached.")
	}
}
