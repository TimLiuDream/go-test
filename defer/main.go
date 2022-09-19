package main

import (
	"fmt"
	"log"
	"runtime/debug"
)

func main() {
	//testFunc1()
	func2()
}

func goFun() {
	defer func() {
		if p := recover(); p != nil {
			log.Printf("%s\n%s", p, debug.Stack())
			log.Printf("%s", debug.Stack())
		}
	}()

	panic("ones")
}

func func1() {
	var err error
	for i := 0; i < 100; i++ {
		defer func() {
			if err != nil {
				fmt.Println("error hanppen")
			}
		}()
		fmt.Println(i)
		if i == 50 {
			err = fmt.Errorf("ones!")
			continue
		}
	}
}

func testFunc1() {
	go goFun()

	select {}
}

type temp struct{}

func (t *temp) Add(elem int) *temp {
	fmt.Println(elem)
	return &temp{}
}

func func2() {
	tt := &temp{}
	defer tt.Add(1).Add(2)
	tt.Add(3)
}
