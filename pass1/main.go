package main

import "fmt"

func main() {
	s := "脑子进煎鱼了"
	fmt.Printf("main 内存地址：%p\n", &s)
	hello(&s)
	fmt.Println(s)
}

func hello(s *string) {
	fmt.Printf("hello 内存地址：%p\n", &s)
	*s = "煎鱼进脑子了"
}
