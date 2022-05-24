package main

import "fmt"

func Contains[T comparable](items []T, item T) bool {
	for _, v := range items {
		if v == item {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(Contains([]string{"a", "b", "c"}, "d"))
	fmt.Println(Contains([]int{1, 2, 3}, 3))
}
