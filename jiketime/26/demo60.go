package main

import (
	"sync"
	"time"
	"log"
)

//counter代表计数器
type counter struct {
	num uint         //计数
	mu  sync.RWMutex //读写锁
}

//number 会返回当前的计数
func (c *counter) number() uint {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.num
}

//add会增加计数器的值，并会返回增加后的计数
func (c *counter) add(increment uint) uint {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.num += increment
	return c.num
}

func main() {
	c := counter{}
	count(&c)
	redundantUnlock()
}

func count(c *counter) {
	//sign用于传递演示完成的信息
	sign := make(chan struct{}, 3)
	go func() {
		defer func() { sign <- struct{}{} }()
		for i := 1; i <= 10; i++ {
			time.Sleep(time.Millisecond * 500)
			c.add(1)
		}
	}()
	go func() {
		defer func() { sign <- struct{}{} }()
		for j := 1; j <= 20; j++ {
			time.Sleep(time.Millisecond * 200)
			log.Printf("the number in counter :%d [%d-%d]", c.number(), 1, j)
		}
	}()
	go func() {
		defer func() { sign <- struct{}{} }()
		for k := 1; k <= 20; k++ {
			time.Sleep(time.Millisecond * 300)
			log.Printf("the number in counter:%d [%d-%d]", c.number(), 2, k)
		}
	}()
	<-sign
	<-sign
	<-sign
}

func redundantUnlock() {
	var rwMu sync.RWMutex

	//ex1
	//rwMu.Unlock()

	//ex2
	rwMu.RUnlock()

	//ex3
	//rwMu.RLock()
	//rwMu.Unlock()
	//rwMu.RUnlock()

	//ex4
	//rwMu.Lock()
	//rwMu.RUnlock()
	//rwMu.Unlock()
}