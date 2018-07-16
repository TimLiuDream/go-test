package main

import "fmt"

type Human struct{
	name string
	age int
	phone string
}

type Student struct{
	Human
	school string
}

type Employee struct{
	Human
	company string
}

func (h *Human)SayHi()  {
	fmt.Printf("hi,i am %s you can call me on %s \n",h.name,h.phone)
}

func main()  {
	mark :=Student{Human{"Mark",25,"15088131357"},"MIT"}
	sam:=Employee{Human{"Sam",45,"15088131358"},"GOlang inc"}

	mark.SayHi()
	sam.SayHi()
}