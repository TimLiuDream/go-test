package main

import "fmt"

type Calculator struct {
	add func(int, int) int
}

func NewCalculator() *Calculator {
	c := &Calculator{}
	c.add = func(x, y int) int {
		fmt.Printf("func x: %d, y: %d\n", x, y)
		return x + y
	}
	return c
}

func (c *Calculator) Add(x, y int) int {
	return c.add(x, y)
}

func func1() {
	calc := NewCalculator()
	fmt.Println(calc.Add(1, 2))
	fmt.Println(calc.Add(2, 3))
}
