package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "1234567"
	if strings.Contains(s, "1") {
		fmt.Println("1")
	} else if strings.Contains(s, "2") {
		fmt.Println("2")
	}
}
