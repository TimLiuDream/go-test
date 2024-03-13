package main

import "sync"

var i = 0
var _isInitialized = false
var once = sync.Once{}

func CreateInstance() {
	once.Do(func() {
		i++
	})
}
