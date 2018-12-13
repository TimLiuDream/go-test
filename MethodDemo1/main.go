package main

import (
	"math"
	"fmt"
)

type Rectangle struct {
	width, height float64
}

type Circle struct {
	rdius float64
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func (c Circle) area() float64 {
	return c.rdius * c.rdius * math.Pi
}

func main() {
	r1 := Rectangle{12, 2}
	r2 := Rectangle{9, 4}
	c1 := Circle{5}
	c2 := Circle{10}
	fmt.Println("Area of r1 is ", r1.area())
	fmt.Println("Area of r2 is ", r2.area())
	fmt.Println("Area of c1 is ", c1.area())
	fmt.Println("Area of c2 is ", c2.area())
}
