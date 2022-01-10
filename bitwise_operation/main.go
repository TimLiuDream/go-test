package main

import (
	"fmt"
)

// | 操作符只能用于整形
func main() {
	// var a, b float64 = 1.0, 4.0 // 编译不过
	// var a, b int = 1, 4 // 5
	// var a, b string = "1", "4" // 编译不过
	// var a, b bool = true, false // 编译不过
	var a, b int8 = 1, 2
	fmt.Println(a | b)

}
