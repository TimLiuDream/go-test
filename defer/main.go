package main

import (
	"fmt"
	"runtime/debug"

	"github.com/siddontang/go/log"
)

func main() {
	go goFun()

	select {}
}

func goFun() {
	defer func() {
		if p := recover(); p != nil {
			log.Errorf("%s\n%s", p, debug.Stack())
			log.Errorf("%s", debug.Stack())
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
