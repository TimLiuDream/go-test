package main

import (
	"fmt"
	"strconv"

	"github.com/siddontang/go-log/log"
)

func main() {
	parseBool()
	parseInt()
}

func parseBool() {
	isOnlyQueryTitle := "false1"
	result, _ := strconv.ParseBool(isOnlyQueryTitle)
	fmt.Println(result)
}

func parseInt() {
	str := ""
	v, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(v)
}
