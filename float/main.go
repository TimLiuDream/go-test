package main

import (
	"fmt"
	"strconv"
)

func main() {
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
