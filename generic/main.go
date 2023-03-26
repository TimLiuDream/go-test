package main

import (
	"fmt"
)

func Contains[T comparable](items []T, item T) bool {
	for _, v := range items {
		if v == item {
			return true
		}
	}
	return false
}

func main() {
	//fmt.Println(Contains([]string{"a", "b", "c"}, "d"))
	//fmt.Println(Contains([]int{1, 2, 3}, 3))
	runSumAll()
}

type constraint interface {
	~float64 | int
}

func sumAll[T constraint](arr []T) T {
	var s T
	for _, ele := range arr {
		s += ele
	}
	return s
}

func runSumAll() {
	fmt.Println("sum: ", sumAll([]int{1, 2, 3, 5, 6}))
	fmt.Println("sum: ", sumAll([]float64{1.2, 2.1, 3.8, 5.4}))
}
