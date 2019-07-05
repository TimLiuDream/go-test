package main

import (
	"fmt"
	"strconv"
)

func main() {
	isOnlyQueryTitle := "false1"
	result, _ := strconv.ParseBool(isOnlyQueryTitle)
	fmt.Println(result)
}
