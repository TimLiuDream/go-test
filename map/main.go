package main

import "fmt"

func main() {
	func1()
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
