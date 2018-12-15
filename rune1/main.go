package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var str = "hello 你好"
	//golang中的string底层是通过byte数组实现的，
	//所以直接求len实际就是在按字节长度计算
	//所以一个汉字占3个字节算3个长度
	fmt.Println("len(str)", len(str))

	//以下两种都可以得到str的字符串长度

	//golang中的unicode/utf-8包提供了用utf-8获取长度的方法
	fmt.Println("RuneCountInString:", utf8.RuneCountInString(str))

	//通过rune类型处理unicode字符
	fmt.Println("rune:", len([]rune(str)))
}
