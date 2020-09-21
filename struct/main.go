package main

import (
	"fmt"

	"github.com/timliudream/go-test/struct/models"
)

func main() {
	func1()
}

func func1() {
	var tom models.Person
	tom.Name, tom.Age = "Tom", 18
	bob := models.Person{"bob", 20}
	tb_Older, tb_diff := models.Older(tom, bob)
	fmt.Printf("of %s and %s,%s is older by %d years \n", tom.Name, bob.Name, tb_Older.Name, tb_diff)
}

func func2() {
	jane := models.Student{Human: models.Human{"Jane", 35, 100}, Speciality: "Biology"}
	fmt.Println("Her name is ", jane.Name)
	fmt.Println("Her age is ", jane.Age)
	fmt.Println("Her weight is ", jane.Weight)
	fmt.Println("Her speciality is ", jane.Speciality)

	jane.Skills = []string{"anatomy"}
	fmt.Println("Her skills are ", jane.Skills)
	fmt.Println("She acquired two new ones ")
	jane.Skills = append(jane.Skills, "physics", "golang")
	fmt.Println("Her skills now are ", jane.Skills)

	jane.Int = 3
	fmt.Println("Her preferred number is ", jane.Int)
}
