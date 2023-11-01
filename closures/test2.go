package main

import "fmt"

func adder() func(int) int { // 外部函数
	sum := 0
	return func(x int) int { // 内部函数
		fmt.Println("func sum: ", sum)
		sum += x
		return sum
	}
}

func func2() {
	a := adder()
	fmt.Println(a(1))
	fmt.Println(a(2))
	fmt.Println(a(3))
}
