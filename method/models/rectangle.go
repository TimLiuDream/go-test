package models

import "math"

type Rectangle struct {
	Width, Height float64
}

type Circle struct {
	Rdius float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Area() float64 {
	return c.Rdius * c.Rdius * math.Pi
}
