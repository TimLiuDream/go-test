package main

import (
	"fmt"
	"strings"

	"github.com/axgle/mahonia"
)

func main() {
	// func1()
	func5()
}

func func1() {
	s := "hello"
	s = "c" + s[1:]
	fmt.Printf("%s\n", s)
}

func func2() {
	str := "lkdjfl000"
	bytes := []rune(str)
	for from, to := 0, len(bytes)-1; from < to; from, to = from+1, to-1 {
		bytes[from], bytes[to] = bytes[to], bytes[from]
	}
	str = string(bytes)
	fmt.Print(str)
}

func func3() {
	var s *string
	s = new(string)
	fmt.Println(s)
}

func func4() {
	str := "乱码的字符串变量"
	str = ConvertToString(str, "gbk", "utf-8")
	fmt.Println(str)
}

func ConvertToString(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}

func func5() {
	str := strings.Replace("ssfsdfsdfsdfs", "dev", "single", -1)
	fmt.Println(str)
}
