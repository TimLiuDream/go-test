package main

import "fmt"

// 0 1 zz zz zz 5
// const (
// 	x = iota
// 	y
// 	z = "zz"
// 	k
// 	j
// 	p = iota
// )

// 0 1 2 3
const (
	a, b = iota, iota + 1
	_, _
	c, d
)

func main() {
	fmt.Println(a, b, c, d)
}
