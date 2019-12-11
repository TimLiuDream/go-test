package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	t := "09:10"
	if VerifyHourMinute(t) {
		ts := strings.Split(t, ":")
		hour, _ := strconv.Atoi(ts[0])
		minute, _ := strconv.Atoi(ts[1])
		fmt.Println(hour)
		fmt.Println(minute)
	}
}

func VerifyHourMinute(t string) bool {
	pattern := "((1|0?)[0-9]|2[0-3]):([0-5][0-9])"
	reg := regexp.MustCompile(pattern)
	result := reg.FindAllString(t, -1)
	return len(result) >= 1
}
