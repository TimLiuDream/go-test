package models

import "fmt"

type Human struct {
	Name  string
	Age   int
	Phone string
}

type Student struct {
	Human
	School string
}

type Employee struct {
	Human
	Company string
}

func (h *Human) SayHi() {
	fmt.Printf("hi,i am %s you can call me on %s \n", h.Name, h.Phone)
}
