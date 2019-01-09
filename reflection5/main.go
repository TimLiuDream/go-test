package main

import (
	"reflect"
	"fmt"
)

func main() {
	x := 5
	y := 6
	fmt.Println(reflect.TypeOf('0'))
	fmt.Println(reflect.ValueOf('0'))
	fmt.Println(reflect.TypeOf(y - x))
	fmt.Println(reflect.ValueOf(y - x))
	fmt.Println(reflect.TypeOf(y - x + '0'))
	fmt.Println(reflect.ValueOf(y - x + '0'))
}
