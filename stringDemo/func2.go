package main

import "fmt"

var item = "hello"

func func2() {
	v := item
	v[0] = 'a'
	fmt.Printf("%s", item)
}
