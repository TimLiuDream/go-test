package main

import (
	"fmt"
	"time"
)

// 判断时间是否在某一分钟内
func main() {
	t := 1571310082
	fmt.Println(judgeTimeMinute(int64(t)))
}

func judgeTimeMinute(t int64) bool {
	now := time.Now()
	nowStamp := now.Unix()
	second := now.Second()
	nowStart := nowStamp - int64(second)
	nowEnd := nowStart + 59

	fmt.Println(t)
	fmt.Println(nowStart)
	fmt.Println(nowEnd)

	return nowStart <= int64(t) && int64(t) <= nowEnd
}
