package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 6}
	cpes := make([]int, len(s))
	cphs := make([]int, len(s))
	copy(cpes, s)
	copy(cphs, s)
	es := cpes[4:]
	fmt.Println(es)
	hs := cphs[:4]
	fmt.Println(hs)
	hs = append(hs, []int{111}...)
	fmt.Println(hs)
	s = append(hs, es...)
	fmt.Println(s)
}
