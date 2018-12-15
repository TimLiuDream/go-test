package main

import (
	"fmt"
	"reflect"
)

func main() {
	x := 123
	v := reflect.ValueOf(&x)
	v.Elem().SetInt(5256)
	fmt.Println(x)
}
