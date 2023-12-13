package main

import (
	"fmt"
	"github.com/shopspring/decimal"
)

func main() {
	var v1, v2 = 0.1, 0.2
	fmt.Println(v1 + v2)
	// 输出：0.30000000000000004

	func1()
}

func func1() {
	a := 0.1
	b := 0.2
	c := decimal.NewFromFloat(a)
	d := decimal.NewFromFloat(b)
	fmt.Println(a, b, c.String(), d.String())
	fmt.Println(a + b)
	fmt.Println(c.Add(d).String())
}
