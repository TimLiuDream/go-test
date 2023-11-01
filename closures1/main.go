package main

import "fmt"

func makeAdder1(x1 int) func(int) int {
	return func(y1 int) int {
		return x1 + y1
	}
}

func makeAdder2(x2 int) func(int) int {
	fmt.Println(x2)
	return func(y2 int) int {
		return x2 + y2
	}
}

func main() {
	a := makeAdder1(5)
	fmt.Println(a(1))

	b := makeAdder2(6)
	fmt.Println(b(1))
}
