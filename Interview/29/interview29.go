package main

func test(x int) (func(), func()) {
	return func() {
			println(x)
			x += 10
		}, func() {
			println(x)
		}
}

// 100
// 110

func main() {
	a, b := test(100)
	a()
	b()
}
