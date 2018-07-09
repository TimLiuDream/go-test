package main

import "fmt"

func main() {
	slice := []byte{'a', 'b', 'c', 'd'}
	fmt.Println(slice[3])

	var ar = [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
	var a, b []byte
	a = ar[2:5]
	b = ar[3:5]
	fmt.Println(a[0])
	fmt.Println(b[0])
}
