package main

import "fmt"

type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human
	school string
	loan   float32
}

type Employee struct {
	Human
	company string
	money   float32
}

func (h Human) SayHi() {
	fmt.Printf("Hi i am %s you can call me on %s\n", h.name, h.phone)
}

func (h Human) Sing(lyrics string) {
	fmt.Println("la la la ...", lyrics)
}

func (e Employee) SayHi() {
	fmt.Printf("hi i am %s i work at %s call me on %s\n", e.name, e.company, e.phone)
}

type Men interface {
	SayHi()
	Sing(lyrics string)
}

func main() {
	mike := Student{Human{"Mike", 25, "12345"}, "MIT", 0.00}
	//paul := Student{Human{"Paul", 26, "123456"}, "Harvard", 1.00}
	//sam := Employee{Human{"Sam", 36, "12458098765"}, "Golang Inc", 1000}
	tom := Employee{Human{"Tom", 37, "34567876543"}, "Things Ltd", 5000}

	var i Men

	i = mike
	fmt.Println("This is mike a student:")
	i.SayHi()
	i.Sing("November rain")

	i = tom
	fmt.Println("this is tom an employee:")
	i.SayHi()
	i.Sing("born to be wild")

	fmt.Println("let's use a slice of men and see what happens")
	//x := make([]Men, 3)
	//
	//x[0], x[1], x[2] = paul, sam, mike
	//
	//for _, value := range x {
	//	value.SayHi()
	//}
	var x Men = Student{Human{"Mike", 25, "12345"}, "MIT", 0.00}
	x.SayHi()
}
