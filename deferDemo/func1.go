package main

import "fmt"

func func1() {
	Defer("John")
}

func Defer(name string) {
	defer func() {
		fmt.Printf("%s", name)
	}()

	defer func() {
		err := recover()
		if err != nil {
			fmt.Printf("%s", err)
		}
	}()

	name = "Lee"
	panic("error")
	defer func() {
		fmt.Printf("end")
	}()
}
