package main

import "fmt"

func main() {
	func2()
}

func func1() {
	index := 5
	for index := 0; index < 19; index++ {
		fmt.Println(index)
	}
	fmt.Println(index)
}

func func2() {
	for i := 1; i < 10; i++ {
		a, b := i, i+1
		fmt.Println(a, b)
	}
}
