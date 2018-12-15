package main

import "fmt"

//栈的特点是先进后出
//使用一个切片的全局变量来模拟栈
var stack []int

//向栈中添加数据
func push(value int) {
	stack = append(stack, value)
}

//从栈中获取数据
func pop() (int, bool) {
	ok := false
	value := 0
	if len(stack) > 0 {
		value = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ok = true
		return value, ok
	} else {
		return value, ok
	}
}

func main() {
	//向栈中添加数据
	for i := 0; i < 10; i++ {
		push(i)
		fmt.Println(stack)
	}
	//从栈中获取数据
	for {
		v, ok := pop()
		if ok {
			fmt.Println(v)
		} else {
			break
		}
	}
}
