package main

import (
	"fmt"
	"strconv"
)

func main() {
	func1()
}

func func1() {
	o := 5500000
	n := 10000000
	v := o / n
	fmt.Println(v)
}

func func2() {
	fixedValue := ""
	value := "902000"
	rawFloat, err := strconv.ParseFloat(value, 64)
	if err != nil {
		fixedValue = "0"
		fmt.Println(fixedValue)
		return
	}
	f := rawFloat / 100000
	fixedValue = fmt.Sprintf("%g", f)
	fmt.Println(fixedValue)
}
