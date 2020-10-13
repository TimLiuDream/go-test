package main

import (
	"fmt"
	"sort"

	"github.com/timliudream/go-test/sort/models"
	"github.com/timliudream/go-test/sort/sortlib"
)

func main() {
	func3()
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
	sort.SliceStable(a, func(i, j int) bool {
		return a[i].Age < a[j].Age
	})
	// sort.Stable(a)
	fmt.Println(a)
}

func func3() {
	as := make([]*models.Person, 0)
	as = append(as, &models.Person{
		Name: "AAA",
		Age:  55,
	})
	as = append(as, &models.Person{
		Name: "BBB",
		Age:  22,
	})
	as = append(as, &models.Person{
		Name: "CCC",
		Age:  0,
	})
	as = append(as, &models.Person{
		Name: "DDD",
		Age:  22,
	})
	as = append(as, &models.Person{
		Name: "EEE",
		Age:  11,
	})
	sort.SliceStable(as, func(i, j int) bool {
		return as[i].Age < as[j].Age
	})
	for _, a := range as {
		fmt.Println(a.Name)
		fmt.Println(a.Age)
	}
}
