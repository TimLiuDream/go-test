package main

import (
	"fmt"
	"time"
)

const (
	DayMinutes = 24 * HourMinute
	HourMinute = 60
)

func main() {
	d, h, m := culTime(1573056000)
	fmt.Println(d)
	fmt.Println(h)
	fmt.Println(m)
	d, h, m = culTime(1573228800)
	fmt.Println(d)
	fmt.Println(h)
	fmt.Println(m)
}

func culTime(targetTime int64) (d, h, m int) {
	tm := time.Unix(targetTime, 0)
	nowStamp := time.Now().Unix()
	var duration time.Duration
	if targetTime >= nowStamp {
		duration = time.Until(tm)
	} else {
		duration = time.Since(tm)
	}
	minutes := duration.Minutes()
	rawMinutes := int(minutes)
	if rawMinutes >= DayMinutes {
		d = rawMinutes / DayMinutes
		h = (rawMinutes - d*DayMinutes) / HourMinute
		m = (rawMinutes - d*DayMinutes) - h*HourMinute
	} else {
		h = rawMinutes / HourMinute
		m = (rawMinutes - h*HourMinute)
	}

	return
}
