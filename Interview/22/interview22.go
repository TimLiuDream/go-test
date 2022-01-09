package main

const cl = 100

var bl = 123

// 0x426180 123
// 100 100
func main() {
	println(&bl, bl)
	println(cl, cl)
}
