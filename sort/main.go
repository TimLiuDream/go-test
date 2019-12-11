package main

import (
	"fmt"
	"sort"

	"github.com/timliudream/golangtraining/sort/models"
	"github.com/timliudream/golangtraining/sort/sortlib"
)

func main() {

}

func func1() {
	s := []int{1, 5, 8, 3, 2, 1, 5, 8, 2, 4}
	sortlib.BubbleSort(s)
	fmt.Println(s)

	s1 := []int{1, 5, 8, 3, 2, 1, 5, 8, 2, 4}
	sortlib.InsertionSort(s1)
	fmt.Println(s1)
}

func func2() {
	a := models.PersonSlice{
		{
			Name: "AAA",
			Age:  55,
		},
		{
			Name: "BBB",
			Age:  22,
		},
		{
			Name: "CCC",
			Age:  0,
		},
		{
			Name: "DDD",
			Age:  22,
		},
		{
			Name: "EEE",
			Age:  11,
		},
	}
	sort.Stable(a)
	fmt.Println(a)
}
