package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "23rojijf我993我8f"
	// tmp := "_副本"

	s = TruncateString(s, 10)
	fmt.Println(s)
}

func TruncateString(s string, maxRuneCount int) string {
	if utf8.RuneCountInString(s) <= maxRuneCount {
		return s
	}
	runes := []rune(s)[:maxRuneCount]
	return string(runes)
}
