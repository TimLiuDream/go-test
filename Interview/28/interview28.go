package main

func test() []func() {
	var funs []func()
	for i := 0; i < 2; i++ {
		x := i
		funs = append(funs, func() {
			println(&x, x)
		})
	}
	return funs
}

// 0xc000014030 0
// 0xc000014038 1
func main() {
	funs := test()
	for _, f := range funs {
		f()
	}
}
