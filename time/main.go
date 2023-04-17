package main

import (
	"fmt"
	"time"
)

func main() {
	func2()
}

// TodayLastSecondTimestampFromTimestamp 从时间戳获取当天最后一秒, local时区
func TodayLastSecondTimestampFromTimestamp(timestamp int64) int64 {
	if timestamp == -1 {
		return -1
	}
	str := TimeInt64ToString(timestamp)
	t := TimeStringToInt64(str, time.Local)
	return t + 24*60*60 - 1
}

// TimeInt64ToString 时间戳转字符串
func TimeInt64ToString(i int64) string {
	tm := time.Unix(i, 0)
	return tm.Format("2006-01-02")
}

// TimeStringToInt64 时间字符串转时间戳
func TimeStringToInt64(s string, loc *time.Location) int64 {
	tm, _ := time.ParseInLocation("2006-01-02", s, loc)
	return tm.Unix()
}

func func1() {
	fmt.Println(time.Now().In(time.UTC).Unix())
	fmt.Println(time.Now().In(time.Local).Unix())
}

func func2() {
	t1 := time.Now()
	fmt.Println(t1)
	time.Sleep(time.Second)
	t2 := time.Now()
	fmt.Println(time.Now())
	fmt.Println(t2.Sub(t1))
}

func func3() {
	timeStr := "2050-05-17 10:00:00"
	loc, _ := time.LoadLocation("Local")
	the_time, err := time.ParseInLocation("2006-01-02 15:04:05", timeStr, loc)
	if err == nil {
		unix_time := the_time.Unix()
		fmt.Println(unix_time)
	}
}

func func4() {
	dateStr := "2050-05-17"
	loc, _ := time.LoadLocation("Local")
	the_time, err := time.ParseInLocation("2006-01-02", dateStr, loc)
	if err == nil {
		unix_time := the_time.Unix()
		fmt.Println(unix_time)
	}

	createTime := time.Now()
	m, _ := time.ParseDuration("1m")
	endTime := time.Now().Add(5 * m)

	d := createTime.Sub(endTime)
	fmt.Println(d.Seconds() > 0)
}
