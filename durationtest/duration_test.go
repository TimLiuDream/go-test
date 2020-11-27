package main

import (
	"math"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDuration(t *testing.T) {
	// str := "49h8m2s"
	// d, err := time.ParseDuration(str)
	// if err != nil {
	// 	assert.Nil(t, err)
	// }
	// day, hour, minute := GetDayHourMinuteByDuration(d)
	// assert.Equal(t, int64(2), day)
	// assert.Equal(t, int64(1), hour)
	// assert.Equal(t, int64(8), minute)

	// var i float64 = 0.124
	// var e float64 = 0.12

	// var i float64 = 0.125
	// var e float64 = 0.13

	// var i float64 = 0.126
	// var e float64 = 0.13

	var i float64 = -0.124
	var e float64 = -0.12

	// var i float64 = -0.125
	// var e float64 = -0.12

	// var i float64 = -0.126
	// var e float64 = -0.13

	f := RoundFloat64(i, 2)
	assert.Equal(t, e, f)
}

func GetDayHourMinuteByDuration(d time.Duration) (day, hour, minute int64) {
	second := int64(d.Seconds())
	day = second / 60 / 60 / 24
	hour = (second - day*24*60*60) / 60 / 60
	minute = (second - day*24*60*60 - hour*60*60) / 60

	// rawDay := d.Hours() / 24
	// rawHour := d.Hours() - rawDay*24
	// rawMin := d.Minutes() - rawDay*24 - rawHour*60
	// day = int64(rawDay)
	// hour = int64(rawHour)
	// minute = int64(rawMin)
	return
}

// RoundFloat64 保留 precision 位小数，并进行四舍五入
func RoundFloat64(f float64, precision int) float64 {
	if precision < 0 || precision > 5 {
		precision = 5
	}
	var result float64
	if f >= 0 {
		f1 := f * math.Pow10(precision)
		ff := math.Floor(f1)
		if ff+0.5 > f1 {
			result = ff / math.Pow10(precision)
		} else {
			result = (ff + 1) / math.Pow10(precision)
		}
	} else {
		f1 := f * math.Pow10(precision)
		ff := math.Floor(f1)
		if ff+0.5 > f1 {
			result = ff / math.Pow10(precision)
		} else {
			result = (ff + 1) / math.Pow10(precision)
		}
	}
	return result
}
