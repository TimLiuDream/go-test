package main

import (
	"fmt"

	"github.com/timliudream/golangtraining/method/models"
)

func main() {
	func1()
}

func func1() {
	r1 := models.Rectangle{12, 2}
	r2 := models.Rectangle{9, 4}
	c1 := models.Circle{5}
	c2 := models.Circle{10}
	fmt.Println("Area of r1 is ", r1.Area())
	fmt.Println("Area of r2 is ", r2.Area())
	fmt.Println("Area of c1 is ", c1.Area())
	fmt.Println("Area of c2 is ", c2.Area())
}

func func2() {
	mark := models.Student{models.Human{"Mark", 25, "15088131357"}, "MIT"}
	sam := models.Employee{models.Human{"Sam", 45, "15088131358"}, "GOlang inc"}

	mark.SayHi()
	sam.SayHi()
}

func func3() {
	boxes := models.BoxList{
		models.Box{4, 4, 4, models.RED},
		models.Box{10, 10, 1, models.YELLOW},
		models.Box{1, 1, 20, models.BLACK},
		models.Box{10, 10, 1, models.BLUE},
		models.Box{10, 30, 1, models.WHITE},
		models.Box{20, 20, 20, models.YELLOW},
	}

	fmt.Printf("We have %d boxes in our set \n", len(boxes))
	fmt.Println("The volunme of the first one is ", boxes[0].Volume(), "cm^3")
	fmt.Println("The color of the last one is ", boxes[len(boxes)-1].Color.String())
	fmt.Println("The biggest one is ", boxes.BiggestColor().String())

	fmt.Println("Let's paint them all black")
	boxes.PaintItBlack()
	fmt.Println("THe color of the second one is ", boxes[1].Color.String())
	fmt.Println("Obviously, now the biggest one is ", boxes.BiggestColor().String())
}
