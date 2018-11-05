package main

import (
	"fmt"
)

//输出所有的子集
//n是要输出的前几个字母
func CaptitalLetter(n uint) {
	letter := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J",
		"K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

	//这里为什么是2^n呢
	//详情看：
	//https://blog.csdn.net/K346K346/article/details/80436430
	var maxCount uint = 1 << n
	var i uint
	var j uint
	for i = 0; i < maxCount; i++ {
		for j = 0; j < n; j++ {
			if (i & (1 << j)) != 0 { //在做位运算的时候需要注意数据类型为uint
				fmt.Printf("%s", letter[j])
			}
		}
		fmt.Println()
	}
}

func main() {
	CaptitalLetter(5)
}
