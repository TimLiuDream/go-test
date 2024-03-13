package main

import (
	"fmt"
)

func func71() (val int) {
	val = 10
	defer func() {
		val += 1
	}()
	return 100
}

func func7() {
	fmt.Println(func71())
}
