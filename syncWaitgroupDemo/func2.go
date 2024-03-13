package main

import (
	"fmt"
	"sync"
)

var a int

func Add() {
	a++
}

func func2() {
	wg := sync.WaitGroup{}
	for i := 0; i < 500; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			Add()
		}()
	}
	wg.Wait()
	fmt.Println(a)
}
