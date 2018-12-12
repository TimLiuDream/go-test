package main

import (
	"fmt"
	"time"
)

func main() {
	createTime := time.Now()
	m, _ := time.ParseDuration("1m")
	endTime := time.Now().Add(5 * m)

	d := createTime.Sub(endTime)
	fmt.Println(d.Seconds() > 0)
}
