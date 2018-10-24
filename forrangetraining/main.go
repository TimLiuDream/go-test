package main

func ForSlice(s []string) {
	len := len(s)
	for i := 0; i < len; i++ {
		_, _ = i, s[i]
	}
}

func RangeForSlice(s []string) {
	for i, v := range s {
		_, _ = i, v
	}
}

func main() {
	s:=make([]string,5)
	ForSlice(s)
}
