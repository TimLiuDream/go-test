package main

import "fmt"

func main() {
	type MyInt1 int
	type MyInt2 = int
	var i int = 9
	var i1 MyInt1 = MyInt1(i)
	var i2 MyInt2 = i
	fmt.Println(i1, i2)
}
