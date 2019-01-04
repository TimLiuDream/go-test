package main

import "fmt"

func main() {
	m := map[int]string{1: "1", 2: "2", 3: "3"}

	for key, value := range m {
		fmt.Printf("key:%d----value:%s", key, value)
	}
}
