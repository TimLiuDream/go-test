package main

import (
	"fmt"
)

func main() {
	// 数组声明
	// var array [5]int

	// 用具体值初始化每个元素
	array := [5]int{10, 20, 30, 40, 50}

	// 用具体值初始化每个元素
	// 容量由初始化值得数量决定
	// array:=[...]int{10,20,30,40,50}

	// 用具体值初始化索引为1和2的元素
	// 其余元素保持零值
	// array:=[5]int{1:10,2:20}

	fmt.Println(array)

	// 修改索引为2的元素的值
	array[2] = 35
	fmt.Println(array)

	// 用整数指针初始化索引为0和1的数组元素
	array1 := [5]*int{0: new(int), 1: new(int)}

	*array1[0] = 10
	*array1[1] = 20

	fmt.Println(array1) //打印的结果是数组元素所在的地址

	array2 := [5]string{"Red", "Blue", "Green", "Yellow", "Pink"}
	array3 := array2
	fmt.Println(array2)
	fmt.Println(array3)
	fmt.Printf("%p\n", array2)
	fmt.Printf("%p\n", array3)

	// 声明第一个包含3个元素的指向字符串的指针数组
	var array4 [3]*string

	// 声明第二个包含3个元素的指向字符串的指针数组
	// 使用字符串指针初始化这个数组
	array5 := [3]*string{new(string), new(string), new(string)}

	// 使用颜色为每个元素赋值
	*array5[0] = "Red"
	*array5[1] = "Blue"
	*array5[2] = "Green"

	// 将array5复制array4
	array4 = array5

	fmt.Println(array4)
	fmt.Println(array5)
}
