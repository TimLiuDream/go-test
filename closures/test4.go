package main

import "fmt"

func makeCounter() func() int {
	i := 0
	return func() int {
		fmt.Println("func i: ", i)
		i++
		return i
	}
}

func func4() {
	counter := makeCounter()
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
}
