package main

import (
	"fmt"
)

func main() {
	c := getClient()
	err := c.Set("key", "value")
	if err != nil {
		panic(err)
	}
	getVal, err := c.Get("key")
	if err != nil {
		panic(err)
	}
	fmt.Printf("get key: %s, value: %s\n", "key", getVal)
	getVal2, err := c.Get("key2")
	if err != nil {
		panic(err)
	}
	fmt.Printf("get key: %s, value: %s\n", "key", getVal2)
}
