package main

import (
	"GolangTraining/sort/sortlib"
	"fmt"
)

func main() {
	s := []int{1, 5, 8, 3, 2, 1, 5, 8, 2, 4}
	sortlib.BubbleSort(s)
	fmt.Println(s)

	s1 := []int{1, 5, 8, 3, 2, 1, 5, 8, 2, 4}
	sortlib.InsertionSort(s1)
	fmt.Println(s1)
}
