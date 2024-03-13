package main

import (
	"fmt"
	"unicode/utf8"
)

func func1() {
	name := "张三"
	fmt.Printf("%d\n", len(name))
	fmt.Printf("%d\n", utf8.RuneCountInString(name))
}
