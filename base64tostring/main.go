package main

import (
	"encoding/base64"
	"fmt"
	"strings"
)

func main() {
	base64str := "MTExMTEgMjIyMjIgMzQ0NTM0NQo="
	data, err := base64.StdEncoding.DecodeString(base64str)
	if err != nil {
		panic(err)
	}
	str := string(data)
	fmt.Println(str)
	strs := strings.Split(str, " ")
	for _, s := range strs {
		fmt.Println(s)
	}
}
