package main

import "fmt"

func main() {
	func1()
}

func func1() {
	var err error
	for i := 0; i < 100; i++ {
		defer func() {
			if err != nil {
				fmt.Println("error hanppen")
			}
		}()
		fmt.Println(i)
		if i == 50 {
			err = fmt.Errorf("ones!")
			continue
		}
	}
}
