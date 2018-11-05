package main

import "fmt"

func main() {
	month := 9

	countSlice := make([]int, month)
	for i := 0; i < month; i++ {
		if i == 0 || i == 1 {
			countSlice[i] = 1
		} else {
			countSlice[i] = countSlice[i-1] + countSlice[i-2]
		}
		fmt.Printf("第%d个月的兔子总数为：%d\n", i+1, countSlice[i])
	}
}
