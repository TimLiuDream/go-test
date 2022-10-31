package main

import "fmt"

func main() {
	const c = 8
	a := &c
	*a = 12
	fmt.Println(*a)
}

const cl = 100

var bl = 123

func func2() {
	fmt.Println(&bl, bl)
	fmt.Println(&cl, cl) // Cannot take the address of 'cl'
}
