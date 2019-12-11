package main

import (
	"fmt"
	"unsafe"

	"github.com/timliudream/golangtraining/pointer/models"
)

func main() {

}

func func1() {
	num := 5
	numPointer := &num

	//flnum:=(*float32)(numPointer)//不同的类型不能够进行赋值、计算等跨类型的操作
	flnum := (*float32)(unsafe.Pointer(numPointer)) //先把numPointer转换为Pointer，再转换为*float32
	fmt.Println(flnum)
}

func func2() {
	n := models.Num{I: "timliu", J: 1}
	nPointer := unsafe.Pointer(&n)

	niPointer := (*string)(unsafe.Pointer(nPointer))
	*niPointer = "咸鱼"

	//基于结构体的成员地址去计算偏移量。就能够得出其他成员变量的内存地址
	njPointer := (*int64)(unsafe.Pointer(uintptr(nPointer) + unsafe.Offsetof(n.J)))
	*njPointer = 2

	fmt.Printf("n.i:%s;n.j:%d", n.I, n.J)
}
