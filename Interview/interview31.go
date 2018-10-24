package main

import (
	"fmt"
	"time"
)

func print(pi *int) {
	fmt.Println("d", pi)
}

func main() {
	for i := 0; i < 10; i++ {
		//defer fmt.Println("a", i)              //倒序依次调用
		//defer func() { fmt.Println("b", i) }() //指向i这个变量，i经过循环之后的值为10
		//defer func(i int) { fmt.Println("c", i) }(i)//倒序依次调用
		//defer print(&i)//输出i的内存地址
		//go fmt.Println("e", i)//0到9随机输出
		//go func() {fmt.Println("f",i)}()//0到9随机输出
		//go func(i int) {fmt.Println("g",i)}(i)//0到9随机输出
	}

	//fmt.Println("h")
	//defer fmt.Println("i")
	time.Sleep(3 * time.Second)
}
