package main

import "fmt"

func main() {
	s := "hello"
	s = "c" + s[1:]
	fmt.Printf("%s\n", s)
}
