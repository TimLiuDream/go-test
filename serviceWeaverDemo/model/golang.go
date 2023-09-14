package model

import "github.com/ServiceWeaver/weaver"

type Golang struct {
	weaver.AutoMarshal
	Channel   string
	Goroutine string
}
