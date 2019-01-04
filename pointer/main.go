package main

import (
	"fmt"
	"unsafe"
)

func main() {
	num:=5
	numPointer:=&num

	//flnum:=(*float32)(numPointer)//不同的类型不能够进行赋值、计算等跨类型的操作
	flnum:=(*float32)(unsafe.Pointer(numPointer))//先把numPointer转换为Pointer，再转换为*float32
	fmt.Println(flnum)
}
