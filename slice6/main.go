package main

import "fmt"

// 比较对nil切片和空切片做append操作是否有同样的效果
// 对nil切片做append操作也是可以的
func main() {
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
