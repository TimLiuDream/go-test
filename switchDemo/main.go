package main

import (
	"fmt"
	"reflect"
)

func main() {
	var object interface{}
	fmt.Println(reflect.TypeOf(object))
	object = interface{}(1)
	fmt.Println(reflect.TypeOf(object))
	switch object {
	case 1:
		fmt.Printf("1")
	case 2:
		fmt.Printf("2")
	case 3:
		fmt.Printf("3")
		break
	case 4:
		fmt.Printf("4")
		break
	default:
		fmt.Printf("5 ")
	}
}
