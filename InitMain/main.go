package main

import (
	"fmt"
	_ "GolangTraining/InitLib1"
	_ "GolangTraining/InitLib2"
)

func init() {
	fmt.Println("libmain init")
}

func main() {
	fmt.Println("libmian main")
}
