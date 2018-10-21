package main

import "fmt"

func main() {
	value1 := [...]int8{0, 1, 2, 3, 4, 5, 6}
	switch int8(1 + 3) {
	case value1[0], value1[1]:
		fmt.Println("0 or 1")
	case value1[2], value1[3]:
		fmt.Println("2 or 3")
	case value1[4], value1[5]:
		fmt.Println("4 or 5")
	default:
		fmt.Println("nothing")
	}
	fmt.Println()

	value2 := [...]int8{0, 1, 2, 3, 4, 5, 6}
	switch value2[4] {
	case 0, 1:
		fmt.Println("0 or 1")
	case 2, 3:
		fmt.Println("2 or 3")
	case 4, 5:
		fmt.Println("4 or 5")
	default:
		fmt.Println("nothing")
	}
}
