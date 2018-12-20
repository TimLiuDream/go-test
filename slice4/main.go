package main

import "fmt"

func main() {
	s := []int{1, 3, 4, 7, 87, 3, 6, 3247, 88, 4, 9}
	for index := 0; index < len(s); index++ {
		fmt.Println(s[index])
	}
	for index, value := range s {
		// if index == 5 {
		// 	value = 10
		// }
		s[index] = 10
		fmt.Println("range:", index, value)
	}
	fmt.Println(s)
}
