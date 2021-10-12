package main

import (
	"fmt"
	"sync/atomic"
)

// 原子操作 减法
func main() {
	var (
		n uint64 = 97
		m uint64 = 1
		k int    = 2
	)
	const (
		a        = 3
		b uint64 = 4
		c uint32 = 5
		d int    = 6
	)

	show := fmt.Println
	atomic.AddUint64(&n, -m)
	show(n) // 96 (97 - 1)
	atomic.AddUint64(&n, -uint64(k))
	show(n) // 94 (95 - 2)
	atomic.AddUint64(&n, ^uint64(a-1))
	show(n) // 91 (94 - 3)
	atomic.AddUint64(&n, ^(b - 1))
	show(n) // 87 (91 - 4)
	atomic.AddUint64(&n, ^uint64(c-1))
	show(n) // 82 (87 - 5)
	atomic.AddUint64(&n, ^uint64(d-1))
	show(n) // 76 (82 - 6)
	x := b
	atomic.AddUint64(&n, -x)
	show(n) // 72 (76 - 4)
	atomic.AddUint64(&n, ^(m - 1))
	show(n) // 71 (72 - 1)
	atomic.AddUint64(&n, ^uint64(k-1))
	show(n) // 69 (71 - 2)
}
