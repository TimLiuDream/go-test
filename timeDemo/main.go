package main

import (
	"fmt"
	"time"
)

func StartOfMonth(date time.Time) time.Time {
	return time.Date(date.Year(), date.Month(), 1, 0, 0, 0, 0, date.Location())
}

func EndOfMonth(date time.Time) time.Time {
	firstDayOfNextMonth := StartOfMonth(date).AddDate(0, 1, 0)
	return firstDayOfNextMonth.Add(-time.Second)
}

func StartOfDayOfWeek(date time.Time) time.Time {
	daysSinceSunday := int(date.Weekday())
	return date.AddDate(0, 0, -daysSinceSunday+1)
}

func EndOfDayOfWeek(date time.Time) time.Time {
	daysUntilSaturday := 7 - int(date.Weekday())
	return date.AddDate(0, 0, daysUntilSaturday)
}

func StartAndEndOfWeeksOfMonth(year, month int) []struct{ Start, End time.Time } {
	startOfMonth := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	weeks := make([]struct{ Start, End time.Time }, 0)

	for current := startOfMonth; current.Month() == time.Month(month); current = current.AddDate(0, 0, 7) {
		startOfWeek := StartOfDayOfWeek(current)
		endOfWeek := EndOfDayOfWeek(current)

		if endOfWeek.Month() != time.Month(month) {
			endOfWeek = EndOfMonth(current)
		}
		weeks = append(weeks, struct{ Start, End time.Time }{startOfWeek, endOfWeek})
	}

	return weeks
}

func WeekNumberInMonth(date time.Time) int {
	startOfMonth := StartOfMonth(date)
	_, week := date.ISOWeek()
	_, startWeek := startOfMonth.ISOWeek()
	return week - startWeek + 1
}

func StartOfYear(date time.Time) time.Time {
	return time.Date(date.Year(), time.January, 1, 0, 0, 0, 0, date.Location())
}

func EndOfYear(date time.Time) time.Time {
	startOfNextYear := StartOfYear(date).AddDate(1, 0, 0)
	return startOfNextYear.Add(-time.Second)
}

func StartOfQuarter(date time.Time) time.Time {
	// you can directly use 0, 1, 2, 3 quarter
	quarter := (int(date.Month()) - 1) / 3
	startMonth := time.Month(quarter*3 + 1)
	return time.Date(date.Year(), startMonth, 1, 0, 0, 0, 0, date.Location())
}

func EndOfQuarter(date time.Time) time.Time {
	startOfNextQuarter := StartOfQuarter(date).AddDate(0, 3, 0)
	return startOfNextQuarter.Add(-time.Second)
}

func CurrentWeekRange(timeZone string) (startOfWeek, endOfWeek time.Time) {
	loc, _ := time.LoadLocation(timeZone)

	now := time.Now().In(loc)
	startOfWeek = StartOfDayOfWeek(now)
	endOfWeek = EndOfDayOfWeek(now)

	return startOfWeek, endOfWeek
}

func DurationBetween(start, end time.Time) time.Duration {
	return end.Sub(start)
}

func GetDatesForDayOfWeek(year, month int, day time.Weekday) []time.Time {
	var dates []time.Time

	firstDayOfMonth := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	diff := int(day) - int(firstDayOfMonth.Weekday())
	if diff < 0 {
		diff += 7
	}

	firstDay := firstDayOfMonth.AddDate(0, 0, diff)
	for current := firstDay; current.Month() == time.Month(month); current = current.AddDate(0, 0, 7) {
		dates = append(dates, current)
	}

	return dates
}

func AddBusinessDays(startDate time.Time, daysToAdd int) time.Time {
	currentDate := startDate
	for i := 0; i < daysToAdd; {
		currentDate = currentDate.AddDate(0, 0, 1)
		if currentDate.Weekday() != time.Saturday && currentDate.Weekday() != time.Sunday {
			i++
		}
	}
	return currentDate
}

func FormatDuration(duration time.Duration) string {
	days := int(duration.Hours() / 24)
	hours := int(duration.Hours()) % 24
	minutes := int(duration.Minutes()) % 60
	seconds := int(duration.Seconds()) % 60
	return fmt.Sprintf("%d天 %02d小时 %02d分 %02d秒", days, hours, minutes, seconds)
}

func main() {
	time1 := time.Now()
	time.Sleep(50 * time.Millisecond)
	fmt.Printf("cost: %v", time.Now().Sub(time1))
	//fmt.Println(FormatDuration(time.Hour*24*3 + time.Hour*4 + time.Minute*15 + time.Second*30))
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

func IsSameHour(t1, t2 int64) bool {
	time1 := time.Unix(t1, 0)
	time2 := time.Unix(t2, 0)
	y1, m1, d1 := time1.Date()
	h1 := time1.Hour()

	y2, m2, d2 := time2.Date()
	h2 := time2.Hour()
	return y1 == y2 && m1 == m2 && d1 == d2 && h1 == h2
}
