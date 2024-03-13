package main

import (
	"fmt"
	"sync"
	"time"
)

func read(str string, c *sync.Cond) {
	c.L.Lock()
	defer c.L.Unlock()

	c.Wait()
	fmt.Println(str, "start reading")
}

func write(str string, c *sync.Cond) {
	fmt.Println(str, "start writing")
	time.Sleep(2 * time.Second)
	c.L.Lock()
	// do something
	c.L.Unlock()
	fmt.Println(str, "weak up all")
	c.Broadcast()
}

func func3() {
	m := &sync.Mutex{}
	c := sync.NewCond(m)

	go read("reader1", c)
	go read("reader2", c)
	write("writer", c)

	time.Sleep(5 * time.Second)
}
