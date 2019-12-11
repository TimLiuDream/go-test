package models

type Person struct {
	Name string
	Age  int
}

type Human struct {
	Name   string
	Age    int
	Weight int
}

type Student struct {
	Human
	Skills     []string
	Int        int
	Speciality string
}

func Older(p1, p2 Person) (Person, int) {
	if p1.Age > p2.Age {
		return p1, p1.Age - p2.Age
	}
	return p2, p2.Age - p1.Age
}
