package main

import (
	"fmt"
	"unsafe"
)

func main() {
	//说先定义一个切片，只限定长度为1
	s := make([]int, 1)

	//打印出slice的长度，容量以及内存地址
	fmt.Printf("len :%d cap:%d array ptr :%v \n", len(s), cap(s), *(*unsafe.Pointer)(unsafe.Pointer(&s)))

	for i := 0; i < 5; i++ {
		s = append(s, i)
		fmt.Printf("len :%d cap:%d array ptr :%v \n", len(s), cap(s), *(*unsafe.Pointer)(unsafe.Pointer(&s)))
	}
	//打印出slice
	fmt.Println("array:", s)
}
