package main

import "fmt"

// carPooling([][]int{{3, 2, 7}, {3, 7, 9}, {8, 3, 9}}, 11)
func carPooling(trips [][]int, capacity int) bool {
	nums := make([]int, 1000)
	da := NewDiffArr(nums)
	for _, trip := range trips {
		da.increment(trip[1], trip[2], trip[0])
	}
	result := da.result()
	fmt.Println(result)
	for _, i := range result {
		if i > capacity {
			return false
		}
	}
	return true
}

type DiffArr struct {
	diff []int
}

func NewDiffArr(nums []int) *DiffArr {
	if len(nums) == 0 {
		return &DiffArr{diff: make([]int, 0)}
	}
	// 构建差分数组
	diff := make([]int, len(nums))
	diff[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		diff[i] = nums[i] - nums[i-1]
	}
	return &DiffArr{diff: diff}
}

/* 给闭区间 [i,j] 增加 val（可以是负数）*/
func (d *DiffArr) increment(i, j int, val int) {
	d.diff[i] += val
	// if (j + 1) < len(d.diff) {
	// 	d.diff[j+1] -= val
	// }
	d.diff[j] -= val
}

/* 返回结果数组 */
func (d *DiffArr) result() []int {
	result := make([]int, len(d.diff))
	// 根据差分数组构造结果数组
	result[0] = d.diff[0]
	for i := 1; i < len(d.diff); i++ {
		result[i] = result[i-1] + d.diff[i]
	}
	return result
}
