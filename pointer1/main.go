package main

import (
	"fmt"
	"unsafe"
)

type Num struct {
	i string
	j int64
}

func main() {
	n := Num{i: "timliu", j: 1}
	nPointer := unsafe.Pointer(&n)

	niPointer := (*string)(unsafe.Pointer(nPointer))
	*niPointer = "咸鱼"

	//基于结构体的成员地址去计算偏移量。就能够得出其他成员变量的内存地址
	njPointer:=(*int64)(unsafe.Pointer(uintptr(nPointer)+unsafe.Offsetof(n.j)))
	*njPointer=2

	fmt.Printf("n.i:%s;n.j:%d",n.i,n.j)
}
