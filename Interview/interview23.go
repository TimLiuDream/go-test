package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		println(i)
		if i > 3 {
			goto loop
		}
	}

	loop:
	fmt.Print("end\n")
}
