package main

import (
	"fmt"
	"unicode"

	"github.com/mozillazg/go-pinyin"
)

func IsChineseChar(str string) bool {
	for _, r := range str {
		if unicode.Is(unicode.Scripts["Han"], r) {
			return true
		}
	}
	return false
}

func main() {
	hans := "刘佛添"
	if !IsChineseChar(hans) {
		fmt.Println("不是汉字")
		return
	}
	fmt.Println(pinyin.LazyConvert(hans, nil))
}
