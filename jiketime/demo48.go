package main

import "fmt"

func main() {
	fmt.Println("enter function main.")
	caller1()
	fmt.Println("exit function main")
}

func caller1() {
	fmt.Println("enter function caller1")
	caller2()
	fmt.Println("exit function caller1")
}

func caller2() {
	fmt.Println("enter function caller2")
	s1 := []int{0, 1, 2, 3, 4}
	e5 := s1[5]
	_ = e5
	fmt.Println("exit function caller2")
}
