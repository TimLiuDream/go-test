package main

import "fmt"

func main() {
	i, err := funcMui(1, 2)
	if err == nil {
		fmt.Print(i)
	}
}

func funcMui(x, y int) (sum int, err error) {
	return x + y, nil
}
