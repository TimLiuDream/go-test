package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(foo())
	// c
	// nil
	// nil
}

func foo() (err error) {
	defer func() {
		fmt.Println(err)
		err = errors.New("a")
	}()
	defer func(e error) {
		fmt.Println(e)
		e = errors.New("b")
	}(err)
	return errors.New("c")
}
