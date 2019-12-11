package main

import (
	"fmt"
	"strings"
)

// strings.TrimLeft：
// TrimLeft returns a slice of the string s with all leading Unicode code points contained in cutset removed.
// To remove a prefix, use TrimPrefix instead.

// strings.TrimPrefix
// TrimPrefix returns s without the provided leading prefix string. If s doesn't start with prefix, s is returned unchanged.

func main() {
	s1 := "field--3004.BWViJmUg.-priority"
	s2 := "field--3004.BWViJmUg.-priority"
	s1 = strings.TrimLeft(s1, "field-")
	fmt.Println(s1)
	s2 = strings.TrimPrefix(s2, "field-")
	fmt.Println(s2)
}
