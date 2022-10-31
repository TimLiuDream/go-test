package main

import "fmt"

func main() {
	fmt.Println(func5(10))
}

func func1() {
	index := 5
	for index := 0; index < 19; index++ {
		fmt.Println(index)
	}
	fmt.Println(index)
}

func func2() {
	for i := 1; i < 10; i++ {
		a, b := i, i+1
		fmt.Println(a, b)
	}
}

func func3(n int) int {
	result := 0
	for i := 0; i <= n; i++ {
		result += i
	}
	return result
}

func func4(n int) int {
	sum := 0
	for i := 1; i < n; i++ {
		for j := 1; j < n; j++ {
			sum = sum + i
		}
	}
	return sum
}

func func5(n int) int {
	var i, count = 1, 0
	for i <= n {
		i = i * 2
		count++
	}
	return count
}

func func6(n int) {
	sum_1, a := 0, 1
	for a <= 100 {
		sum_1 += a
	}

	sum_2, b := 0, 1
	for b <= n {
		sum_2 += b
	}

	sum_3 := 0
	for i := 1; i <= n; i++ {
		for j := 1; j < n; j++ {
			sum_3 += i + j
		}
	}
}
