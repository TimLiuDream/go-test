package main

import (
	"fmt"
	"time"
)

func main() {
	timeStr := "2050-05-17 10:00:00"
	loc, _ := time.LoadLocation("Local")
	the_time, err := time.ParseInLocation("2006-01-02 15:04:05", timeStr, loc)
	if err == nil {
		unix_time := the_time.Unix()
		fmt.Println(unix_time)
	}

	// dateStr := "2050-05-17"
	// loc, _ := time.LoadLocation("Local")
	// the_time, err := time.ParseInLocation("2006-01-02", dateStr, loc)
	// if err == nil {
	// 	unix_time := the_time.Unix()
	// 	fmt.Println(unix_time)
	// }

	// createTime := time.Now()
	// m, _ := time.ParseDuration("1m")
	// endTime := time.Now().Add(5 * m)

	// d := createTime.Sub(endTime)
	// fmt.Println(d.Seconds() > 0)
}
