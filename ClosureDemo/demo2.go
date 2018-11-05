package main

import "fmt"

func f(i int) func() int{
	return func() int {
		i++
		return i
	}
}

func main() {
	c1:=f(0)
	c2:=f(0)
	fmt.Println(c1())
	fmt.Println(c2())
}