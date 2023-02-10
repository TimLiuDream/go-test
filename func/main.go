package main

import "fmt"

type testInt func(int) bool

func isOdd(integer int) bool {
	if integer%2 == 0 {
		return false
	}
	return true
}

func isEven(integer int) bool {
	if integer%2 == 0 {
		return true
	}
	return false
}

func filter(slice []int, f testInt) []int {
	var result []int
	for _, value := range slice {
		if f(value) {
			result = append(result, value)
		}
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func SumAndProduct(A, B int) (add int, Multiplied int) {
	add = A + B
	Multiplied = A * B
	return
}

func main() {
	//x := 3
	//y := 4
	//z := 5
	//
	//max_xy := max(x, y)
	//max_xz := max(x, z)
	//
	//fmt.Printf("max(%d,%d) = %d\n", x, y, max_xy)
	//fmt.Printf("max(%d,%d) = %d\n", x, z, max_xz)
	//fmt.Printf("max(%d,%d) = %d\n", y, z, max(y, z))
	//
	//xPLUSy, xTIMESy := SumAndProduct(x, y)
	//
	//fmt.Printf("%d+%d=%d\n", x, y, xPLUSy)
	//fmt.Printf("%d*%d=%d\n", x, y, xTIMESy)
	//
	//slice := []int{1, 2, 3, 4, 5, 7}
	//fmt.Println("slice = ", slice)
	//odd := filter(slice, isOdd)
	//fmt.Println("ODD element of slice are :", odd)
	//even := filter(slice, isEven)
	//fmt.Println("even element of slice are :", even)
	if func1 != nil {
		fmt.Println(func1(1, 2))
	} else {
		fmt.Println("func1 is nil")
	}
}
