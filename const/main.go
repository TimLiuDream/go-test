package main

import "fmt"

func main() {
	const c = 8
	a := &c
	*a = 12
	fmt.Println(*a)
}
