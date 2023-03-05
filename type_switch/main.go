package main

import "fmt"

func do(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("value is int type")
	case string:
		fmt.Println("value is string type")
	case bool:
		fmt.Println("value is bool type")
	default:
		fmt.Println("can not found value type")
	}
}

func main() {
	do(1)
	do("joker")
	do(true)
}
