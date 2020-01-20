package main

import "fmt"

func main() {
	slice := []int{1, 9, 10, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 0}
	newSlice := UniqueIntSlice(slice...)
	fmt.Println(newSlice)
}

func UniqueIntSlice(slice ...int) (newSlice []int) {
	found := make(map[int]bool)
	for _, val := range slice {
		if _, ok := found[val]; !ok {
			found[val] = true
			newSlice = append(newSlice, val)
		}
	}
	return
}
