package main

import (
	"fmt"
	"time"
)

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

func main() {
	go func1()
	time.Sleep(5 * time.Second)
	go func2()
	time.Sleep(5 * time.Second)
}
