package main

import (
	"fmt"
	"sync/atomic"
	"unsafe"
)

type T struct{ x int }

var pT *T

func main() {
	var unsafePPT = (*unsafe.Pointer)(unsafe.Pointer(&pT))
	var ta, tb = T{1}, T{2}
	// 修改
	atomic.StorePointer(
		unsafePPT, unsafe.Pointer(&ta))
	fmt.Println(pT) // &{1}
	// 读取
	pa1 := (*T)(atomic.LoadPointer(unsafePPT))
	fmt.Println(pa1 == &ta) // true
	// 置换
	pa2 := atomic.SwapPointer(
		unsafePPT, unsafe.Pointer(&tb))
	fmt.Println((*T)(pa2) == &ta) // true
	fmt.Println(pT)               // &{2}
	// 比较置换
	b := atomic.CompareAndSwapPointer(
		unsafePPT, pa2, unsafe.Pointer(&tb))
	fmt.Println(b) // false
	b = atomic.CompareAndSwapPointer(
		unsafePPT, unsafe.Pointer(&tb), pa2)
	fmt.Println(b) // true
}
