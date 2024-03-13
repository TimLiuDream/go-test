package main

import (
	"fmt"
	"sort"
	"time"
)

func buildNums() []int {
	nums := make([]int, 0)
	for i := 1; i <= 100000000; i++ {
		nums = append(nums, i)
	}
	nums[100000] = 99999
	return nums
}

func func1() {
	mp := make(map[int]int, 10)
	for i := 0; i < 10; i++ {
		mp[i] = i
	}
	for key, val := range mp {
		go func() {
			fmt.Println("key: ", key, "val: ", val)
		}()
	}
}

func main() {
	func1()
	time.Sleep(time.Second * 5)

	//nums := buildNums()
	//fmt.Println(func1(nums))
	//fmt.Println(func2(nums))
	//fmt.Println(func3(nums))
}

// 使用 map
//func func1(nums []int) int {
//	m := make(map[int]struct{})
//	for _, num := range nums {
//		if _, ok := m[num]; ok {
//			return num
//		}
//		m[num] = struct{}{}
//	}
//	return -1
//}

// 排序后使用双指针
func func2(nums []int) int {
	sort.Ints(nums)
	for a, b := 0, 1; a < len(nums)-1 && b < len(nums); {
		if nums[a] == nums[b] {
			return nums[a]
		}
		a++
		b++
	}
	return -1
}

// 二分查找
func func3(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	mid := nums[len(nums)/2]
	if mid == len(nums)/2 {
		return mid
	} else if mid > len(nums)/2 {
		return func3(nums[:mid+1])
	} else {
		return func3(nums[mid:])
	}
}

//func main() {
//	func1()
//}
//
//func func1() {
//	nums := []int{1, 2, 3, 4, 5}
//	runtime.GOMAXPROCS(1)
//	g, _ := errgroup.WithContext(context.Background())
//	for _, value := range nums {
//		g.Go(func() error {
//			fmt.Print(value)
//			return nil
//		})
//	}
//	g.Wait()
//}
//
//func func2() {
//	//ss := []int{1, 1}
//	m := make(map[string]string)
//	eg, _ := errgroup.WithContext(context.Background())
//	eg.Go(func() error {
//		for i := 0; i < 100; i++ {
//			//ss[0] = ss[0] + 1
//			m["1"] = "2"
//		}
//		return nil
//	})
//	eg.Go(func() error {
//		for i := 0; i < 100; i++ {
//			m["2"] = "3"
//		}
//		return nil
//	})
//	eg.Wait()
//	fmt.Println("%+v\n", m)
//}
