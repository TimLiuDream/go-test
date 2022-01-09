package main

import "fmt"

// map[1:1 3:3]

func main() {
	m := make(map[int]int)
	m[1] = 1
	m[2] = 2
	m[3] = 3
	m[4] = 2
	for k, v := range m {
		if v == 2 {
			delete(m, k)
		}
	}
	fmt.Println(m)
}
