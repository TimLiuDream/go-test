package main

import (
	"time"
	"fmt"
)

func Calculate1(n int) int {
	result := 0
	for i := 1; i <= n; i++ {
		result += i
	}
	return result
}

func Calculate2(n int) int {
	return n * (n + 1) / 2
}

func main() {
	start1 := time.Now()
	Calculate1(100000000)
	fmt.Println(time.Since(start1))
	start2:=time.Now()
	Calculate2(100000000)
	fmt.Println(time.Since(start2))
}
