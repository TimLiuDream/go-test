package main

import "strconv"

func isPalindrome(x int) bool {
	// write code here
	s := strconv.Itoa(x)
	i, j := 0, len(s)-1
	for i <= j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
