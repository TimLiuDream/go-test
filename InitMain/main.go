package main

import (
	_ "GolangTraining/initlib1"
	_ "GolangTraining/initlib2"
	"fmt"
)

func init() {
	fmt.Println("libmain init")
}

func main() {
	fmt.Println("libmian main")
}
