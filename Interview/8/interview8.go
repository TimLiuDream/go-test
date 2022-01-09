package main

import (
	"fmt"
	"sync"
)

type UserAges struct {
	ages map[string]int
	sync.Mutex
}

func (ua *UserAges) Add(name string, age int) {
	ua.Lock()
	defer ua.Unlock()
	ua.ages[name] = age
}

func (ua *UserAges) Get(name string) int {
	ua.Lock()
	defer ua.Unlock()
	if age, ok := ua.ages[name]; ok {
		return age
	}

	return -1
}

func main() {
	ua := UserAges{ages: map[string]int{"jj": 10}}
	ua.Add("ll", 10)
	fmt.Println(ua.Get("ll"))
}
