package main

import (
	"fmt"
	"sort"
	"unsafe"
)

func main() {
	func12()
}

func func1() {
	slice := []byte{'a', 'b', 'c', 'd'}
	fmt.Println(slice[3])

	var ar = [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
	var a, b []byte
	a = ar[2:5]
	b = ar[3:5]
	// c = ar[0:0] // 索引越界
	fmt.Println(a[0])
	fmt.Println(b[0])
	// fmt.Println(c[0])
}

func func2() {
	//说先定义一个切片，只限定长度为1
	s := make([]int, 1)

	//打印出slice的长度，容量以及内存地址
	fmt.Printf("len :%d cap:%d array ptr :%v \n", len(s), cap(s), *(*unsafe.Pointer)(unsafe.Pointer(&s)))

	for i := 0; i < 5; i++ {
		s = append(s, i)
		fmt.Printf("len :%d cap:%d array ptr :%v \n", len(s), cap(s), *(*unsafe.Pointer)(unsafe.Pointer(&s)))
	}
	//打印出slice
	fmt.Println("array:", s)
}

func func3() {
	//说先定义一个切片，只限定长度为1
	s := make([]int, 1)

	//打印出slice的长度，容量以及内存地址
	fmt.Printf("len :%d cap:%d array ptr :%v \n", len(s), cap(s), *(*unsafe.Pointer)(unsafe.Pointer(&s)))

	for i := 1; i < 1024*2; i++ {
		s = append(s, i)
		fmt.Printf("len :%d cap:%d array ptr :%v \n", len(s), cap(s), *(*unsafe.Pointer)(unsafe.Pointer(&s)))
	}
	//打印出slice
	fmt.Println("array:", s)
}

func func4() {
	a := []int{1, 3, 6, 10, 15, 21, 28, 36, 45, 55}
	x := 6

	i := sort.Search(len(a), func(i int) bool { return a[i] == 2 })
	if i < len(a) && a[i] == x {
		fmt.Printf("found %d at index %d in %v\n", x, i, a)
	} else {
		fmt.Printf("%d not found in %v\n", x, a)
	}
}

func func5() {
	s := []int{1, 3, 4, 7, 87, 3, 6, 3247, 88, 4, 9}
	for index := 0; index < len(s); index++ {
		fmt.Println(s[index])
	}
	for index, value := range s {
		// if index == 5 {
		// 	value = 10
		// }
		s[index] = 10
		fmt.Println("range:", index, value)
	}
	fmt.Println(s)
}

func func6() {
	s := []int{1, 2, 3, 4, 5, 6, 6}
	cpes := make([]int, len(s))
	cphs := make([]int, len(s))
	copy(cpes, s)
	copy(cphs, s)
	es := cpes[4:]
	fmt.Println(es)
	hs := cphs[:4]
	fmt.Println(hs)
	hs = append(hs, []int{111}...)
	fmt.Println(hs)
	s = append(hs, es...)
	fmt.Println(s)
}

// 比较对nil切片和空切片做append操作是否有同样的效果
// 对nil切片做append操作也是可以的
func func7() {
	// 定义一个nil的切片
	var nums []int
	// 定义一个空切片
	nums1 := make([]int, 0)
	// 分别nil切片和空切片append
	nums = append(nums, 1)
	nums1 = append(nums1, 1)
	fmt.Println(nums)
	fmt.Println(nums1)
}

// 长度与容量的区别
func func8() {
	slice := make([]int, 3, 5)
	slice[0] = 10
	slice[1] = 20
	slice[2] = 30
	fmt.Println(slice[2])
	fmt.Println(slice[4]) // index out of range
}

func func9() {
	s := []string{"A", "B", "C"}

	t := s[:1]
	fmt.Println(&s[0] == &t[0])

	u := append(s[:1], s[2:]...)
	fmt.Println(&s[0] == &u[0])
}

func func10() {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(s, len(s), cap(s))
	sResult := append(s, 11)
	fmt.Println(sResult, len(sResult), cap(sResult))

	s1 := s[0:5]
	fmt.Println(s1, len(s1), cap(s1))

	s2 := s[5:]
	fmt.Println(s2, len(s2), cap(s2))
}

func func11() {
	a := []int{5}
	for range a {
		a = append(a, 1)
	}
	fmt.Println(len(a))
}

func func12() {
	s := []int{0}
	sValue := 1
	newS := appendFunc121(s, sValue)
	fmt.Println(s)
	fmt.Println(newS)
}

func appendFunc12(fs []int, v int) {
	fs = append(fs, v)
}

func appendFunc121(fs []int, v int) []int {
	newFs := make([]int, 0)
	newFs = append(fs, v)
	return newFs
}

// 确定两个 slice 是否包含
func func13() {
	s1 := []int{1, 2, 3, 4, 5, 6}
	s2 := []int{2, 5}

	// 方式一
	m := make(map[int]struct{})
	for _, v := range s1 {
		m[v] = struct{}{}
	}
	for _, v := range s2 {
		if _, ok := m[v]; !ok {
			fmt.Println(false)
		}
	}

	// 方式二
	for _, v2 := range s2 {
		found := false
		for _, v1 := range s1 {
			if v2 == v1 {
				found = true
				break
			}
		}
		if !found {
			fmt.Println(false)
		}
	}
}
