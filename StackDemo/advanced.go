package main

import (
	"fmt"
)

//栈的特点是先进后出
//使用一个切片的全局变量来模拟栈
var stack1 []int

//此通道用于通知主协程已经完成操作了
//但是此操作有可能不会输出全部数据
//因为添加数据和获取数据是异步的
//当获取数据的速度快于写入数据
//便不会输出全部数据
var e chan int = make(chan int)

//向栈中添加数据
func push1(value int) {
	stack1 = append(stack1, value)
	fmt.Println(stack1)
}

//从栈中获取数据
func pop1() {
	for {
		if len(stack1) > 0 {
			value := stack1[len(stack1)-1]
			stack1 = stack1[:len(stack1)-1]
			fmt.Println(value)
		} else {
			e <- 0
		}
	}
}

func main() {
	for i := 0; i < 10; i++ {
		go push1(i)
	}

	go pop1()

	<-e
}
