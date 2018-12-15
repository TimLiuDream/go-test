package main

import "fmt"

func main() {
	complexArray1 := [3][]string{
		[]string{"d", "e", "f"},
		[]string{"g", "h", "i"},
		[]string{"j", "k", "l"},
	}

	fmt.Printf("The array: %v\n", complexArray1)
	array2 := modifyArray(complexArray1)
	fmt.Printf("The modified array: %v\n", array2)
	fmt.Printf("The original array: %v\n", complexArray1)
	array3 := modifyArray1(complexArray1)
	fmt.Printf("The modified array: %v\n", array3)
	fmt.Printf("The original array: %v\n", complexArray1)
}

func modifyArray(a [3][]string) [3][]string {
	a[1] = []string{"d", "e", "p"}
	return a
}

func modifyArray1(a [3][]string) [3][]string {
	a[1][1] = "v"
	return a
}
