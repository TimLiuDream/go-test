package main

import (
	"fmt"
	"regexp"
)

func IsChineseChar(str string) bool {
	var result bool
	result = true
	if m, _ := regexp.MatchString("^\\p{Han}+$", str); !m {
		result = false
	}
	return result
}

func main() {
	s := "我是中国人"
	fmt.Println(IsChineseChar(s))
}
