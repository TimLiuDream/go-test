package main

import (
	"GolangTraining/SortTraining/SortLib"
	"fmt"
)

func main() {
	s := []int{1, 5, 8, 3, 2, 1, 5, 8, 2, 4}
	SortLib.BubbleSort(s)
	fmt.Println(s)

	s1 := []int{1, 5, 8, 3, 2, 1, 5, 8, 2, 4}
	SortLib.InsertionSort(s1)
	fmt.Println(s1)
}
