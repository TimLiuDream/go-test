package main

import (
	"fmt"
	"time"
)

// 获取明天的起止时间戳
func main() {
	start, end := getTomorrowStartEndStamp()
	fmt.Println(start)
	fmt.Println(end)
}

func getTomorrowStartEndStamp() (start int64, end int64) {
	tomorrow := time.Now().AddDate(0, 0, 1)
	tomorrowStart := time.Date(tomorrow.Year(), tomorrow.Month(), tomorrow.Day(), 0, 0, 0, 0, tomorrow.Location())
	tomorrowEnd := time.Date(tomorrow.Year(), tomorrow.Month(), tomorrow.Day(), 23, 59, 59, 59, tomorrow.Location())
	start = tomorrowStart.Unix()
	end = tomorrowEnd.Unix()
	return
}
