package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func rune2bytes(rs []rune) []byte {
	return []byte(string(rs))
}

func bytes2rune(bs []byte) []rune {
	return []rune(string(bs))
}

func rune2string1(rs []rune) string {
	return string(rs)
}

func string2rune1(str string) []rune {
	rs := make([]rune, 0)
	for _, r := range str {
		rs = append(rs, r)
	}
	return rs
}

func string2rune2(str string) []rune {
	return []rune(str)
}

func bytes2string1(bs []byte) string {
	return string(bs)
}

func bytes2string2(bs []byte) string {
	return *(*string)(unsafe.Pointer(&bs))
}

func string2bytes1(str string) []byte {
	bs := make([]byte, 0)
	for i := 0; i < len(str); i++ {
		bs = append(bs, str[i])
	}
	return bs
}

func string2bytes2(str string) []byte {
	return []byte(str)
}

func string2bytes3(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}

func other() {
	str := "Hello, 中国!"

	rs := string2rune1(str)
	bs := string2bytes1(str)

	convertedBytes := rune2bytes(rs)
	convertedRunes := bytes2rune(bs)
	fmt.Println(bs)
	fmt.Println(convertedBytes)
	fmt.Println(rs)
	fmt.Println(convertedRunes)
}
