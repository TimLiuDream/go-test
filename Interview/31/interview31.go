package main

import (
	"fmt"
	"runtime"
	"time"
)

func print(pi *int) {
	fmt.Println("d", pi)
}

func main() {
	for i := 0; i < 10; i++ {
		//当defer被声明时，其参数就会被实时解析，而且defer是栈结构，所以是0-9倒序依次调用
		//对于这个普通的defer
		//defer fmt.Println("a", i)

		//闭包，闭包里引用了不作为参数传递进去的值,都是引用传递
		//指向i这个变量，i经过循环之后的值为10
		//defer func() { fmt.Println("b", i) }()

		//闭包，闭包里引用了不作为参数传递进去的值,都是引用传递,所以这里是值传递
		//倒序依次调用
		//defer func(i int) { fmt.Println("c", i) }(i)

		//输出i的内存地址
		//defer print(&i)

		//for循环开启10条goroutine
		//0到9随机输出
		//go fmt.Println("e", i)

		//闭包，闭包里引用了不作为参数传递进去的值,都是引用传递
		//但是开了10条goroutine，所以这10个goroutine会不断抢占去访问i的值
		//0到10随机输出，但是说不定会输出哪个值，10出现最多次
		//协程执行时机问题,除非设置单核执行，那就是全部输出10
		runtime.GOMAXPROCS(1)
		fmt.Println("p", i)
		go func() { fmt.Println("f", i) }()

		////闭包，闭包里引用了不作为参数传递进去的值,都是引用传递,所以这里是值传递
		//0到9随机输出,但是0到9各出现一次
		//go func(i int) {fmt.Println("g",i)}(i)
	}

	//fmt.Println("h")
	//defer fmt.Println("i")
	time.Sleep(3 * time.Second)
}
