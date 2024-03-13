package main

import "fmt"

func TestSlice(arr []int) {
	fmt.Printf("[TestSlice] %p\n", &arr)
	arr = append(arr, 3)
	arr = append(arr, 4)
	fmt.Printf("[TestSlice] %p\n", &arr)
}

func func1() {
	arr := make([]int, 1, 10)
	arr = append(arr, 1)
	arr = append(arr, 2)
	TestSlice(arr)
	fmt.Printf("[func1] %p\n", &arr)
	fmt.Printf("%d,%d,%d", arr[0], arr[1], len(arr))
}
