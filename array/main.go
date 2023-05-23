package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	func12()
}

type tt struct {
	aa string
}

func joker() {
	ts := []tt{{"aa"}, {"bb"}}
	result := make([]*tt, 0, len(ts))
	for _, t := range ts {
		newT := t
		result = append(result, &newT)
	}
	fmt.Println(result)
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

	doubleArray := [2][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}}
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

type person struct {
	name string
	age  int
	job  string
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

func func4() {
	matris := [][]int{{3, 0, 1, 4, 2}, {5, 6, 3, 2, 1}, {1, 2, 0, 1, 5}, {4, 1, 0, 1, 7}, {1, 0, 3, 0, 5}}
	numMatiris := Constructor(matris)
	fmt.Println(numMatiris.SumRegion(2, 1, 4, 3))
}

type NumMatrix struct {
	matrixArr [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	m := len(matrix)
	n := len(matrix[0])
	if m == 0 || n == 0 {
		return NumMatrix{}
	}
	matrixArr := make([][]int, m+1)
	for i, _ := range matrixArr {
		matrixArr[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			matrixArr[i][j] = matrixArr[i-1][j] + matrixArr[i][j-1] + matrix[i-1][j-1] - matrixArr[i-1][j-1]
		}
	}
	return NumMatrix{matrixArr}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.matrixArr[row2+1][col2+1] - this.matrixArr[row2+1][col1] - this.matrixArr[row1][col2+1] + this.matrixArr[row1][col1]
}

func func5() {
	fmt.Println(summaryRanges([]int{0, 1, 2, 4, 5, 7}))
}

func summaryRanges(nums []int) (ans []string) {
	for i, n := 0, len(nums); i < n; {
		left := i
		for i++; i < n && nums[i-1]+1 == nums[i]; i++ {
		}
		s := strconv.Itoa(nums[left])
		if left < i-1 {
			s += "->" + strconv.Itoa(nums[i-1])
		}
		ans = append(ans, s)
	}
	return
}

func func6() {
	var a = [5]int{1, 2, 3, 4, 5}
	var r [5]int
	var b = []int{1, 2, 3, 4, 5}

	for i, v := range a { //range 表达式是副本参与循环，就是说例子中参与循环的是 a 的副本，而不是真正的 a
		if i == 0 {
			a[1] = 12
			a[2] = 13
		}
		r[i] = v
	}
	fmt.Println("r = ", r)
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.TypeOf(r))
	fmt.Println(reflect.TypeOf(b))
}

// 数组是可以比较的
func func7() {
	type pos [2]int
	a := pos{4, 5}
	b := pos{4, 5}
	fmt.Println(a == b)
}

//func func8() {
//	var nums1 []interface{}
//	nums2 := []int{1, 3, 4}
//	nums3 := append(nums1, nums2...)
//	fmt.Println(len(nums3))
//}

// 不确定长度数组，长度依赖所赋予的最大下标
// 此题中，最大的下标为 'd' 为 101，所以数组长度为 101
func func9() {
	m := [...]int{
		'a': 1,
		'b': 2,
		'c': 3,
		'd': 4,
	}
	fmt.Println(reflect.TypeOf(m))
	m['a'] = 3
	fmt.Println('a')
	fmt.Println(len(m))
}

func func10() {
	pairs := [][2]string{
		{"a", "apple"},
		{"a", "ant"},
		{"b", "bee"},
	}

	m := map[string]string{
		pairs[0][0]: pairs[0][1],
		pairs[1][0]: pairs[1][1],
		pairs[2][0]: pairs[2][1],
	}
	fmt.Println(m["a"])
}

func StringArrayDifference(old []string, new []string) (additions []string, deletions []string) {
	additionsMap := make(map[string]struct{})
	deletionsMap := make(map[string]struct{})
	for _, s := range old {
		if len(s) > 0 {
			deletionsMap[s] = struct{}{}
		}
	}
	for _, s := range new {
		if len(s) > 0 {
			additionsMap[s] = struct{}{}
		}
	}
	for s, _ := range additionsMap {
		if _, ok := deletionsMap[s]; !ok {
			additions = append(additions, s)
		}
	}
	for s, _ := range deletionsMap {
		if _, ok := additionsMap[s]; !ok {
			deletions = append(deletions, s)
		}
	}
	return
}

func func11() {
	old := []string{"4j2faC8i", "537e8Qvx", "5apEeEnu", "63vByhP7", "7AE2pRu4", "7FyBw3tg", "7JinzWNV", "7kme55qA", "7tLomwKi", "9ALdUNb2", "9yTLhCVr", "A4YVjM4P", "AB8hSwEJ", "AVtmud6H", "BBWFFHEq", "BCzt4Djq", "CE9kUdKX", "CMSMaaXa", "CZo5brMR", "CiaXT1TP", "D8YCjwXB", "DiZ6JV8P", "ETpj9fhJ", "EWxfP4JB", "F1WDUt5s", "FbTkKgog", "FczWm9Vb", "Fh4nwKF3", "FqULVjsB", "G2HL2Ss8", "G66u2kDi", "G6NmZPgx", "GLS8PMw4", "Gd1SYuAN", "Gg2eUXZe", "GoQF9u5e", "H2Ds4Nb2", "JHs3KVb1", "JUhi1yRy", "JcMRGX6A", "K46BoBYX", "KXBoeHVE", "LM4X1u5t", "LfXKyuxH", "Lq3jyFaE", "LxDHDexC", "MH4fmWAJ", "Mkuj9K5L", "N7iTTaor", "NJ7YfdCH", "QeE5UdWi", "QuARTxrk", "Qw6YBw1K", "R29VbMcV", "RBbR9jnn", "RcVtumD9", "RhC5x3pt", "Rj422Sk9", "SJaR4FsG", "T2vA22gJ", "TbUmz4Ep", "TkTFrZKe", "UDEcEZJs", "VJtf7YQN", "VNqPfhma", "VV7i4QSn", "WChdsXkc", "Wp5q24Q8", "Xevn9YcA", "XtkpjbER", "Y5zKYaLs", "YREcwt47", "jh5r86Q3"}

	new := []string{"H2Ds4Nb2",
		"JcMRGX6A",
		"RcVtumD9",
		"RhC5x3pt",
		"Wp5q24Q8",
		"7tLomwKi",
		"EWxfP4JB",
		"Y5zKYaLs",
		"jh5r86Q3",
		"9yTLhCVr",
		"Gd1SYuAN",
		"SJaR4FsG",
		"AB8hSwEJ",
		"RBbR9jnn",
		"Xevn9YcA",
		"7FyBw3tg",
		"BBWFFHEq",
		"GLS8PMw4",
		"Lq3jyFaE",
		"MH4fmWAJ",
		"T2vA22gJ",
		"XtkpjbER",
		"FczWm9Vb",
		"KXBoeHVE",
		"YREcwt47",
		"7kme55qA",
		"Mkuj9K5L",
		"N7iTTaor",
		"TkTFrZKe",
		"VNqPfhma",
		"A4YVjM4P",
		"CZo5brMR",
		"CiaXT1TP",
		"DiZ6JV8P",
		"ETpj9fhJ",
		"K46BoBYX",
		"4j2faC8i",
		"7JinzWNV",
		"AVtmud6H",
		"CE9kUdKX",
		"G6NmZPgx",
		"JHs3KVb1",
		"LxDHDexC",
		"WChdsXkc",
		"9ALdUNb2",
		"FqULVjsB",
		"Qw6YBw1K",
		"5apEeEnu",
		"G66u2kDi",
		"Gg2eUXZe",
		"Rj422Sk9",
		"TbUmz4Ep",
		"UDEcEZJs",
		"VV7i4QSn",
		"F1WDUt5s",
		"537e8Qvx",
		"63vByhP7",
		"CMSMaaXa",
		"Fh4nwKF3",
		"GoQF9u5e",
		"LM4X1u5t",
		"QeE5UdWi",
		"VJtf7YQN",
		"D8YCjwXB",
		"G2HL2Ss8",
		"JUhi1yRy",
		"LfXKyuxH",
		"NJ7YfdCH",
		"BCzt4Djq",
		"QuARTxrk",
		"R29VbMcV",
		"7AE2pRu4",
		"FbTkKgog",
		"63vByhP7",
		"7kme55qA",
		"BBWFFHEq",
		"F1WDUt5s",
		"Rj422Sk9",
		"EWxfP4JB",
		"G66u2kDi",
		"R29VbMcV",
		"UDEcEZJs",
		"5apEeEnu",
		"AB8hSwEJ",
		"CiaXT1TP",
		"K46BoBYX",
		"LM4X1u5t",
		"Qw6YBw1K",
		"RBbR9jnn",
		"VV7i4QSn",
		"XtkpjbER",
		"9ALdUNb2",
		"A4YVjM4P",
		"G6NmZPgx",
		"GoQF9u5e",
		"FbTkKgog"}

	addi, dele := StringArrayDifference(old, new)
	fmt.Println(addi)
	fmt.Println(dele)
}

// 初始化费率列表
var rates = []struct {
	min  float64
	max  float64
	rate float64
}{
	{0, 1000, 0.05},
	{1000, 5000, 0.04},
	{5000, 10000, 0.03},
	{10000, 20000, 0.02},
	{20000, 0, 0.01},
}

// 计算手续费
func calcFee(amount float64) float64 {
	var fee float64
	for _, r := range rates {
		if r.max == 0 || amount < r.max {
			fee += (amount - r.min) * r.rate
			break
		} else {
			fee += (r.max - r.min) * r.rate
			amount -= r.max - r.min
		}
	}
	return fee
}

func func12() {
	// 测试
	amount := 15000.0
	fee := calcFee(amount)
	fmt.Printf("交易金额: %.2f 元，手续费率为: %.2f%%，手续费为: %.2f 元\n", amount, fee/amount*100, fee)
}
