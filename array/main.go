package main

import "fmt"

func main() {
	func1()
}

func func1() {
	var arr [10]int
	arr[0] = 42
	arr[1] = 13
	a := [3]int{1, 2, 3}
	b := [10]int{4, 5, 6}
	c := [...]int{7, 8, 9}
	fmt.Printf("the first element is %d\n", arr[0])
	fmt.Printf("the last element is %d\n", arr[9])
	fmt.Println(a[2])
	fmt.Println(b[4])
	fmt.Println(c[1])

	doubleArray := [2][4]int{[4]int{1, 2, 3, 4}, [4]int{5, 6, 7, 8}}
	easyArray := [2][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}}
	fmt.Println(doubleArray[0][2])
	fmt.Println(easyArray[0][1])
}

func func2() {
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

func func3() {
	// 声明一个二维整形数组，两个维度分别存储4个元素和2个元素
	// var array [4][2]int

	// 使用数组字面量来声明并初始化一个二维整形数组
	// array := [4][2]int{{10, 11}, {20, 21}, {30, 31}, {40, 41}}

	// 声明并初始化外层数组中索引为1和3的元素
	// array := [4][2]int{1: {20, 21}, 3: {40, 41}}

	// 声明并初始化外层数组和内层数组的单个元素
	// array := [4][2]int{1: {0: 20}, 3: {1: 41}}

	var array [2][2]int

	array[0][0] = 10
	array[0][1] = 20
	array[1][0] = 30
	array[1][1] = 40

	fmt.Println(array)
}
