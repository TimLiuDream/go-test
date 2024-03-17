package main

//func main() {
//	x := 2
//	square(x)
//}
//
//func square(x int) int {
//	return x * x
//}
//# github.com/timliudream/go-test/EscapeDemo
//./main.go:8:6: can inline square
//./main.go:3:6: can inline main
//./main.go:5:8: inlining call to square

//func main() {
//	x := 2
//	square(x)
//}
//
//func square(x int) *int {
//	y := x * x
//	return &y
//}
//# github.com/timliudream/go-test/EscapeDemo
//./main.go:21:6: can inline square
//./main.go:16:6: can inline main
//./main.go:18:8: inlining call to square
//./main.go:22:2: moved to heap: y

//func main() {
//	x := 4
//	square(&x)
//}
//
//func square(x *int) {
//	*x = *x * *x
//}
//// # github.com/timliudream/go-test/EscapeDemo
////./main.go:36:6: can inline square
////./main.go:31:6: can inline main
////./main.go:33:8: inlining call to square
////./main.go:36:13: x does not escape

func main() {
	x := 4
	square(&x)
}

func square(x *int) *int {
	y := *x * *x
	return &y
}

// # github.com/timliudream/go-test/EscapeDemo
//./main.go:50:6: can inline square
//./main.go:45:6: can inline main
//./main.go:47:8: inlining call to square
//./main.go:50:13: x does not escape
//./main.go:51:2: moved to heap: y
