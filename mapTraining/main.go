package main

import "fmt"

func main() {
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
