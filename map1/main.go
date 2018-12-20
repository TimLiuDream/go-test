package main

import "fmt"

func changeMapItem(m map[string]string, key string) bool {
	if key == "hello" {
		m[key] = "hello2345678987654"
		return true
	}
	return false
}

func main() {
	m := map[string]string{"hello": "hello", "world": "world"}
	changeMapItem(m, "hello")
	fmt.Println(m)
}
