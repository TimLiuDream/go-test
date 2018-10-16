package main

import "fmt"

func main() {
	array1 := [3]string{"a", "b", "c"}
	fmt.Printf("the array:%v\n", array1)

	array2 := modifyArray(array1)
	fmt.Printf("the array:%v\n", array1)
	fmt.Printf("the array:%v\n", array2)
}

func modifyArray(a [3]string) [3]string {
	a[1] = "x"
	return a
}
