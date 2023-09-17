package main

import "log"

func main() {
	c, err := NewCache("")
	if err != nil {
		log.Fatalln(err)
	}
	c.Set("hello", "world")
	c.Get("hello")
}
