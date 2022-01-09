package main

import "fmt"

// [1]

func main() {
	list := make([]int, 0)
	list = append(list, 1)
	fmt.Println(list)
}
