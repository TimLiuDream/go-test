package main

import (
	"fmt"
	"errors"
)

func main() {
	fmt.Println("enter function main.")

	defer func() {
		fmt.Println("enter defer function.")
		if p := recover(); p != nil {
			fmt.Printf("panic:%s\n", p)
		}
		fmt.Println("exit defer function.")
	}()

	fmt.Printf("no panic:%v\n", recover())

	panic(errors.New("something wrong"))

	p := recover()
	fmt.Printf("panic:%s\n", p)

	fmt.Println("exit function main.")
}
