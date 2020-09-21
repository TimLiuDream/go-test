package main

import "fmt"

const (
	x = iota
	y
	z = "zz"
	k
	j
	p = iota
)

func main()  {
	fmt.Println(x,y,z,k,j,p)
}
