package main

func main() {
	var i int8 = 1
	read(i)
}

//go:noinline
func read(i interface{}) {
	println(i)
}
