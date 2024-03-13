package main

import (
	"fmt"
	"sync"
	"time"
)

var m = sync.Map{}

func Set(key string, value interface{}) {
	m.Store(key, value)
}

func Get(key string) interface{} {
	v, ok := m.Load(key)
	if !ok {
		return nil
	}
	fmt.Printf("Get value: %d\n", v)
	return v
}

func main() {
	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("%d", i)
		value := i
		go Set(key, value)
		go Get(key)
	}
	time.Sleep(time.Second * 5)
}
