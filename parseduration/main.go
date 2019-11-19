package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)
	hour, err := time.ParseDuration(fmt.Sprintf("%dh", 0))
	if err != nil {
		fmt.Println(err)
	}
	minute, err := time.ParseDuration(fmt.Sprintf("%dm", 0))
	if err != nil {
		fmt.Println(err)
	}
	d := now.AddDate(0, 0, 1)
	d = d.Add(hour)
	d = d.Add(minute)
	fmt.Println(d)
}
