package main

import (
	"errors"
	"fmt"
	"log"
	"runtime/debug"
)

func main() {
	func6()
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

// 当异常被抛出时，程序的控制流会立即跳转到最近的恢复点，也就是包含当前 defer 语句的代码块的末尾。
func defer_call() {
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()

	panic("触发异常")
}

func func3() (r int) {
	i := 1
	defer func() {
		i = i + 1
	}()
	return i
}

func func4() (r int) {
	defer func(rq int) {
		r = r + 2
	}(r)
	return 2
}

func func5() (r int) {
	defer func(r *int) {
		*r = *r + 2
	}(&r)
	return 2
}

func func6() {
	e1()
	e2()
	e3()
}

func e1() {
	var err error
	defer fmt.Println(err)
	err = errors.New("e1 defer err")
}

func e2() {
	var err error
	defer func() {
		fmt.Println(err)
	}()
	err = errors.New("e2 defer err")
}

func e3() {
	var err error
	defer func(err error) {
		fmt.Println(err)
	}(err)
	err = errors.New("e3 defer err")
}
