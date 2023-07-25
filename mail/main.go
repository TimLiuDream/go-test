package main

import (
	"fmt"
	"net/mail"
)

func main() {
	ss := "～·！@#¥$%……^&*()_-—+=【】「」{}[]\\|、：;’”“‘，。.《》？<>/±＋－×÷:\""
	part := " <liufotian@gmail.com>"
	for _, r := range ss {
		s := string(r) + part
		_, err := mail.ParseAddress(s)
		if err != nil {
			fmt.Printf("add: %s has error: %s\n", s, err.Error())
		}
	}
}
