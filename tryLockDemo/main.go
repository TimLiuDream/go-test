package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Mutex struct {
	state int32
}

func (m *Mutex) TryLock() bool {
	return atomic.CompareAndSwapInt32(&m.state, 0, 1)
}

func (m *Mutex) Unlock() {
	atomic.StoreInt32(&m.state, 0)
}

func main() {
	var wg sync.WaitGroup
	m := &Mutex{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		if m.TryLock() {
			defer m.Unlock()
			fmt.Println("lock success!----1")
		} else {
			fmt.Println("lock failed!----1")
		}
	}()

	go func() {
		defer wg.Done()
		if m.TryLock() {
			defer m.Unlock()
			fmt.Println("lock success!----2")
		} else {
			fmt.Println("lock failed!----2")
		}
	}()

	wg.Wait()
}
