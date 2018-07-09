package main

import "fmt"

func main() {
	var arr [10]int
	arr[0] = 42
	arr[1] = 13
	a := [3]int{1, 2, 3}
	b := [10]int{4, 5, 6}
	c := [...]int{7, 8, 9}
	fmt.Printf("the first element is %d\n", arr[0])
	fmt.Printf("the last element is %d\n", arr[9])
	fmt.Println(a[2])
	fmt.Println(b[4])
	fmt.Println(c[1])

	doubleArray := [2][4]int{[4]int{1, 2, 3, 4}, [4]int{5, 6, 7, 8}}
	easyArray := [2][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}}
	fmt.Println(doubleArray[0][2])
	fmt.Println(easyArray[0][1])
}
