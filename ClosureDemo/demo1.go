package main

import "fmt"

func main() {
	var flist []func()
	for i := 0; i < 3; i++ {

		i := i //给i变量重新赋值，
		fmt.Println(i)
		flist = append(flist, func() {
			fmt.Println(i)
		})
	}
	for _, f := range flist {
		f()
	}
}
