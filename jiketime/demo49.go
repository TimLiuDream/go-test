package main

import (
	"fmt"
	"errors"
)

func main() {
	fmt.Println("enter function main.")
	caller()
	fmt.Println("exit function main.")
}

func caller() {
	fmt.Println("enter function caller")
	panic(errors.New("something wrong"))
	panic(fmt.Println)
	fmt.Println("exit function caller")
}
