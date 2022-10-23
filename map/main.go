package main

import (
	"fmt"
	"reflect"
)

func main() {
	func8()
}

func func1() {
	// var numbers map[string]int	//不能这样初始化map
	numbers := make(map[string]int)
	numbers["one"] = 1
	numbers["ten"] = 10
	numbers["three"] = 3
	fmt.Println("第三个数字是：", numbers["three"])

	rating := map[string]float32{"C": 5, "Go": 4.5, "Python": 4.5, "C++": 2}
	csharpRating, ok := rating["c#"]
	if ok {
		fmt.Println("c# is in the mapand its rating is ", csharpRating)
	} else {
		fmt.Println("We have no rating associated with C# in the map")
	}

	delete(rating, "C")
	fmt.Println(rating["C"])
}

func func2() {
	m := map[string]string{"hello": "hello", "world": "world"}
	changeMapItem(m, "hello")
	fmt.Println(m)
}

func changeMapItem(m map[string]string, key string) bool {
	if key == "hello" {
		m[key] = "hello2345678987654"
		return true
	}
	return false
}

func func3() {
	m := map[int]string{1: "1", 2: "2", 3: "3"}

	for key, value := range m {
		fmt.Printf("key:%d----value:%s", key, value)
	}
}

func func4() {
	oldMap := map[string]string{"1": "1", "2": "2", "3": "3"}
	newFields := []string{"1", "2", "4"}
	for _, field := range newFields {
		if _, found := oldMap[field]; found {
			delete(oldMap, field)
		}
	}
	if len(oldMap) > 0 {
		for key := range oldMap {
			newFields = append(newFields, key)
		}
	}
	fmt.Println(newFields)
}

func func5() {
	var m map[string]map[string]*string
	s := m["123"]["1"]
	fmt.Println(s)
}

func func6() {
	sn1 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}

	sn2 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}

	if sn1 == sn2 {
		fmt.Println("sn1 == sn2")
	}

	// sm1 := struct {
	// 	age int
	// 	m   map[string]string
	// }{age: 11, m: map[string]string{"a": "1"}}

	// sm2 := struct {
	// 	age int
	// 	m   map[string]string
	// }{age: 11, m: map[string]string{"a": "1"}}

	// if sm1 == sm2 {
	// 	fmt.Println("sm1 == sm2")
	// }

	// map 实例不能直接比较

	m1 := map[string]string{"a": "1"}
	m2 := map[string]string{"a": "1"}
	ok := reflect.DeepEqual(m1, m2)
	fmt.Println(ok)
	// if m1 == m2 {
	// 	fmt.Println("sm1 == sm2")
	// }
}

func func7() {
	m := make(map[int]int, 3)
	x := len(m)
	m[1] = m[1]
	y := len(m)
	println(x, y)
}

func func8() {
	var m map[string]int
	delete(m, "on")
	fmt.Println(m)
}
