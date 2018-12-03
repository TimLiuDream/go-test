package main

import (
	"fmt"
)

func main() {
	str := "Go爱好者"
	fmt.Printf("The string:%q\n", str)
	fmt.Printf("  => runes(char):%q\n", []rune(str))
	fmt.Printf("  => runes(hex):%x\n", []rune(str))
	fmt.Printf("  => bytes(hex):[% x]\n", []byte(str)) //因为一个中文字符的 UTF-8 编码值需要用三个字节来表达。
}
