package main

import (
	"fmt"
	"sync"
)

//下面的迭代会有什么问题？

type threadSafeSet struct {
	sync.Map
	s []interface{}
}

func (set *threadSafeSet) Iter() <-chan interface{} {
	// 解除注释看看！
	//Iter: 0 (0x5a04e0,0x5c8ca8)
	// ch := make(chan interface{})

	ch := make(chan interface{}, len(set.s))
	// 	Iter: 0 (0x8504e0,0x878ca8)
	// Iter: 1 (0x8504e0,0x878cb8)

	go func() {
		for elem, value := range set.s {
			ch <- elem
			println("Iter:", elem, value)
		}

		close(ch)
	}()
	return ch
}

func main() {
	th := threadSafeSet{
		s: []interface{}{"1", "2"},
	}
	v := <-th.Iter()
	fmt.Sprintf("%s%v", "ch", v)
}
