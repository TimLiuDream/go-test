package main

import "fmt"

type Person struct {
	ID   int
	Name string
	Age  int
}

func main() {
	var p = Person{ID: 1, Name: "Tony", Age: 25}
	fmt.Println(p)
	fmt.Printf("type of &l: %T\n", &p)
	fmt.Printf("type of  l: %T\n", p)
}

func (p *Person) String() string {
	return fmt.Sprintf("Person Type\nID: %d\nName: %s\nAge: %d", p.ID, p.Name, p.Age)
}
